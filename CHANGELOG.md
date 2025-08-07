# Changelog

## [0.10.1](https://github.com/googleapis/google-cloudevents-go/compare/v0.10.0...v0.10.1) (2025-08-07)


### Bug Fixes

* **deps:** Update google.golang.org/genproto digest to f173205 ([#298](https://github.com/googleapis/google-cloudevents-go/issues/298)) ([727ebde](https://github.com/googleapis/google-cloudevents-go/commit/727ebde97a3863c735d60b3632f94a428dbb0ea8))

## [0.10.0](https://github.com/googleapis/google-cloudevents-go/compare/v0.9.0...v0.10.0) (2025-06-27)


### Features

* Add type definitions for API Hub events ([#274](https://github.com/googleapis/google-cloudevents-go/issues/274)) ([e3a3be6](https://github.com/googleapis/google-cloudevents-go/commit/e3a3be63702e01c6c8a898059fd9050126ac5a21))


### Bug Fixes

* **deps:** Update google.golang.org/genproto digest to 513f239 ([#293](https://github.com/googleapis/google-cloudevents-go/issues/293)) ([6d41f39](https://github.com/googleapis/google-cloudevents-go/commit/6d41f39d2f109146f0336b169e02dece462af8ec))
* **deps:** Update google.golang.org/genproto digest to b45e905 ([#278](https://github.com/googleapis/google-cloudevents-go/issues/278)) ([b4a7f18](https://github.com/googleapis/google-cloudevents-go/commit/b4a7f18dd0bd8d5f55938ba203db602c41279bba))
* **deps:** Update google.golang.org/genproto/googleapis/rpc digest to 513f239 ([#294](https://github.com/googleapis/google-cloudevents-go/issues/294)) ([80ad9e1](https://github.com/googleapis/google-cloudevents-go/commit/80ad9e1be3f57623046cabaf1c703901d115ce90))
* **deps:** Update google.golang.org/genproto/googleapis/rpc digest to b45e905 ([#279](https://github.com/googleapis/google-cloudevents-go/issues/279)) ([1f79af3](https://github.com/googleapis/google-cloudevents-go/commit/1f79af3dee5a1d9077474baec78a8aec8e81afdf))
* **deps:** Update module google.golang.org/protobuf to v1.36.6 ([#280](https://github.com/googleapis/google-cloudevents-go/issues/280)) ([b1faa7d](https://github.com/googleapis/google-cloudevents-go/commit/b1faa7d1515257cbcb14624394485cfbf1d3e727))

## [0.9.0](https://github.com/googleapis/google-cloudevents-go/compare/v0.8.0...v0.9.0) (2024-09-26)


### Features

* Add Cloud Deploy Event Types ([#204](https://github.com/googleapis/google-cloudevents-go/issues/204)) ([378fd48](https://github.com/googleapis/google-cloudevents-go/commit/378fd4852f729d7590fc8852cae1911dd71ba604))
* Add Cloud Scheduler created, updated, and deleted event types ([#206](https://github.com/googleapis/google-cloudevents-go/issues/206)) ([16d7643](https://github.com/googleapis/google-cloudevents-go/commit/16d7643c1a2c1229c1712a55571a812a1be5ad11))
* Run the code generator (a5f31a9) ([#269](https://github.com/googleapis/google-cloudevents-go/issues/269)) ([e23cb17](https://github.com/googleapis/google-cloudevents-go/commit/e23cb178f5a40f8f6cceb3e8208d38698d612c51))


### Bug Fixes

* **ci:** Fix go version min dependency for failing builds ([#267](https://github.com/googleapis/google-cloudevents-go/issues/267)) ([6f31753](https://github.com/googleapis/google-cloudevents-go/commit/6f31753828df4cecc5cf3746f8946c01282de9f4))
* **deps:** Update google.golang.org/genproto digest to 2c9e96a ([#227](https://github.com/googleapis/google-cloudevents-go/issues/227)) ([1b32592](https://github.com/googleapis/google-cloudevents-go/commit/1b32592c45bf5c4ffff97297beb25c803f8d0ec5))
* **deps:** Update google.golang.org/genproto/googleapis/rpc digest to 8af14fe ([#228](https://github.com/googleapis/google-cloudevents-go/issues/228)) ([2490ba9](https://github.com/googleapis/google-cloudevents-go/commit/2490ba97de64131e1a6e95ecd7230c6d0f308435))
* **deps:** Update google.golang.org/genproto/googleapis/rpc digest to 8c6c420 ([#203](https://github.com/googleapis/google-cloudevents-go/issues/203)) ([3530e19](https://github.com/googleapis/google-cloudevents-go/commit/3530e191618698dc3e9c74bfbdf32abb3c0d5841))
* **deps:** Update module google.golang.org/protobuf to v1.34.2 ([#229](https://github.com/googleapis/google-cloudevents-go/issues/229)) ([dfac9c3](https://github.com/googleapis/google-cloudevents-go/commit/dfac9c39a37991676e777fc3758456e82aed1dbd))
* **renovate:** Schedule syntax ([#271](https://github.com/googleapis/google-cloudevents-go/issues/271)) ([aa388e3](https://github.com/googleapis/google-cloudevents-go/commit/aa388e3994c3c965d43f06eb5c573d03b1d4d207))

## [0.8.0](https://github.com/googleapis/google-cloudevents-go/compare/v0.7.1...v0.8.0) (2024-03-29)


### Features

* Add event types for [Cloud Batch](https://cloud.google.com/eventarc/docs/reference/supported-events#batch) ([#197](https://github.com/googleapis/google-cloudevents-go/issues/197)) ([97b28ef](https://github.com/googleapis/google-cloudevents-go/commit/97b28ef36421b64a645b5f330549d15f5895ec48))
* Add event types for [Network Services](https://cloud.google.com/eventarc/docs/reference/supported-events#network-services) ([#197](https://github.com/googleapis/google-cloudevents-go/issues/197)) ([97b28ef](https://github.com/googleapis/google-cloudevents-go/commit/97b28ef36421b64a645b5f330549d15f5895ec48))
* Add event types for Speech-to-Text ([#197](https://github.com/googleapis/google-cloudevents-go/issues/197)) ([97b28ef](https://github.com/googleapis/google-cloudevents-go/commit/97b28ef36421b64a645b5f330549d15f5895ec48))
* Add google.cloud.datastore.entity.v1.*.withAuthContext event types ([#196](https://github.com/googleapis/google-cloudevents-go/issues/196)) ([e596ab7](https://github.com/googleapis/google-cloudevents-go/commit/e596ab744e99445a8a2c6b23d0cc59714207cf4f))


### Bug Fixes

* **deps:** Update google.golang.org/genproto digest to 995d672 ([#189](https://github.com/googleapis/google-cloudevents-go/issues/189)) ([40e253f](https://github.com/googleapis/google-cloudevents-go/commit/40e253fd80c5b5433916e0e21effe700fcdf423d))
* **deps:** Update google.golang.org/genproto/googleapis/rpc digest to 995d672 ([#190](https://github.com/googleapis/google-cloudevents-go/issues/190)) ([9291a1f](https://github.com/googleapis/google-cloudevents-go/commit/9291a1f14eef3a33cf99a64d284ce034d21b4696))
* **deps:** Update module google.golang.org/protobuf to v1.32.0 ([#191](https://github.com/googleapis/google-cloudevents-go/issues/191)) ([2cf5cb4](https://github.com/googleapis/google-cloudevents-go/commit/2cf5cb4415e862beeb1234243eade91ffae3573c))
* **deps:** Update module google.golang.org/protobuf to v1.33.0 [security] ([#200](https://github.com/googleapis/google-cloudevents-go/issues/200)) ([50b5201](https://github.com/googleapis/google-cloudevents-go/commit/50b5201af9aa8d10ad1650933ea6ecd37dc8fe5e))

## [0.7.1](https://github.com/googleapis/google-cloudevents-go/compare/v0.7.0...v0.7.1) (2023-11-02)


### Bug Fixes

* **deps:** Update google.golang.org/genproto digest to d783a09 ([#185](https://github.com/googleapis/google-cloudevents-go/issues/185)) ([4aa2d89](https://github.com/googleapis/google-cloudevents-go/commit/4aa2d89097e98f7de07f772980c6322575e4bb29))
* **deps:** Update google.golang.org/genproto/googleapis/rpc digest to d783a09 ([#186](https://github.com/googleapis/google-cloudevents-go/issues/186)) ([67d9c81](https://github.com/googleapis/google-cloudevents-go/commit/67d9c81912fd037f8297de6a7fb206396dac3cf9))

## [0.7.0](https://github.com/googleapis/google-cloudevents-go/compare/v0.6.0...v0.7.0) (2023-07-13)


### Features

* Run the code generator (0f27e75) ([#183](https://github.com/googleapis/google-cloudevents-go/issues/183)) ([a1dc2f9](https://github.com/googleapis/google-cloudevents-go/commit/a1dc2f9b451966a0a06830b60ea82b9bb51773d2))


### Bug Fixes

* **deps:** Update google.golang.org/genproto digest to ccb25ca ([#179](https://github.com/googleapis/google-cloudevents-go/issues/179)) ([e856a0b](https://github.com/googleapis/google-cloudevents-go/commit/e856a0b7e32fa2c685d77c883cdbbd11bea27095))
* **deps:** Update module google.golang.org/protobuf to v1.31.0 ([#180](https://github.com/googleapis/google-cloudevents-go/issues/180)) ([30d2bf4](https://github.com/googleapis/google-cloudevents-go/commit/30d2bf4f584832c2c3dec9adb69510e1386f2013))
* **docs:** Correct product names ([[#181](https://github.com/googleapis/google-cloudevents-go/issues/181)](https://github.com/googleapis/google-cloudevents-go/issues/181)) ([cb5d980](https://github.com/googleapis/google-cloudevents-go/commit/cb5d9801c1d10139a4bafaafbd58cc06a739873a))

## [0.6.0](https://github.com/googleapis/google-cloudevents-go/compare/v0.5.0...v0.6.0) (2023-06-27)


### Features

* Run the code generator (1300be8) ([#171](https://github.com/googleapis/google-cloudevents-go/issues/171)) ([74c44f2](https://github.com/googleapis/google-cloudevents-go/commit/74c44f20b0ce858b128c1e19b8d0436470ccd757))
* Run the code generator (45395fe) ([#173](https://github.com/googleapis/google-cloudevents-go/issues/173)) ([dd77933](https://github.com/googleapis/google-cloudevents-go/commit/dd779339a48877382e5eb816a50b890ccd9405c8))
* Run the code generator (45395fe) ([#176](https://github.com/googleapis/google-cloudevents-go/issues/176)) ([238d825](https://github.com/googleapis/google-cloudevents-go/commit/238d825a1c9fea635bc4cf60120c39b530f6a0b5))
* Run the code generator (45395fe) ([#177](https://github.com/googleapis/google-cloudevents-go/issues/177)) ([af89ff7](https://github.com/googleapis/google-cloudevents-go/commit/af89ff7873cf17caac21af1b2f9ea9b96a4e23ef))


### Bug Fixes

* Remove library version from test headers ([#175](https://github.com/googleapis/google-cloudevents-go/issues/175)) ([86fd204](https://github.com/googleapis/google-cloudevents-go/commit/86fd204b7a2d50b791cc2bae3133ecf6164b50fd))

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
