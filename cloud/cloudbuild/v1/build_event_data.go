// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cloudbuild

import "encoding/json"

// Build event data for Google Cloud Platform API operations.
type BuildEventData struct {
	Artifacts        *Artifacts             `json:"artifacts,omitempty"`       // Artifacts produced by the build that should be uploaded upon; successful completion of all build steps.
	BuildTriggerID   *string                `json:"buildTriggerId,omitempty"`  // The ID of the `BuildTrigger` that triggered this build, if it; was triggered automatically.
	CreateTime       *string                `json:"createTime,omitempty"`      // Time at which the request to create the build was received.
	FinishTime       *string                `json:"finishTime,omitempty"`      // Time at which execution of the build was finished.; ; The difference between finish_time and start_time is the duration of the; build's execution.
	ID               *string                `json:"id,omitempty"`              // Unique identifier of the build.
	Images           []string               `json:"images,omitempty"`          // A list of images to be pushed upon the successful completion of all build; steps.; ; The images are pushed using the builder service account's credentials.; ; The digests of the pushed images will be stored in the `Build` resource's; results field.; ; If any of the images fail to be pushed, the build status is marked; `FAILURE`.
	LogsBucket       *string                `json:"logsBucket,omitempty"`      // Google Cloud Storage bucket where logs should be written (see; [Bucket Name; Requirements](https://cloud.google.com/storage/docs/bucket-naming#requirements)).; Logs file names will be of the format `${logs_bucket}/log-${build_id}.txt`.
	LogURL           *string                `json:"logUrl,omitempty"`          // URL to logs for this build in Google Cloud Console.
	Options          *Options               `json:"options,omitempty"`         // Special options for this build.
	ProjectID        *string                `json:"projectId,omitempty"`       // ID of the project.
	QueueTTL         *QueueTTL              `json:"queueTtl,omitempty"`        // TTL in queue for this build. If provided and the build is enqueued longer; than this value, the build will expire and the build status will be; `EXPIRED`.; ; The TTL starts ticking from create_time.
	Results          *Results               `json:"results,omitempty"`         // Results of the build.
	Secrets          []Secret               `json:"secrets,omitempty"`         // Secrets to decrypt using Cloud Key Management Service.
	Source           *Source                `json:"source,omitempty"`          // The location of the source files to build.
	SourceProvenance *SourceProvenance      `json:"sourceProvenance,omitempty"`// A permanent fixed identifier for source.
	StartTime        *string                `json:"startTime,omitempty"`       // Time at which execution of the build was started.
	Status           *BuildEventDataStatus  `json:"status"`                    // Status of the build.
	StatusDetail     *string                `json:"statusDetail,omitempty"`    // Customer-readable message about the current status.
	Steps            []Step                 `json:"steps,omitempty"`           // The operations to be performed on the workspace.
	Substitutions    map[string]string      `json:"substitutions,omitempty"`   // Substitutions data for `Build` resource.
	Tags             []string               `json:"tags,omitempty"`            // Tags for annotation of a `Build`. These are not docker tags.
	Timeout          *BuildEventDataTimeout `json:"timeout,omitempty"`         // Amount of time that this build should be allowed to run, to second; granularity. If this amount of time elapses, work on the build will cease; and the build status will be `TIMEOUT`.
	Timing           map[string]TimeSpan    `json:"timing,omitempty"`          // Stores timing information for phases of the build. Valid keys; are:; ; * BUILD: time to execute all build steps; * PUSH: time to push all specified images.; * FETCHSOURCE: time to fetch source.; ; If the build does not specify source or images,; these keys will not be included.
}

