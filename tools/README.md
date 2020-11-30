# Tools

Different tool

# Tools

Different tools for the Google CloudEvent – Go repository.

## Gen

The `gen.sh` script generates this library. To run this script, navigate to the
root directory and run `tools/gen.sh`

This will generate golang files using the [`qt` tool](https://github.com/googleapis/google-cloudevents/tree/master/tools/quicktype-wrapper).

The generator will also run a postgen step, adding unmarshalling functions.

## Docs

The `npm run docs` command generates the documentation for the main README using
the JSON schema catalog.