import * as fs from 'fs';
import * as path from 'path';
import {fix64BitNumberFields} from './postgen-64types';
import {fixByteFields} from './postgen-bytetypes';
const recursive = require("recursive-readdir");

/**
 * This tool makes the generated code from the quicktype wrapper usable to Go clients.
 * 
 * Prerequisites:
 * - A `cloud/` folder with `*Data.go` files within subfolders.
 * - A `firebase/` folder with `*Data.go` files within subfolders.
 * - The cloned Google CloudEvents repo in `../google-cloudevents`
 * 
 * After gen, this postgen scripts modififies the files this way:
 * - Add package based on folder (i.e. "pubsub")
 * - Fixes 64 bit types on the structs.
 * - Adds the correct Go package.
 * - Fixes map of strings interfaces.
 * - Fixes godoc string prefixes.
 * - Fixes []byte fields.
 */

// The abs path of the repo root
const REPO_ROOT = path.dirname(process.cwd());
const PROTO_ROOT = path.resolve(REPO_ROOT, '..', 'google-cloudevents', 'proto');

/**
 * Runs post-gen processing on the generated files.
 */
async function main() {
  const filePaths: string[] = [
    ...await recursive(`${REPO_ROOT}/cloud`),
    ...await recursive(`${REPO_ROOT}/firebase`),
  ];
  
  // For each schema
  filePaths.forEach(filePath => {
    // Get relative path info
    const relativePath = filePath.substr(REPO_ROOT.length + 1);

    // Get proto file
    const relativeProtoPath = `google/events/${relativePath.split('/').slice(0, -1).join('/')}/data.proto`;
    const protoPath = path.join(PROTO_ROOT, relativeProtoPath);
    const protoFile = fs.readFileSync(protoPath).toString();

    // Read file and apply fixes
    const [typeFileContent] = [fs.readFileSync(filePath).toString()]
      .map(fix64BitNumberFields)
      .map((s: string) => fixByteFields(s, protoFile))
      .map(fixMapStringToInterface)
      .map(fixTypeGoDocPrefix)
    ;

    // Get package info
    // 
    // IN: relativePath = cloud/storage/v1/StorageObjectData.go
    // OUT: goPackage = storage
    // OUT: version = v1
    // OUT: dataFilename = StorageObjectData.go
    const splitPath = relativePath.split('/');
    const [goPackage, version, dataFilename] = splitPath.slice(-3); // last 3 elements

    // Add Go package header (but under Apache header, with newline)
    const filePackageHeader =
`
package ${goPackage}
`;
    let typeFileWithPackage =
    typeFileContent.split('\n');
    typeFileWithPackage.splice(
      13, // length of Apache header from quicktype wrapper
      0, // don't delete anything
      filePackageHeader, // insert header at this index
    );
    const typeFile = typeFileWithPackage.join('\n');

    // Write final type file
    fs.writeFileSync(filePath, typeFile);
  });
}

/**
 * Converts:
 * - map[string]map[string]interface{}
 * to
 * - map[string]interface{}
 * @param {string} s The string to update.
 */
function fixMapStringToInterface(s: string): string {
  const mapStringToStringToInterface = 'map[string]map[string]interface{}';
  const mapStringToInterface = 'map[string]interface{}';
  return s.split(mapStringToStringToInterface).join(mapStringToInterface);
}

/**
 * Prefixes a golang type with "TypeName: " to match golang convention.
 * @param {string} file The file as a string.
 * @returns {string} The file with each type prefixed with the type name.
 */
function fixTypeGoDocPrefix(file: string): string {
  const lines = file.split('\n');
  lines.forEach((line: string, i: number) => {
    // If we're defining a struct, there must be a comment!
    // Update the comment.
    if (line.includes(' struct {')) {
      const typeName = line.split(' ')[1];
      // Look back at the comments until we hit the a blank newline
      let commentIndex = i - 1;
      while (lines[commentIndex].startsWith('//')) {
        --commentIndex;
      }
      // If there was no comment before the type, there is nothing to fix. Return.
      if (commentIndex === i - 1) return;

      // We want the line _after_ the blank line (this is the first line of the comment)
      ++commentIndex;

      // Add the "TypeName: " string at the first line
      // Before: "// Last set value of user property."
      // After: "// Value: Last set value of user property."
      const updateLine = `// ${typeName}: ${lines[commentIndex].substr('// '.length)}`;
      lines[commentIndex] = updateLine;
    }
  });
  return lines.join('\n');
}

main();