// Artifacts produced by the build that should be uploaded upon
// successful completion of all build steps.
type Artifacts struct {
	Images  []string `json:"images,omitempty"` // A list of images to be pushed upon the successful completion of all build; steps.; ; The images will be pushed using the builder service account's credentials.; ; The digests of the pushed images will be stored in the Build resource's; results field.; ; If any of the images fail to be pushed, the build is marked FAILURE.
	Objects *Objects `json:"objects,omitempty"`// A list of objects to be uploaded to Cloud Storage upon successful; completion of all build steps.; ; Files in the workspace matching specified paths globs will be uploaded to; the specified Cloud Storage location using the builder service account's; credentials.; ; The location and generation of the uploaded objects will be stored in the; Build resource's results field.; ; If any objects fail to be pushed, the build is marked FAILURE.
}

// A list of objects to be uploaded to Cloud Storage upon successful
// completion of all build steps.
//
// Files in the workspace matching specified paths globs will be uploaded to
// the specified Cloud Storage location using the builder service account's
// credentials.
//
// The location and generation of the uploaded objects will be stored in the
// Build resource's results field.
//
// If any objects fail to be pushed, the build is marked FAILURE.
type Objects struct {
	Location *string        `json:"location,omitempty"`// Cloud Storage bucket and optional object path, in the form; "gs://bucket/path/to/somewhere/". (see [Bucket Name; Requirements](https://cloud.google.com/storage/docs/bucket-naming#requirements)).; ; Files in the workspace matching any path pattern will be uploaded to; Cloud Storage with this location as a prefix.
	Paths    []string       `json:"paths,omitempty"`   // Path globs used to match files in the build's workspace.
	Timing   *ObjectsTiming `json:"timing,omitempty"`  // Stores timing information for pushing all artifact objects.
}

// Stores timing information for pushing all artifact objects.
//
// Start and end times for a build execution phase.
type ObjectsTiming struct {
	EndTime   *string `json:"endTime,omitempty"`  // End of time span.
	StartTime *string `json:"startTime,omitempty"`// Start of time span.
}

// Special options for this build.
type Options struct {
	DiskSizeGB            *int64                        `json:"diskSizeGb,omitempty"`          // Requested disk size for the VM that runs the build. Note that this is *NOT*; "disk free"; some of the space will be used by the operating system and; build utilities. Also note that this is the minimum disk size that will be; allocated for the build -- the build may run with a larger disk than; requested. At present, the maximum disk size is 1000GB; builds that request; more than the maximum are rejected with an error.
	Env                   []string                      `json:"env,omitempty"`                 // A list of global environment variable definitions that will exist for all; build steps in this build. If a variable is defined in both globally and in; a build step, the variable will use the build step value.; ; The elements are of the form "KEY=VALUE" for the environment variable "KEY"; being given the value "VALUE".
	Logging               *Logging                      `json:"logging"`                       // Option to specify the logging mode, which determines where the logs are; stored.
	LogStreamingOption    *LogStreamingOption           `json:"logStreamingOption"`            // Option to define build log streaming behavior to Google Cloud; Storage.
	MachineType           *MachineType                  `json:"machineType"`                   // Compute Engine machine type on which to run the build.
	RequestedVerifyOption *RequestedVerifyOption        `json:"requestedVerifyOption"`         // Requested verifiability options.
	SecretEnv             []string                      `json:"secretEnv,omitempty"`           // A list of global environment variables, which are encrypted using a Cloud; Key Management Service crypto key. These values must be specified in the; build's `Secret`. These variables will be available to all build steps; in this build.
	SourceProvenanceHash  []SourceProvenanceHashElement `json:"sourceProvenanceHash,omitempty"`// Requested hash for SourceProvenance.
	SubstitutionOption    *SubstitutionOption           `json:"substitutionOption"`            // Option to specify behavior when there is an error in the substitution; checks.
	Volumes               []Volume                      `json:"volumes,omitempty"`             // Global list of volumes to mount for ALL build steps; ; Each volume is created as an empty volume prior to starting the build; process. Upon completion of the build, volumes and their contents are; discarded. Global volume names and paths cannot conflict with the volumes; defined a build step.; ; Using a global volume in a build with only one step is not valid as; it is indicative of a build request with an incorrect configuration.
	WorkerPool            *string                       `json:"workerPool,omitempty"`          // Option to specify a `WorkerPool` for the build.; Format: projects/{project}/locations/{location}/workerPools/{workerPool}
}

