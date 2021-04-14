const pascalcase = require('pascalcase');

/**
 * This file converts string types with:
 * - byte data
 * to
 * - byte types
 *
 * ## Example
 * ### Before
 * type Message struct {
 *   Data        *string        `json:"data,omitempty"`        // The binary data in the message.
 * }
 * 
 * ### After
 * type Message struct {
 *   Data        []byte        `json:"data,omitempty"`        // The binary data in the message.
 * }
 * 
 * Also handles repeated proto fields.
 * 
 * ---
 * 
 * Does this by inspecting the proto, and modifing the source where appropriate.
 * 
 * This modification currently cannot be done at the JSON schema + Quicktype level,
 * as JSON schema + Quicktype does create base64 formatted strings.
 * @see https://github.com/quicktype/quicktype/issues/138
 * @param {string} golangFile The golang file source code
 * @param {string} protoFile The proto source
 * @returns {string} The golang file with fixed base64 fields
 */
export const fixByteFields = (golangFile: string, protoFile: string) => {
  // A proto line with bytes
  type LineWithByte = {
    field: string;
    isRepeated: boolean;
  };
  
  // Get all proto lines with the "bytes" type:
  const protoLines = protoFile.split('\n');
  const protoLinesWithBytes: LineWithByte[] = [];
  protoLines.map((line: string) => {
    /**
     * Include lines from our proto files.
     * The two spaces in front are part of the file identifiers as we do not do AST parsing.
     * An eventual solution to this would be to fix the byte fields upstream in the jsonschema
     * and quicktype generators.
     * @example repeated bytes build_step_outputs = 6;
     * @example bytes data = 1;
     * @example bytes custom_data = 1;
     */
    const isBytes = line.includes('  bytes ');
    const isRepeatedBytes = line.includes('  repeated bytes ');
    if (isBytes || isRepeatedBytes) {
      // The field is right before the " = " sign:
      const fieldTokens = line.split(' = ')[0].split(' ');
      const field = pascalcase(fieldTokens[fieldTokens.length - 1]);
      
      // Add the line to the array
      protoLinesWithBytes.push({
        field,
        isRepeated: isRepeatedBytes,
      });
    }
  });

  // For all proto "bytes" fields,
  // Change the type to []byte instead of *string.
  let updatedGolangFile = golangFile;
  protoLinesWithBytes.forEach((byteLine: LineWithByte) => {
    // Replace the golang field with the same name
    updatedGolangFile = golangFile.split('\n').map((golangLine: string, i: number) => {
      // Match on the exact field (with tab and space)
      const lineIncludesFieldname = golangLine.includes(`	${byteLine.field} `);
      if (lineIncludesFieldname) {
        if (byteLine.isRepeated) {
          // If repeated proto bytes, replace the type "[]string" to "[][]byte"
          return golangLine.replace('[]string', '[][]byte');
        } else {
          // If non-repeated proto bytes, replace the type "*string" to "[]byte"
          return golangLine.replace('*string', '[]byte');
        }
      }
      return golangLine;
    }).join('\n');
  });
  return updatedGolangFile;
};
