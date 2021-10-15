// Copyright 2021 Google LLC
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

package testlab

// TestMatrixEventData: The data within all Firebase test matrix events.
type TestMatrixEventData struct {
	ClientInfo           *ClientInfo     `json:"clientInfo,omitempty"`           // Information provided by the client that created the test matrix.
	CreateTime           *string         `json:"createTime,omitempty"`           // Time the test matrix was created.
	InvalidMatrixDetails *string         `json:"invalidMatrixDetails,omitempty"` // Code that describes why the test matrix is considered invalid. Only set for; matrices in the INVALID state.
	OutcomeSummary       *OutcomeSummary `json:"outcomeSummary,omitempty"`       // Outcome summary of the test matrix.
	ResultStorage        *ResultStorage  `json:"resultStorage,omitempty"`        // Locations where test results are stored.
	State                *State          `json:"state,omitempty"`                // State of the test matrix.
}

// ClientInfo: Information provided by the client that created the test matrix.
//
// Information about the client which invoked the test.
type ClientInfo struct {
	Client  *string           `json:"client,omitempty"`  // Client name, such as "gcloud".
	Details map[string]string `json:"details,omitempty"` // Map of detailed information about the client.
}

// ResultStorage: Locations where test results are stored.
type ResultStorage struct {
	GcsPath              *string `json:"gcsPath,omitempty"`              // Location in Google Cloud Storage where test results are written to.; In the form "gs://bucket/path/to/somewhere".
	ResultsURI           *string `json:"resultsUri,omitempty"`           // URI to the test results in the Firebase Web Console.
	ToolResultsExecution *string `json:"toolResultsExecution,omitempty"` // Tool Results execution resource containing test results. Format is; `projects/{project_id}/histories/{history_id}/executions/{execution_id}`.; Optional, can be omitted in erroneous test states.; See https://firebase.google.com/docs/test-lab/reference/toolresults/rest; for more information.
	ToolResultsHistory   *string `json:"toolResultsHistory,omitempty"`   // Tool Results history resource containing test results. Format is; `projects/{project_id}/histories/{history_id}`.; See https://firebase.google.com/docs/test-lab/reference/toolresults/rest; for more information.
}

// Outcome summary of the test matrix.
type OutcomeSummary string

const (
	Failure                   OutcomeSummary = "FAILURE"
	Inconclusive              OutcomeSummary = "INCONCLUSIVE"
	OutcomeSummaryUnspecified OutcomeSummary = "OUTCOME_SUMMARY_UNSPECIFIED"
	Skipped                   OutcomeSummary = "SKIPPED"
	Success                   OutcomeSummary = "SUCCESS"
)

// State of the test matrix.
type State string

const (
	Error                State = "ERROR"
	Finished             State = "FINISHED"
	Invalid              State = "INVALID"
	Pending              State = "PENDING"
	TestStateUnspecified State = "TEST_STATE_UNSPECIFIED"
	Validating           State = "VALIDATING"
)
