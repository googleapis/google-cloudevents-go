const fetch = require('node-fetch');
const fs = require('fs');

// Config
const CATALOG_URL = 'https://googleapis.github.io/google-cloudevents/jsonschema/catalog.json';
const README_START = '<!-- GENERATED START -->';
const README_END = '<!-- GENERATED END -->';

/**
 * Gets markdown type documentation for this library.
 * @param {object} catalog The JSON schema event catalog.
 * @example
 * {
 *   url: 'https://googleapis.github.io/google-cloudevents/jsonschema/google/events/cloud/firestore/v1/DocumentEventData.json',
 *   product: 'Cloud Firestore',
 *   name: 'DocumentEventData',
 *   description: 'The data within all Firestore document events.',
 *   datatype: 'google.events.cloud.firestore.v1.DocumentEventData',
 *   cloudeventTypes: ['first', 'second']
 * }
 * @returns {string} Markdown for the Node type documentation.
 */
function getTypeDocumentation(catalog: any): string {
  const result: any[] = catalog.schemas.map((schema: any) => {
    // Get Go package and path.
    // E.g. "google.events.firebase.database.v1.ReferenceEventData" -> "database"
    const dataTypeSplit = schema.datatype.split('.');
    dataTypeSplit.shift(); // Remove "google"
    dataTypeSplit.shift(); // Remove "events"
    dataTypeSplit.pop(); // Remove data (i.e. ReferenceEventData)
    const goPackage = dataTypeSplit[dataTypeSplit.length - 2]; // 2nd to last element
    const goPath = `github.com/googleapis/google-cloudevents-go/${dataTypeSplit.join('/')}`;
    return '' +
`### ${schema.product}

${schema.description}

#### CloudEvent Types:

${schema.cloudeventTypes?.map((t: string) => `- \`${t}\``).join('\n')}

#### Sample
\`\`\`go
package main

import (
	"fmt"

	${goPackage} "${goPath}"
)

func main() {
	data := []byte("CloudEvent Data in bytes")

	e, err := ${goPackage}.Unmarshal${schema.name}(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\\n", e)
}
\`\`\`
`;
  });
  return result.join('\n');
}

/**
 * Replaces the contents of a string's GENERATED comments with a replacement.
 * @param {string} s The string to replace.
 * @param {string} replacement The replacement string.
 */
const getGeneratedStringWithReplacement = (s: string, replacement: string) => {
  const README_GEN_START = '<!-- GENERATED START -->';
  const README_GEN_END = '<!-- GENERATED END -->';

  const README_BEFORE_TABLE = s.substring(0, s.indexOf(README_GEN_START) + README_GEN_START.length);
  const README_AFTER_TABLE = s.substring(s.indexOf(README_GEN_END));

  // Return result (with newlines)
  return `${README_BEFORE_TABLE}
${replacement}
${README_AFTER_TABLE}`;
};

// Runs the generator
(async () => {
  const catalog = await fetch(CATALOG_URL);
  const catalogJSON = await catalog.json();

  const typeDocumentation = getTypeDocumentation(catalogJSON);

  // Root reference.md
  const README_PATH = `${__dirname}/../../../reference.md`;
  const readmeContents = fs.readFileSync(README_PATH).toString();
  const updatedReadmeContents = getGeneratedStringWithReplacement(readmeContents, typeDocumentation);

  // Save updated reference
  fs.writeFileSync(README_PATH, updatedReadmeContents);
})();
