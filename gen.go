/**
Copyright 2020 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
**/

// This Go file helps generate the library.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const workplacePath = "workplace"
const schemaFileName = "events.json"
const pkgPrefix = "google.events"
const funcsTmplFileName = "marshalUnmarshalFuncs.tmpl"
const readMeTmplFileName = "README.tmpl"

const funcsToReplace = `func UnmarshalEvents(data []byte) (Events, error) {
	var r Events
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Events) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
`

type eventSpec struct {
	Title      string
	GoPackage  string
	Properties map[string]map[string]interface{}
}

type event struct {
	Package     string
	Name        string
	Description string
}

// locateConfigs scans the workplace directory and collects the paths to event
// specifications.
func locateConfigs(workplacePath string) []string {
	configPaths := []string{}
	err := filepath.Walk(workplacePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(fmt.Sprintf("error when processing path %q: %v\n", path, err))
		}
		if !info.IsDir() && info.Name() == schemaFileName {
			configPaths = append(configPaths, path)
		}
		return nil
	})
	if err != nil {
		panic(fmt.Sprintf("error when walking the directory: %v\n", err))
	}

	return configPaths
}

// extractFileStructAndPkgName unmarshals the event specifications and
// collects information in eventSpec structs
func extractFileStructAndPkgName(configPath string) eventSpec {
	dat, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(fmt.Sprintf("error when opening the event spec %q: %v\n", configPath, err))
	}

	spec := eventSpec{}
	err = json.Unmarshal(dat, &spec)
	if err != nil {
		panic(fmt.Sprintf("error when unmarshaling the event spec %q: %v\n", configPath, err))
	}

	return spec
}

// genSrc generates the source code using quicktype.
func genSrc(configPath string) string {
	args := []string{
		"--src",
		configPath,
		"--src-lang",
		"schema",
		"--lang",
		"go",
	}
	cmd := exec.Command("quicktype", args...)
	out, err := cmd.Output()
	if err != nil {
		panic(fmt.Sprintf("error when calling quicktype library: %v\n", err))
	}

	return string(out)
}

// mkAllDirs creates the directories for generated code.
func mkAllDirs(spec eventSpec) string {
	path := spec.Title
	path = strings.TrimPrefix(path, pkgPrefix)
	path = strings.TrimPrefix(path, ".")
	path = strings.ReplaceAll(path, ".", "/")

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("error when creating directories %q: %v\n", path, err))
	}
	return path
}

// updatePkdStmt patches the package statements in the generated code
func updatePkgStmt(spec eventSpec, out string) string {
	pkg := spec.GoPackage
	return strings.Replace(out, "package main", fmt.Sprintf("package %s", pkg), 1)
}

// replaceMarshalUnMarshalFuncs patches the utility funcs in the generated code
func replaceMarshalUnMarshalFuncs(out string, events []event) string {
	replacementFlag := "**flag for replacement**"
	cleanedOut := strings.Replace(out, funcsToReplace, replacementFlag, 1)

	funcsTmplDat, err := ioutil.ReadFile(funcsTmplFileName)
	if err != nil {
		panic(fmt.Sprintf("error when reading a template: %v", err))
	}
	tmpl, err := template.New("marshalUnmarshalFuncs").Parse(string(funcsTmplDat))
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, events)
	if err != nil {
		panic(err)
	}
	newMarshalUnmarshalFuncs := buf.String()

	return strings.Replace(cleanedOut, replacementFlag, newMarshalUnmarshalFuncs, 1)
}

// writeReadMe generates and writes the README file.
func writeReadMe(events []event) {
	readMeTmplDat, err := ioutil.ReadFile(readMeTmplFileName)
	if err != nil {
		panic(fmt.Sprintf("error when reading a template: %v", err))
	}
	tmpl, err := template.New("readMe").Parse(string(readMeTmplDat))
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, events)
	if err != nil {
		panic(err)
	}
	readMe := buf.String()

	err = ioutil.WriteFile("README.md", []byte(readMe), 0644)
	if err != nil {
		panic(fmt.Sprintf("error when writing README file: %v", err))
	}
}

func main() {
	configPaths := locateConfigs(workplacePath)
	allEvents := []event{}
	for _, path := range configPaths {
		spec := extractFileStructAndPkgName(path)
		out := genSrc(path)
		path := mkAllDirs(spec)
		out = updatePkgStmt(spec, out)

		events := []event{}
		for k, v := range spec.Properties {
			desc, found := v["description"]
			if !found {
				desc = "N/A"
			}
			events = append(events, event{Package: spec.GoPackage, Name: k, Description: fmt.Sprintf("%s", desc)})
		}
		allEvents = append(allEvents, events...)
		out = replaceMarshalUnMarshalFuncs(out, events)

		err := ioutil.WriteFile(fmt.Sprintf("%s/events.go", path), []byte(out), 0644)
		if err != nil {
			panic(fmt.Sprintf("error when writing file %q: %v", path, err))
		}
	}
	writeReadMe(allEvents)
}