// Volume describes a Docker container volume which is mounted into build steps
// in order to persist files across build step execution.
type Volume struct {
	Name *string `json:"name,omitempty"`// Name of the volume to mount.; ; Volume names must be unique per build step and must be valid names for; Docker volumes. Each named volume must be used by at least two build steps.
	Path *string `json:"path,omitempty"`// Path at which to mount the volume.; ; Paths must be absolute and cannot conflict with other volume paths on the; same build step or with certain reserved volume paths.
}

// TTL in queue for this build. If provided and the build is enqueued longer
// than this value, the build will expire and the build status will be
// `EXPIRED`.
//
// The TTL starts ticking from create_time.
type QueueTTL struct {
	Nanos   *int64 `json:"nanos,omitempty"`  // Signed fractions of a second at nanosecond resolution of the span; of time. Durations less than one second are represented with a 0; `seconds` field and a positive or negative `nanos` field. For durations; of one second or more, a non-zero value for the `nanos` field must be; of the same sign as the `seconds` field. Must be from -999,999,999; to +999,999,999 inclusive.
	Seconds *int64 `json:"seconds,omitempty"`// Signed seconds of the span of time. Must be from -315,576,000,000; to +315,576,000,000 inclusive. Note: these bounds are computed from:; 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years
}

// Results of the build.
type Results struct {
	ArtifactManifest *string         `json:"artifactManifest,omitempty"`// Path to the artifact manifest. Only populated when artifacts are uploaded.
	ArtifactTiming   *ArtifactTiming `json:"artifactTiming,omitempty"`  // Time to push all non-container artifacts.
	BuildStepImages  []string        `json:"buildStepImages,omitempty"` // List of build step digests, in the order corresponding to build step; indices.
	BuildStepOutputs []string        `json:"buildStepOutputs,omitempty"`// List of build step outputs, produced by builder images, in the order; corresponding to build step indices.; ; [Cloud Builders](https://cloud.google.com/cloud-build/docs/cloud-builders); can produce this output by writing to `$BUILDER_OUTPUT/output`.; Only the first 4KB of data is stored.
	Images           []Image         `json:"images,omitempty"`          // Container images that were built as a part of the build.
	NumArtifacts     *int64          `json:"numArtifacts,omitempty"`    // Number of artifacts uploaded. Only populated when artifacts are uploaded.
}

// Time to push all non-container artifacts.
//
// Stores timing information for pushing all artifact objects.
//
// Start and end times for a build execution phase.
type ArtifactTiming struct {
	EndTime   *string `json:"endTime,omitempty"`  // End of time span.
	StartTime *string `json:"startTime,omitempty"`// Start of time span.
}

// An image built by the pipeline.
type Image struct {
	Digest     *string     `json:"digest,omitempty"`    // Docker Registry 2.0 digest.
	Name       *string     `json:"name,omitempty"`      // Name used to push the container image to Google Container Registry, as; presented to `docker push`.
	PushTiming *PushTiming `json:"pushTiming,omitempty"`// Stores timing information for pushing the specified image.
}

// Stores timing information for pushing the specified image.
//
// Stores timing information for pushing all artifact objects.
//
// Start and end times for a build execution phase.
type PushTiming struct {
	EndTime   *string `json:"endTime,omitempty"`  // End of time span.
	StartTime *string `json:"startTime,omitempty"`// Start of time span.
}

