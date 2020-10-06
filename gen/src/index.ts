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
const README_TEMPLATE = 'README.squirrelly';

interface Event {
  package: string;
  eventName: string;
  eventDescription: string;
};

async function main() {
  if (!IN) console.error('Error in config: `IN` not set');
  if (!OUT) console.error('Error in config: `OUT` not set');
  if (!IN || !OUT) return;
  if (IN.endsWith('/')) IN = IN.substring(0, IN.length - 1);
  if (OUT.endsWith('/')) OUT = OUT.substring(0, OUT.length - 1);

  const templateDirectoryPath = `${__dirname}/../../${TEMPLATE_DIRECTORY}`;

  // Collect all schemas and generate Go code from them using the
  // quicktype-wrapper package
  const schemasAndGenFiles = await qt.getJSONSchemasAndGenFiles(IN, LANGUAGE);
  const allEvents: Event[] = [];
  schemasAndGenFiles.map(([schema, genFile]: [any, string]) => {
    // Make the directory for all events in the schema
    let pkg = schema['$id'];
    pkg = pkg.replace('google.events.', '');
    const pkgPath = pkg.replace(/\./g, '/');
    mkdirp.sync(`${OUT}/${pkgPath}`);
    const eventName = schema.name;

    // Update the package statement
    const goPkg = schema['goPackage'];
    const goPkgStmt = `package ${goPkg}`;
    genFile = genFile.replace('package main', goPkgStmt);

    // Collect all events in the schema
    const eventDescription = schema.description;
    allEvents.push({
      package: goPkg,
      eventName: eventName,
      eventDescription: eventDescription,
    });

    // Write the patched generated code
    fs.writeFileSync(`${OUT}/${pkgPath}/${eventName}.go`, genFile);
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
