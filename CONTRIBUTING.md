# How to Contribute

We'd love to accept your patches and contributions to this project. There are
just a few small guidelines you need to follow.

## Contributor License Agreement

Contributions to this project must be accompanied by a Contributor License
Agreement. You (or your employer) retain the copyright to your contribution;
this simply gives us permission to use and redistribute your contributions as
part of the project. Head over to <https://cla.developers.google.com/> to see
your current agreements on file or to sign a new one.

You generally only need to submit a CLA once, so if you've already submitted one
(even if it was for a different project), you probably don't need to do it
again.

## Code reviews

All submissions, including submissions by project members, require review. We
use GitHub pull requests for this purpose. Consult
[GitHub Help](https://help.github.com/articles/about-pull-requests/) for more
information on using pull requests.

## Community Guidelines

This project follows [Google's Open Source Community
Guidelines](https://opensource.google/conduct/).

## Generating the Library

```sh
git clone https://github.com/googleapis/google-cloudevents-go
cd google-cloudevents-go
sh ./tools/setup-generator.sh
export GENERATE_DATA_SOURCE=tmp/google-cloudevents
export GENERATE_PROTOC_PATH=tmp/protobuf/bin/protoc
sh ./generate-code.sh
```

### Generating the Library (Locally modified protos)

The "data source" for code generation is a collection of protos maintained in
[googleapis/google-cloudevets](https://github.com/googleapis/google-cloudevents).

If you have a local copy of this repository, such as for trying modifications to
those protos, you can use these instructions to use that copy instead of
retrieving a new clone.

```sh
cd /path/to/shared/repositories
git clone https://github.com/googleapis/google-cloudevents --depth 1

cd path/to/the/project
git clone https://github.com/googleapis/google-cloudevents-go

# Configure this before running the setup script to reuse the repository.
export GENERATE_DATA_SOURCE=/path/to/shared/repositories/google-cloudevents
sh ./tools/setup-generator.sh
export GENERATE_PROTOC_PATH=tmp/protobuf/bin/protoc
sh ./generate-code.sh
```

A similar "skip" does not exist for the protobuf library.
