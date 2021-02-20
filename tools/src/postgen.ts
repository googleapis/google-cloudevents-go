import * as fs from 'fs';
import * as path from 'path';
import {fix64BitNumberFields} from './postgen-64types';
const pascalcase = require('pascalcase');
const recursive = require("recursive-readdir");

/**
 * This tool makes the generated code from the quicktype wrapper usable to Go clients.
 * 
 * Prerequisites:
 * - A `cloud/` folder with `*Data.go` files within subfolders.
 * 
 * After gen, this postgen scripts modififies the files this way:
 * - Add package based on folder + version (i.e. "pubsubv1")
 * - Add 'import "encoding/json"' for unmashal function
 * - Add "Unmarshal(data []byte)" function
 * - Add "(d *MyData) Marshal() ([]byte error)" function
 */

// The abs path of the repo root
const REPO_ROOT = path.dirname(process.cwd());

/**
 * Creates the Unmarshal and Marshal functions for this struct.
 * @param {string} dataFieldName The Pascal Case data name.
 * @example marshalUnmarshalFunctions('MessagePublishedData')
 */
const marshalUnmarshalFunctions = (dataFieldName: string) => `
func Unmarshal${dataFieldName}(data []byte) (${dataFieldName}, error) {
	var d ${dataFieldName}
	err := json.Unmarshal(data, &d)
	return d, err
}

func (p *${dataFieldName}) Marshal${dataFieldName}() ([]byte, error) {
	return json.Marshal(p)
}`;

/**
 * Runs post-gen processing on the generated files:
 * - Fixes 64 bit types on the structs.
 * - Adds the correct Go package.
 * - Adds JSON Unmarshal and Marshal functions.
 */
async function main() {
  const filePaths: string[] = [
    ...await recursive(`${REPO_ROOT}/cloud`),
    ...await recursive(`${REPO_ROOT}/firebase`),
  ];
  
  // For each schema
  filePaths.forEach(filePath => {
    // Read file
    const typeFileContent = fs.readFileSync(filePath).toString();
    const fixedFileContents = fix64BitNumberFields(typeFileContent);

    // Get relative path info
    const relativePath = filePath.substr(REPO_ROOT.length + 1);

    // Get Data field name from file path, in PascalCase
    const dataField = pascalcase(path.parse(path.basename(filePath)).name);

    // Get package info
    // 
    // IN: relativePath = cloud/storage/v1/StorageObjectData.go
    // OUT: goPackage = storage
    // OUT: version = v1
    // OUT: dataFilename = StorageObjectData.go
    const splitPath = relativePath.split('/');
    const [goPackage, version, dataFilename] = splitPath.slice(-3); // last 3 elements
    
    // Create the full file
    const typeFileWithMarshalling = 
      fixedFileContents +
      marshalUnmarshalFunctions(dataField);

    // Add Go package header (but under Apache header)
    const filePackageHeader =
`package ${goPackage}

import "encoding/json"`;
    let typeFileWithMarshallingAndPackage =
      typeFileWithMarshalling.split('\n');
    typeFileWithMarshallingAndPackage.splice(
      13, // length of Apache header from quicktype wrapper
      0, // don't delete anything
      filePackageHeader, // insert header at this index
    );
    const typeFile = typeFileWithMarshallingAndPackage.join('\n');

    // Write final type file
    fs.writeFileSync(filePath, typeFile);
  });
}

main();