// Pairs a set of secret environment variables containing encrypted
// values with the Cloud KMS key to use to decrypt the value.
type Secret struct {
	KmsKeyName *string           `json:"kmsKeyName,omitempty"`// Cloud KMS key name to use to decrypt these envs.
	SecretEnv  map[string]string `json:"secretEnv,omitempty"` // Map of environment variable name to its encrypted value.; ; Secret environment variables must be unique across all of a build's; secrets, and must be used by at least one build step. Values can be at most; 64 KB in size. There can be at most 100 secret values across all of a; build's secrets.
}

// The location of the source files to build.
type Source struct {
	RepoSource    *RepoSourceClass    `json:"repoSource,omitempty"`   // If provided, get the source from this location in a Cloud Source; Repository.
	StorageSource *StorageSourceClass `json:"storageSource,omitempty"`// If provided, get the source from this location in Google Cloud Storage.
}

// If provided, get the source from this location in a Cloud Source
// Repository.
//
// Location of the source in a Google Cloud Source Repository.
type RepoSourceClass struct {
	BranchName    *string           `json:"branchName,omitempty"`   // Regex matching branches to build.; ; The syntax of the regular expressions accepted is the syntax accepted by; RE2 and described at https://github.com/google/re2/wiki/Syntax
	CommitSHA     *string           `json:"commitSha,omitempty"`    // Explicit commit SHA to build.
	Dir           *string           `json:"dir,omitempty"`          // Directory, relative to the source root, in which to run the build.; ; This must be a relative path. If a step's `dir` is specified and is an; absolute path, this value is ignored for that step's execution.
	InvertRegex   *bool             `json:"invertRegex,omitempty"`  // Only trigger a build if the revision regex does NOT match the revision; regex.
	ProjectID     *string           `json:"projectId,omitempty"`    // ID of the project that owns the Cloud Source Repository.
	RepoName      *string           `json:"repoName,omitempty"`     // Name of the Cloud Source Repository.
	Substitutions map[string]string `json:"substitutions,omitempty"`// Substitutions to use in a triggered build.; Should only be used with RunBuildTrigger
	TagName       *string           `json:"tagName,omitempty"`      // Regex matching tags to build.; ; The syntax of the regular expressions accepted is the syntax accepted by; RE2 and described at https://github.com/google/re2/wiki/Syntax
}

// If provided, get the source from this location in Google Cloud Storage.
//
// Location of the source in an archive file in Google Cloud Storage.
type StorageSourceClass struct {
	Bucket     *string `json:"bucket,omitempty"`    // Google Cloud Storage bucket containing the source (see; [Bucket Name; Requirements](https://cloud.google.com/storage/docs/bucket-naming#requirements)).
	Generation *int64  `json:"generation,omitempty"`// Google Cloud Storage generation for the object. If the generation is; omitted, the latest generation will be used.
	Object     *string `json:"object,omitempty"`    // Google Cloud Storage object containing the source.
}

// A permanent fixed identifier for source.
type SourceProvenance struct {
	FileHashes            map[string]FileHashValue    `json:"fileHashes,omitempty"`           // Hash(es) of the build source, which can be used to verify that; the original source integrity was maintained in the build. Note that; `FileHashes` will only be populated if `BuildOptions` has requested a; `SourceProvenanceHash`.; ; The keys to this map are file paths used as build source and the values; contain the hash values for those files.; ; If the build source came in a single package such as a gzipped tarfile; (`.tar.gz`), the `FileHash` will be for the single path to that file.
	ResolvedRepoSource    *ResolvedRepoSourceClass    `json:"resolvedRepoSource,omitempty"`   // A copy of the build's `source.repo_source`, if exists, with any; revisions resolved.
	ResolvedStorageSource *ResolvedStorageSourceClass `json:"resolvedStorageSource,omitempty"`// A copy of the build's `source.storage_source`, if exists, with any; generations resolved.
}

type FileHashValue struct {
	FileHash []FileHashElement `json:"fileHash,omitempty"`// Collection of file hashes.
}

