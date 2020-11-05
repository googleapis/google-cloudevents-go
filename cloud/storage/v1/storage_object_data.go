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
package storage

import "encoding/json"

// An object within Google Cloud Storage.
type StorageObjectData struct {
	Bucket                  *string             `json:"bucket,omitempty"`                 // The name of the bucket containing this object.
	CacheControl            *string             `json:"cacheControl,omitempty"`           // Cache-Control directive for the object data, matching; [https://tools.ietf.org/html/rfc7234#section-5.2"][RFC 7234 §5.2].
	ComponentCount          *int64              `json:"componentCount,omitempty"`         // Number of underlying components that make up this object. Components are; accumulated by compose operations.; Attempting to set this field will result in an error.
	ContentDisposition      *string             `json:"contentDisposition,omitempty"`     // Content-Disposition of the object data, matching; [https://tools.ietf.org/html/rfc6266][RFC 6266].
	ContentEncoding         *string             `json:"contentEncoding,omitempty"`        // Content-Encoding of the object data, matching; [https://tools.ietf.org/html/rfc7231#section-3.1.2.2][RFC 7231 §3.1.2.2]
	ContentLanguage         *string             `json:"contentLanguage,omitempty"`        // Content-Language of the object data, matching; [https://tools.ietf.org/html/rfc7231#section-3.1.3.2][RFC 7231 §3.1.3.2].
	ContentType             *string             `json:"contentType,omitempty"`            // Content-Type of the object data, matching; [https://tools.ietf.org/html/rfc7231#section-3.1.1.5][RFC 7231 §3.1.1.5].; If an object is stored without a Content-Type, it is served as; `application/octet-stream`.
	Crc32C                  *string             `json:"crc32C,omitempty"`                 // CRC32c checksum. For more information about using the CRC32c; checksum, see; [https://cloud.google.com/storage/docs/hashes-etags#_JSONAPI][Hashes and; ETags: Best Practices].
	CustomerEncryption      *CustomerEncryption `json:"customerEncryption,omitempty"`     // Metadata of customer-supplied encryption key, if the object is encrypted by; such a key.
	Etag                    *string             `json:"etag,omitempty"`                   // HTTP 1.1 Entity tag for the object. See; [https://tools.ietf.org/html/rfc7232#section-2.3][RFC 7232 §2.3].
	EventBasedHold          *bool               `json:"eventBasedHold,omitempty"`         // Whether an object is under event-based hold.
	Generation              *Generation         `json:"generation"`                       // The content generation of this object. Used for object versioning.; Attempting to set this field will result in an error.
	ID                      *string             `json:"id,omitempty"`                     // The ID of the object, including the bucket name, object name, and; generation number.
	Kind                    *string             `json:"kind,omitempty"`                   // The kind of item this is. For objects, this is always "storage#object".
	KmsKeyName              *string             `json:"kmsKeyName,omitempty"`             // Cloud KMS Key used to encrypt this object, if the object is encrypted by; such a key.
	Md5Hash                 *string             `json:"md5Hash,omitempty"`                // MD5 hash of the data; encoded using base64 as per; [https://tools.ietf.org/html/rfc4648#section-4][RFC 4648 §4]. For more; information about using the MD5 hash, see; [https://cloud.google.com/storage/docs/hashes-etags#_JSONAPI][Hashes and; ETags: Best Practices].
	MediaLink               *string             `json:"mediaLink,omitempty"`              // Media download link.
	Metadata                map[string]string   `json:"metadata,omitempty"`               // User-provided metadata, in key/value pairs.
	Metageneration          *Generation         `json:"metageneration"`                   // The version of the metadata for this object at this generation. Used for; preconditions and for detecting changes in metadata. A metageneration; number is only meaningful in the context of a particular generation of a; particular object.
	Name                    *string             `json:"name,omitempty"`                   // The name of the object.
	RetentionExpirationTime *string             `json:"retentionExpirationTime,omitempty"`// A server-determined value that specifies the earliest time that the; object's retention period expires.
	SelfLink                *string             `json:"selfLink,omitempty"`               // The link to this object.
	Size                    *Generation         `json:"size"`                             // Content-Length of the object data in bytes, matching; [https://tools.ietf.org/html/rfc7230#section-3.3.2][RFC 7230 §3.3.2].
	StorageClass            *string             `json:"storageClass,omitempty"`           // Storage class of the object.
	TemporaryHold           *bool               `json:"temporaryHold,omitempty"`          // Whether an object is under temporary hold.
	TimeCreated             *string             `json:"timeCreated,omitempty"`            // The creation time of the object.; Attempting to set this field will result in an error.
	TimeDeleted             *string             `json:"timeDeleted,omitempty"`            // The deletion time of the object. Will be returned if and only if this; version of the object has been deleted.
	TimeStorageClassUpdated *string             `json:"timeStorageClassUpdated,omitempty"`// The time at which the object's storage class was last changed.
	Updated                 *string             `json:"updated,omitempty"`                // The modification time of the object metadata.
}

// Metadata of customer-supplied encryption key, if the object is encrypted by
// such a key.
type CustomerEncryption struct {
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty"`// The encryption algorithm.
	KeySha256           *string `json:"keySha256,omitempty"`          // SHA256 hash value of the encryption key.
}

type Generation struct {
	Integer *int64
	String  *string
}

func UnmarshalStorageObjectData(data []byte) (StorageObjectData, error) {
	var d StorageObjectData
	err := json.Unmarshal(data, &d)
	return d, err
}

func (p *StorageObjectData) MarshalStorageObjectData() ([]byte, error) {
	return json.Marshal(p)
}