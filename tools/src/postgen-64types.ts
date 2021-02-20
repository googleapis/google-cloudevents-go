
/**
 * This file adds the "string" annotation to all numbers with 64 bits.
 *
 * Properties that are numbers with 64 bits are (always) converted into a string
 * when serialized into JSON.
 *
 * @see https://developers.google.com/protocol-buffers/docs/proto3#json
 *
 * Note: This library strictly expects JSON serialization for CloudEvent data payloads.
 *
 * @example
 * {
 *   "generation": "1610564600595212"
 * }
 *
 * To parse these values into an idiomatic value, i.e. an "int64" rather than "string", we modify
 * the Golang fields for 64 bit types (int64, uint64, float64):
 *
 * Before:
 *     Generation              *int64              `json:"generation,omitempty"`
 * After:
 *     Generation              *int64              `json:"generation,string,omitempty"`
 *
 * This modification currently cannot be done at the JSON schema + Quicktype level,
 * as JSON schema does not have a representation for 64 bit numbers serialized as strings.
 * @param {string} golangFile
 * @returns {string} The golang file with fixed 64
 */
export const fix64BitNumberFields = (golangFile: string) => {
  const lines = golangFile.split('\n');
  const patchedLines = lines.map((line: string) => {
    // For lines that have a 64 bit field
    if (line.includes('*int64') ||
        line.includes('*uint64') ||
        line.includes('*floatint64')) {
      line = line.replace(",omitempty", ",string,omitempty");
    }
    return line;
  });
  return patchedLines.join('\n');
};
