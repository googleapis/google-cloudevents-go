package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const TypePrefix = "google.events."
const SrcPrefix = "google/events/"

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(0)

	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			// Protos intended for the shared/ directory are common resources.
			// They are not event data types.
			if strings.Contains(f.GoImportPath.String(), "third_party") {
				continue
			}

			generateTests(gen, f)
			generateDocs(gen, f)

		}
		return nil
	})
}

var validationTestTpl *template.Template

func init() {
	p := os.Getenv("GENERATE_TEMPLATE_DIR")
	if p == "" {
		log.Fatal("Missing environment variable GENERATE_TEMPLATE_DIR")
	}
	validationTestTpl = template.Must(template.ParseFiles(filepath.Join(p, "validationtest.gotpl")))
}

// TestParams organizes all template parameters for test generation.
type TestParams struct {
	// DataTypes is a list of payload data structure names like "LogEntryData".
	DataTypes []string
	// The type prefix identifies a versioned family of events, such as "google.cloud.audit.log.v1"
	TypePrefix string
	// SrcPath is the package path within the module, such as cloud/auditdata
	SrcPath string
	// Pkg is the package name.
	Pkg string
	// TestDataPath is the path from GENERATOR_DATA_PATH to examples for this event type family.
	TestDataPath string
	// SourceFile is the proto used to generate code.
	SourceFile string
	// PrepareFunction defines a function to call to clean up the data.
	// If set, the strict subtest will be skipped and a "compatibility" subtest created instead.
	// It is mapped to the DataType.
	PrepareFunction map[string]string

	// Protoc tooling versions.
	ProtocVersion      string
	ProtocGenGoVersion string
	LibraryVersion     string
}

// generateTests generates test coverage per type.
func generateTests(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := filepath.Join(filepath.Dir(file.GeneratedFilenamePrefix), "data_test.go")
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	params := TestParams{
		Pkg:        string(file.GoPackageName),
		SrcPath:    strings.TrimPrefix(strings.Trim(file.GoImportPath.String(), "\""), SrcPrefix),
		TypePrefix: string(file.Desc.Package()),
		SourceFile: file.Desc.Path(),
	}

	// Derive the protoc version.
	protocVersion := "(unknown)"
	if v := gen.Request.GetCompilerVersion(); v != nil {
		protocVersion = fmt.Sprintf("v%v.%v.%v", v.GetMajor(), v.GetMinor(), v.GetPatch())
	}
	params.ProtocVersion = protocVersion

	// Derive protoc-gen-go version.
	pgcVersion := "(unknown)"
	if v := os.Getenv("PROTOC_GEN_GO_VERSION"); v != "" {
		pgcVersion = v
	}
	params.ProtocGenGoVersion = pgcVersion

	// Derive current library version.
	libVersion := "(unknown)"
	if v := os.Getenv("LIBRARY_VERSION"); v != "" {
		libVersion = v
	}
	params.LibraryVersion = libVersion

	// Notes on parsing a proto message:
	// msg.GoIdent.GoName => Event Name like "FunctionCreatedEvent"
	// msg.Desc.FullName() => CloudEvent Type like "google.events.cloud.functions.v2.FunctionCreatedEvent"
	// msg.Desc.Fields().ByName("data").Message().FullName() => Event Data Type like "FunctionEventData"

	// Assemble DataTypes
	dataTypeMap := make(map[string]string, len(file.Messages))
	for _, msg := range file.Messages {
		dataType := filepath.Ext(string(msg.Desc.Fields().ByName("data").Message().FullName()))[1:]
		dataTypeMap[dataType] = msg.GoIdent.GoName
	}
	var dataTypeSlice []string
	for i := range dataTypeMap {
		dataTypeSlice = append(dataTypeSlice, i)
	}
	sort.Strings(dataTypeSlice)
	params.DataTypes = dataTypeSlice

	// Create TestDataPath like "google/events/cloud/functions/v1"
	re := regexp.MustCompile(`(v\d)$`)
	version := re.FindString(string(file.GoPackageName))
	product := strings.TrimSuffix(string(file.GoPackageName), version)
	product = strings.TrimSuffix(product, "data")
	if version == "" {
		version = "v1"
	}

	testDataPath := filepath.Join(SrcPrefix, filepath.Dir(params.SrcPath), product, version)
	params.TestDataPath = testDataPath

	// Identify if a data cleaning step is required and assign the cleanup function.
	params.PrepareFunction = make(map[string]string, len(params.DataTypes))
	for _, dataType := range params.DataTypes {
		switch {
		case params.TypePrefix == "google.events.cloud.pubsub.v1" && dataType == "MessagePublishedData":
			params.PrepareFunction[dataType] = "testhelper.PreparePubSubMessagePublishedData"
		case params.TypePrefix == "google.events.cloud.audit.v1" && dataType == "LogEntryData":
			params.PrepareFunction[dataType] = "testhelper.PrepareAuditLogEntryData"
		}
	}

	var b bytes.Buffer
	if err := validationTestTpl.Execute(&b, params); err != nil {
		log.Fatal(err)
	}

	g.P(b.String())
	b.Reset()

	logGeneratedFileFromProto(file, filename)

	return g
}