// Container message for hash values.
type FileHashElement struct {
	Type  *Type   `json:"type"`           // The type of hash that was performed.
	Value *string `json:"value,omitempty"`// The hash value.
}

// A copy of the build's `source.repo_source`, if exists, with any
// revisions resolved.
//
// If provided, get the source from this location in a Cloud Source
// Repository.
//
// Location of the source in a Google Cloud Source Repository.
type ResolvedRepoSourceClass struct {
	BranchName    *string           `json:"branchName,omitempty"`   // Regex matching branches to build.; ; The syntax of the regular expressions accepted is the syntax accepted by; RE2 and described at https://github.com/google/re2/wiki/Syntax
	CommitSHA     *string           `json:"commitSha,omitempty"`    // Explicit commit SHA to build.
	Dir           *string           `json:"dir,omitempty"`          // Directory, relative to the source root, in which to run the build.; ; This must be a relative path. If a step's `dir` is specified and is an; absolute path, this value is ignored for that step's execution.
	InvertRegex   *bool             `json:"invertRegex,omitempty"`  // Only trigger a build if the revision regex does NOT match the revision; regex.
	ProjectID     *string           `json:"projectId,omitempty"`    // ID of the project that owns the Cloud Source Repository.
	RepoName      *string           `json:"repoName,omitempty"`     // Name of the Cloud Source Repository.
	Substitutions map[string]string `json:"substitutions,omitempty"`// Substitutions to use in a triggered build.; Should only be used with RunBuildTrigger
	TagName       *string           `json:"tagName,omitempty"`      // Regex matching tags to build.; ; The syntax of the regular expressions accepted is the syntax accepted by; RE2 and described at https://github.com/google/re2/wiki/Syntax
}

// A copy of the build's `source.storage_source`, if exists, with any
// generations resolved.
//
// If provided, get the source from this location in Google Cloud Storage.
//
// Location of the source in an archive file in Google Cloud Storage.
type ResolvedStorageSourceClass struct {
	Bucket     *string `json:"bucket,omitempty"`    // Google Cloud Storage bucket containing the source (see; [Bucket Name; Requirements](https://cloud.google.com/storage/docs/bucket-naming#requirements)).
	Generation *int64  `json:"generation,omitempty"`// Google Cloud Storage generation for the object. If the generation is; omitted, the latest generation will be used.
	Object     *string `json:"object,omitempty"`    // Google Cloud Storage object containing the source.
}

