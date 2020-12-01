# Tools

Different tools for the Google CloudEvent – Go repository.

## Gen

The `gen.sh` script generates this library's source code.
To run this script, navigate to the root directory and run `tools/gen.sh`

This will generate golang files using the [`qt` tool](https://github.com/googleapis/google-cloudevents/tree/master/tools/quicktype-wrapper).

The generator will also run a postgen step, adding unmarshalling functions.

## Docs

To update the main README, run the `npm run docs` command in this directory.
This generates the documentation using the Google CloudEvent JSON schema catalog.
