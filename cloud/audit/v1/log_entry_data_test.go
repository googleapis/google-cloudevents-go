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

package audit

import (
	"encoding/json"
	"testing"
)

// https://cloud.google.com/eventarc/docs/cloudevents
func TestMessagePublishedData(t *testing.T) {
	pubsub_data := []byte(`{
		"protoPayload": {
			"status": {},
			"authenticationInfo": {},
			"requestMetadata": {
				"callerIp": "...",
				"callerSuppliedUserAgent": "...",
				"requestAttributes": {
					"time": "YYYY-MM-DDThh:mm:ss.sZ",
					"auth": {}
				},
				"destinationAttributes": {}
			},
			"serviceName": "storage.googleapis.com",
			"methodName": "storage.objects.create",
			"authorizationInfo": [
				{
					"resource": "projects/_/buckets/MY_BUCKET/objects/MY_FILE.txt",
					"permission": "storage.objects.create",
					"granted": true,
					"resourceAttributes": {}
				},
				{
					"resource": "projects/_/buckets/MY_BUCKET/objects/MY_FILE.txt",
					"permission": "storage.objects.delete",
					"granted": true,
					"resourceAttributes": {}
				}
			],
			"resourceName": "projects/_/buckets/MY_PROJECT/objects/MY_FILE.txt",
			"serviceData": {
				"policyDelta": {
					"bindingDeltas": []
				},
				"@type": "type.googleapis.com/google.iam.v1.logging.AuditData"
			},
			"resourceLocation": {
				"currentLocations": [
					"us-central1"
				]
			}
		},
		"insertId": "-qklu1af11306y",
		"resource": {
			"type": "gcs_bucket",
			"labels": {
				"location": "us-central1",
				"bucket_name": "MY_BUCKET",
				"project_id": "MY_PROJECT"
			}
		},
		"timestamp": "2021-04-09T17:57:16.718915757Z",
		"severity": "INFO",
		"logName": "projects/MY_PROJECT/logs/cloudaudit.googleapis.com%2Fdata_access",
		"receiveTimestamp": "2021-04-09T17:57:17.748555474Z"
	}`)

	var e LogEntryData
	err := json.Unmarshal(pubsub_data, &e)
	if err != nil {
		panic(err)
	}

	if data := string(*e.Severity); data != "INFO" {
		t.Errorf("Expected %s, got %s", "INFO", data)
	}
}