// A step in the build pipeline.
type Step struct {
	Args       []string                     `json:"args,omitempty"`      // A list of arguments that will be presented to the step when it is started.; ; If the image used to run the step's container has an entrypoint, the `args`; are used as arguments to that entrypoint. If the image does not define; an entrypoint, the first element in args is used as the entrypoint,; and the remainder will be used as arguments.
	Dir        *string                      `json:"dir,omitempty"`       // Working directory to use when running this step's container.; ; If this value is a relative path, it is relative to the build's working; directory. If this value is absolute, it may be outside the build's working; directory, in which case the contents of the path may not be persisted; across build step executions, unless a `volume` for that path is specified.; ; If the build specifies a `RepoSource` with `dir` and a step with a `dir`,; which specifies an absolute path, the `RepoSource` `dir` is ignored for; the step's execution.
	Entrypoint *string                      `json:"entrypoint,omitempty"`// Entrypoint to be used instead of the build step image's default entrypoint.; If unset, the image's default entrypoint is used.
	Env        []string                     `json:"env,omitempty"`       // A list of environment variable definitions to be used when running a step.; ; The elements are of the form "KEY=VALUE" for the environment variable "KEY"; being given the value "VALUE".
	ID         *string                      `json:"id,omitempty"`        // Unique identifier for this build step, used in `wait_for` to; reference this build step as a dependency.
	Name       *string                      `json:"name,omitempty"`      // The name of the container image that will run this particular; build step.; ; If the image is available in the host's Docker daemon's cache, it; will be run directly. If not, the host will attempt to pull the image; first, using the builder service account's credentials if necessary.; ; The Docker daemon's cache will already have the latest versions of all of; the officially supported build steps; ; ([https://github.com/GoogleCloudPlatform/cloud-builders](https://github.com/GoogleCloudPlatform/cloud-builders)).; The Docker daemon will also have cached many of the layers for some popular; images, like "ubuntu", "debian", but they will be refreshed at the time you; attempt to use them.; ; If you built an image in a previous build step, it will be stored in the; host's Docker daemon's cache and is available to use as the name for a; later build step.
	PullTiming *PullTiming                  `json:"pullTiming,omitempty"`// Stores timing information for pulling this build step's; builder image only.
	SecretEnv  []string                     `json:"secretEnv,omitempty"` // A list of environment variables which are encrypted using a Cloud Key; Management Service crypto key. These values must be specified in the; build's `Secret`.
	Status     *SourceProvenanceHashElement `json:"status"`              // Status of the build step. At this time, build step status is; only updated on build completion; step status is not updated in real-time; as the build progresses.
	Timeout    *StepTimeout                 `json:"timeout,omitempty"`   // Time limit for executing this build step. If not defined, the step has no; time limit and will be allowed to continue to run until either it completes; or the build itself times out.
	Timing     *StepTiming                  `json:"timing,omitempty"`    // Stores timing information for executing this build step.
	Volumes    []Volume                     `json:"volumes,omitempty"`   // List of volumes to mount into the build step.; ; Each volume is created as an empty volume prior to execution of the; build step. Upon completion of the build, volumes and their contents are; discarded.; ; Using a named volume in only one step is not valid as it is indicative; of a build request with an incorrect configuration.
	WaitFor    []string                     `json:"waitFor,omitempty"`   // The ID(s) of the step(s) that this build step depends on.; This build step will not start until all the build steps in `wait_for`; have completed successfully. If `wait_for` is empty, this build step will; start when all previous build steps in the `Build.Steps` list have; completed successfully.
}

// Stores timing information for pulling this build step's
// builder image only.
//
// Stores timing information for pushing all artifact objects.
//
// Start and end times for a build execution phase.
type PullTiming struct {
	EndTime   *string `json:"endTime,omitempty"`  // End of time span.
	StartTime *string `json:"startTime,omitempty"`// Start of time span.
}

// Time limit for executing this build step. If not defined, the step has no
// time limit and will be allowed to continue to run until either it completes
// or the build itself times out.
type StepTimeout struct {
	Nanos   *int64 `json:"nanos,omitempty"`  // Signed fractions of a second at nanosecond resolution of the span; of time. Durations less than one second are represented with a 0; `seconds` field and a positive or negative `nanos` field. For durations; of one second or more, a non-zero value for the `nanos` field must be; of the same sign as the `seconds` field. Must be from -999,999,999; to +999,999,999 inclusive.
	Seconds *int64 `json:"seconds,omitempty"`// Signed seconds of the span of time. Must be from -315,576,000,000; to +315,576,000,000 inclusive. Note: these bounds are computed from:; 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years
}

// Stores timing information for executing this build step.
//
// Stores timing information for pushing all artifact objects.
//
// Start and end times for a build execution phase.
type StepTiming struct {
	EndTime   *string `json:"endTime,omitempty"`  // End of time span.
	StartTime *string `json:"startTime,omitempty"`// Start of time span.
}

// Amount of time that this build should be allowed to run, to second
// granularity. If this amount of time elapses, work on the build will cease
// and the build status will be `TIMEOUT`.
type BuildEventDataTimeout struct {
	Nanos   *int64 `json:"nanos,omitempty"`  // Signed fractions of a second at nanosecond resolution of the span; of time. Durations less than one second are represented with a 0; `seconds` field and a positive or negative `nanos` field. For durations; of one second or more, a non-zero value for the `nanos` field must be; of the same sign as the `seconds` field. Must be from -999,999,999; to +999,999,999 inclusive.
	Seconds *int64 `json:"seconds,omitempty"`// Signed seconds of the span of time. Must be from -315,576,000,000; to +315,576,000,000 inclusive. Note: these bounds are computed from:; 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years
}

