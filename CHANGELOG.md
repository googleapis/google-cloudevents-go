# Changelog

## [0.5.0](https://github.com/googleapis/google-cloudevents-go/compare/v0.4.0...v0.5.0) (2023-05-25)


### Features

* Run the code generator (08ff8b4) ([#162](https://github.com/googleapis/google-cloudevents-go/issues/162)) ([d6ae4ff](https://github.com/googleapis/google-cloudevents-go/commit/d6ae4ffe75b73207a5f3fa9e73152236e6f4680c))


### Bug Fixes

* **deps:** Update google.golang.org/genproto digest to daa745c ([#163](https://github.com/googleapis/google-cloudevents-go/issues/163)) ([1c823a8](https://github.com/googleapis/google-cloudevents-go/commit/1c823a8929478cdeb99a2f43f106d277eafb7b94))
* **deps:** Update module google.golang.org/protobuf to v1.30.0 ([#164](https://github.com/googleapis/google-cloudevents-go/issues/164)) ([e11c4a1](https://github.com/googleapis/google-cloudevents-go/commit/e11c4a1ec21999cbdb2e32a149fe7743246f5d65))

## [0.4.0](https://github.com/googleapis/google-cloudevents-go/compare/v0.3.0...v0.4.0) (2023-03-30)


### Features

* Add support for new Event Type: DataFlow v1beta3 ([8504df3](https://github.com/googleapis/google-cloudevents-go/commit/8504df35ac0d25d907f960019cee432e60653d0d))


### Bug Fixes

* Support event type pre-release versions ([#160](https://github.com/googleapis/google-cloudevents-go/issues/160)) ([58aab52](https://github.com/googleapis/google-cloudevents-go/commit/58aab5260f1a2af84591e0dbc524582f685f6e5d))

## [0.3.0](https://github.com/googleapis/google-cloudevents-go/compare/v0.2.3...v0.3.0) (2023-01-31)


### ⚠ BREAKING CHANGES

* clear json-schema based version of the Go library ([#117](https://github.com/googleapis/google-cloudevents-go/issues/117))

### Features

* Generate type definition package godoc ([#142](https://github.com/googleapis/google-cloudevents-go/issues/142)) ([b9e2108](https://github.com/googleapis/google-cloudevents-go/commit/b9e2108ebc539fe8d2565da60b70c1a3f1478ee7))
* Run the code generator ([#129](https://github.com/googleapis/google-cloudevents-go/issues/129)) ([4d9f2db](https://github.com/googleapis/google-cloudevents-go/commit/4d9f2db181810f528ba351d8c2b3d758c4c93a0b))
* Run the code generator ([#148](https://github.com/googleapis/google-cloudevents-go/issues/148)) ([322c786](https://github.com/googleapis/google-cloudevents-go/commit/322c786abef2dfa47cce9a2d082ed83e47f5b10e))
* Tooling to generate types from protos ([#115](https://github.com/googleapis/google-cloudevents-go/issues/115)) ([82e971b](https://github.com/googleapis/google-cloudevents-go/commit/82e971b95f34e445d1eb51ff202be9441ec4c906))


### Bug Fixes

* Correct testdata lookup for deeply nested types ([#143](https://github.com/googleapis/google-cloudevents-go/issues/143)) ([797a7c1](https://github.com/googleapis/google-cloudevents-go/commit/797a7c1e4bc19f556367330e213a7da3aa95aae7))
* **generator:** Cron schedule in timezone UTC ([#131](https://github.com/googleapis/google-cloudevents-go/issues/131)) ([9f61d6d](https://github.com/googleapis/google-cloudevents-go/commit/9f61d6d0d13bc36246ae7df19258d99247aed0ee))
* **generator:** Github action source-version syntax error ([#127](https://github.com/googleapis/google-cloudevents-go/issues/127)) ([3c2d689](https://github.com/googleapis/google-cloudevents-go/commit/3c2d6895c4043f1043ef4de191b354859858230a))
* **generator:** Use side-by-side repositories ([#130](https://github.com/googleapis/google-cloudevents-go/issues/130)) ([521da22](https://github.com/googleapis/google-cloudevents-go/commit/521da225b7e85d6901bdc316b9542d1060aad29e))
* Skip unsupported pubsub event type ([#140](https://github.com/googleapis/google-cloudevents-go/issues/140)) ([6c4d4e4](https://github.com/googleapis/google-cloudevents-go/commit/6c4d4e4c144439c269a66fd86e076673db7041cd))
* **testing:** Generator workflow syntax error ([#151](https://github.com/googleapis/google-cloudevents-go/issues/151)) ([ef63a63](https://github.com/googleapis/google-cloudevents-go/commit/ef63a63ea939a736fe1d92597826e3abd01565a6))
* **testing:** Sort test cases per data package ([#156](https://github.com/googleapis/google-cloudevents-go/issues/156)) ([030cf6d](https://github.com/googleapis/google-cloudevents-go/commit/030cf6dced4d3ac2748a32e760f53f17fffd698e))


### Code Refactoring

* Clear json-schema based version of the Go library ([#117](https://github.com/googleapis/google-cloudevents-go/issues/117)) ([ee45e63](https://github.com/googleapis/google-cloudevents-go/commit/ee45e636b4393792ab0fbccb237c35081fc87e18))

## [0.2.3](https://github.com/googleapis/google-cloudevents-go/compare/v0.2.2...v0.2.3) (2022-07-14)


### Bug Fixes

* **deps:** update dependency node-fetch to v3.2.8 ([#106](https://github.com/googleapis/google-cloudevents-go/issues/106)) ([f989223](https://github.com/googleapis/google-cloudevents-go/commit/f989223899b6416c31944e47de5ba96039a8f1be))

### [0.2.2](https://github.com/googleapis/google-cloudevents-go/compare/v0.2.1...v0.2.2) (2022-05-27)


### Bug Fixes

* **deps:** update dependency node-fetch to v3.2.4 ([#97](https://github.com/googleapis/google-cloudevents-go/issues/97)) ([a4f4dd4](https://github.com/googleapis/google-cloudevents-go/commit/a4f4dd423f2ff502647efeb52fb6b5cb9b93c19e))

### [0.2.1](https://github.com/googleapis/google-cloudevents-go/compare/v0.2.0...v0.2.1) (2022-05-18)


### Bug Fixes

* **deps:** update dependency pascalcase to v2 ([#102](https://github.com/googleapis/google-cloudevents-go/issues/102)) ([868b08c](https://github.com/googleapis/google-cloudevents-go/commit/868b08cd730a504d4a6c20ad5e3e25aaf14f9b2a))

## [0.2.0](https://www.github.com/googleapis/google-cloudevents-go/compare/v0.1.1...v0.2.0) (2021-10-15)


### Features

* run generator (10-15-2021) ([#85](https://www.github.com/googleapis/google-cloudevents-go/issues/85)) ([0f3150c](https://www.github.com/googleapis/google-cloudevents-go/commit/0f3150c20752d8f66341b2b32892e3847b571d44))

### [0.1.1](https://www.github.com/googleapis/google-cloudevents-go/compare/v0.1.0...v0.1.1) (2021-10-01)


### Bug Fixes

* **deps:** update dependency node-fetch to v2.6.5 ([#79](https://www.github.com/googleapis/google-cloudevents-go/issues/79)) ([9540d8c](https://www.github.com/googleapis/google-cloudevents-go/commit/9540d8c68b0d26c8b43b566a2feacb4cee3d9a9b))