// generateDocs generates package docs.
func generateDocs(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := filepath.Join(filepath.Dir(file.GeneratedFilenamePrefix), "doc.go")
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	product := getCustomField(gen, file.Desc.Options(), "cloud_event_product")
	if product == "" {
		panic("could not parse product name")
	}
	g.P("// Package ", string(file.GoPackageName), " provides ", product, " type definitions for CloudEvent data payloads.")
	g.P("//")
	g.P("// # Supported CloudEvent Types")
	g.P("//")
	for _, msg := range file.Messages {
		// Access the comment describing the event type:
		// d := strings.TrimSpace(strings.TrimPrefix(msg.Comments.Leading.String(), "// "))
		t := getCustomField(gen, msg.Desc.Options(), "cloud_event_type")
		if t != "" {
			g.P("//   - ", t)
		}
	}

	g.P("package ", string(file.GoPackageName))

	logGeneratedFileFromProto(file, filename)

	return g
}

// getCustomField retrieves a field defined in the google/events/cloud_event.proto.
// Custom options are complex to reflect:
// - https://github.com/golang/protobuf/issues/1260.
// - https://github.com/jhump/protoreflect/issues/377
//
// This approach renders all options to a string then parses the value based on
// looking up the field number assigned in the extension descriptor.
func getCustomField(gen *protogen.Plugin, options protoreflect.ProtoMessage, name string) string {
	var id int
	for _, file := range gen.Request.ProtoFile {
		if *file.Name != "google/events/cloudevent.proto" {
			continue
		}
		for _, ext := range file.GetExtension() {
			if *ext.Name == name {
				id = int(*ext.Number)
				break
			}
		}
	}

	return parseValueFromOptions(options, strconv.Itoa(id))
}

// parseValueFromOptionString extracts the value associated with a particular
// field number from the Stringified options.
func parseValueFromOptions(o protoreflect.ProtoMessage, id string) string {
	// Example section of s: 11716487:\"API Gateway\"
	s := fmt.Sprintf("%v", o)
	// Split on the options ID.
	a := strings.Split(s, id)
	if len(a) < 2 {
		return ""
	}

	// Data is formatted as the value in the next set of doublequotes.
	b := strings.Split(a[1], `"`)
	if len(b) < 2 {
		return ""
	}

	return b[1]
}

func logGeneratedFileFromProto(file *protogen.File, filename string) {
	re := regexp.MustCompile("^(.+)data")
	m := re.FindStringSubmatch(string(file.GoPackageName))
	log.Printf("- %s: %s => %s", m[1], file.Desc.Path(), filename)
}