// Stores timing information for pushing all artifact objects.
//
// Start and end times for a build execution phase.
type TimeSpan struct {
	EndTime   *string `json:"endTime,omitempty"`  // End of time span.
	StartTime *string `json:"startTime,omitempty"`// Start of time span.
}

type LogStreamingOptionEnum string
const (
	StreamDefault LogStreamingOptionEnum = "STREAM_DEFAULT"
	StreamOff LogStreamingOptionEnum = "STREAM_OFF"
	StreamOn LogStreamingOptionEnum = "STREAM_ON"
)

type LoggingEnum string
const (
	GcsOnly LoggingEnum = "GCS_ONLY"
	Legacy LoggingEnum = "LEGACY"
	LoggingUnspecified LoggingEnum = "LOGGING_UNSPECIFIED"
)

type MachineTypeEnum string
const (
	N1Highcpu32 MachineTypeEnum = "N1_HIGHCPU_32"
	N1Highcpu8 MachineTypeEnum = "N1_HIGHCPU_8"
	Unspecified MachineTypeEnum = "UNSPECIFIED"
)

type RequestedVerifyOptionEnum string
const (
	NotVerified RequestedVerifyOptionEnum = "NOT_VERIFIED"
	Verified RequestedVerifyOptionEnum = "VERIFIED"
)

type SubstitutionOptionEnum string
const (
	AllowLoose SubstitutionOptionEnum = "ALLOW_LOOSE"
	MustMatch SubstitutionOptionEnum = "MUST_MATCH"
)

type TypeEnum string
const (
	Md5 TypeEnum = "MD5"
	None TypeEnum = "NONE"
	Sha256 TypeEnum = "SHA256"
)

type StatusEnum string
const (
	Cancelled StatusEnum = "CANCELLED"
	Expired StatusEnum = "EXPIRED"
	Failure StatusEnum = "FAILURE"
	InternalError StatusEnum = "INTERNAL_ERROR"
	Queued StatusEnum = "QUEUED"
	StatusUnknown StatusEnum = "STATUS_UNKNOWN"
	Success StatusEnum = "SUCCESS"
	Timeout StatusEnum = "TIMEOUT"
	Working StatusEnum = "WORKING"
)

// Option to define build log streaming behavior to Google Cloud
// Storage.
type LogStreamingOption struct {
	Enum    *LogStreamingOptionEnum
	Integer *int64
}

// Option to specify the logging mode, which determines where the logs are
// stored.
type Logging struct {
	Enum    *LoggingEnum
	Integer *int64
}

// Compute Engine machine type on which to run the build.
type MachineType struct {
	Enum    *MachineTypeEnum
	Integer *int64
}

// Requested verifiability options.
type RequestedVerifyOption struct {
	Enum    *RequestedVerifyOptionEnum
	Integer *int64
}

type SourceProvenanceHashElement struct {
	Integer *int64
	String  *string
}

// Option to specify behavior when there is an error in the substitution
// checks.
type SubstitutionOption struct {
	Enum    *SubstitutionOptionEnum
	Integer *int64
}

// The type of hash that was performed.
type Type struct {
	Enum    *TypeEnum
	Integer *int64
}

// Status of the build.
type BuildEventDataStatus struct {
	Enum    *StatusEnum
	Integer *int64
}

func UnmarshalBuildEventData(data []byte) (BuildEventData, error) {
	var d BuildEventData
	err := json.Unmarshal(data, &d)
	return d, err
}

func (p *BuildEventData) MarshalBuildEventData() ([]byte, error) {
	return json.Marshal(p)
}