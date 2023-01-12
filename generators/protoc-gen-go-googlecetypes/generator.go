package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
)

const TypePrefix = "google.events."
const SrcPrefix = "google/events/"

func main() {
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

	return g
}
