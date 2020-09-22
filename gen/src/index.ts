#!/usr/bin/env node
const yargs = require('yargs');
const mkdirp = require('mkdirp');
const fs = require('fs');
const sqrl = require('squirrelly');

const qt = require('qt');

let IN = yargs.argv.in || process.env.IN;
let OUT = yargs.argv.out || process.env.OUT;

const LANGUAGE = 'go';
const TEMPLATE_DIRECTORY = 'templates';
const GO_FILENAME = 'events.go';
const SNIPPET_DIRECTORY = 'snippets';
const MARSHAL_FUNCS_TO_REPLACE_SNIPPET = 'marshalFuncsToReplace.snippet';
const MARSHAL_FUNCS_TEMPLATE = 'marshalFuncs.squirrelly';
const README_TEMPLATE = 'README.squirrelly';

async function main() {
  if (!IN) console.error('Error in config: `IN` not set');
  if (!OUT) console.error('Error in config: `OUT` not set');
  if (!IN || !OUT) return;
  if (IN.endsWith('/')) IN = IN.substring(0, IN.length - 1);
  if (OUT.endsWith('/')) OUT = OUT.substring(0, OUT.length - 1);

  const templateDirectoryPath = `${__dirname}/../../${TEMPLATE_DIRECTORY}`;
  const snippetDirectoryPath = `${__dirname}/../../${SNIPPET_DIRECTORY}`;

  // Collect all schemas and generate Go code from them using the
  // quicktype-wrapper package
  const schemasAndGenFiles = await qt.getJSONSchemasAndGenFiles(IN, LANGUAGE);
  const allEvents: Array<[string, string, string]> = [];
  schemasAndGenFiles.map((sg: [any, string]) => {
    const schema = sg[0];
    let genFile = sg[1];

    // Make the directory for all events in the schema
    let pkg = schema['$id'];
    pkg = pkg.replace('google.events.', '');
    const pkgPath = pkg.replace(/\./g, '/');
    mkdirp.sync(`${OUT}/${pkgPath}`);

    // Update the package statement
    const goPkg = schema['goPackage'];
    const goPkgStmt = `package ${goPkg}`;
    genFile = genFile.replace('package main', goPkgStmt);

    // Collect all events in the schema
    const pkgEvents: string[] = [];
    Object.keys(schema.properties).map((eventName: string) => {
      const eventDescription = schema.properties[eventName].description;
      pkgEvents.push(eventName);
      allEvents.push([pkg, eventName, eventDescription]);
    });

    // Replace the utility funcs in the generated
    const marshalFuncsSnippet = fs.readFileSync(
      `${snippetDirectoryPath}/${MARSHAL_FUNCS_TO_REPLACE_SNIPPET}`
    );
    const marshalFuncsTmpl = fs.readFileSync(
      `${templateDirectoryPath}/${MARSHAL_FUNCS_TEMPLATE}`
    );
    const marshalFuncs = sqrl.render(String(marshalFuncsTmpl), {
      pkgEvents: pkgEvents,
    });
    genFile = genFile.replace(String(marshalFuncsSnippet), marshalFuncs);

    // Write the patched generated code
    fs.writeFileSync(`${OUT}/${pkgPath}/${GO_FILENAME}`, genFile);
  });

  // Write the README.md file
  const readMeSqrlTmpl = fs.readFileSync(
    `${templateDirectoryPath}/${README_TEMPLATE}`
  );
  const readMe = sqrl.render(String(readMeSqrlTmpl), {allEvents: allEvents});
  fs.writeFileSync(`${OUT}/README.md`, readMe);
}

if (!module.parent) {
  main();
}
