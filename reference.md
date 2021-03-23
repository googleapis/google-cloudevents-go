# Reference

The following sections describe how to use this library with different event data types.

<!-- GENERATED START -->
### Cloud Audit Logs

The data within all Cloud Audit Logs log entry events.

#### CloudEvent Types:

- `google.cloud.audit.log.v1.written`

#### Sample

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/googleapis/google-cloudevents-go/cloud/audit/v1"
)

func main() {
	data := []byte("CloudEvent Data in bytes")

	var e audit.LogEntryData
	err := json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", e)
}
```

### Cloud Build

Build event data for Google Cloud Platform API operations.

#### CloudEvent Types:

- `google.cloud.cloudbuild.build.v1.statusChanged`

#### Sample

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/googleapis/google-cloudevents-go/cloud/cloudbuild/v1"
)

func main() {
	data := []byte("CloudEvent Data in bytes")

	var e cloudbuild.BuildEventData
	err := json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", e)
}
```

### Cloud Firestore

The data within all Firestore document events.

#### CloudEvent Types:

- `google.cloud.firestore.document.v1.created`
- `google.cloud.firestore.document.v1.updated`
- `google.cloud.firestore.document.v1.deleted`
- `google.cloud.firestore.document.v1.written`

#### Sample

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/googleapis/google-cloudevents-go/cloud/firestore/v1"
)

func main() {
	data := []byte("CloudEvent Data in bytes")

	var e firestore.DocumentEventData
	err := json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", e)
}
```

### Cloud Pub/Sub

The event data when a message is published to a topic.

#### CloudEvent Types:

- `google.cloud.pubsub.topic.v1.messagePublished`

#### Sample

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/googleapis/google-cloudevents-go/cloud/pubsub/v1"
)

func main() {
	data := []byte("CloudEvent Data in bytes")

	var e pubsub.MessagePublishedData
	err := json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", e)
}
```

### Cloud Scheduler

Scheduler job data.

#### CloudEvent Types:

- `google.cloud.scheduler.job.v1.executed`

#### Sample

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/googleapis/google-cloudevents-go/cloud/scheduler/v1"
)

func main() {
	data := []byte("CloudEvent Data in bytes")

	var e scheduler.SchedulerJobData
	err := json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", e)
}
```

### Cloud Storage

An object within Google Cloud Storage.

#### CloudEvent Types:

- `google.cloud.storage.object.v1.finalized`
- `google.cloud.storage.object.v1.archived`
- `google.cloud.storage.object.v1.deleted`
- `google.cloud.storage.object.v1.metadataUpdated`

#### Sample

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/googleapis/google-cloudevents-go/cloud/storage/v1"
)

func main() {
	data := []byte("CloudEvent Data in bytes")

	var e storage.StorageObjectData
	err := json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", e)
}
```

### Google Analytics for Firebase

The data within Firebase Analytics log events.

#### CloudEvent Types:

- `google.firebase.analytics.log.v1.written`

#### Sample

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/googleapis/google-cloudevents-go/firebase/analytics/v1"
)

func main() {
	data := []byte("CloudEvent Data in bytes")

	var e analytics.AnalyticsLogData
	err := json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", e)
}
```

### Firebase Authentication

The data within all Firebase Auth events.

#### CloudEvent Types:

- `google.firebase.auth.user.v1.created`
- `google.firebase.auth.user.v1.deleted`

#### Sample

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/googleapis/google-cloudevents-go/firebase/auth/v1"
)

func main() {
	data := []byte("CloudEvent Data in bytes")

	var e auth.AuthEventData
	err := json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", e)
}
```

### Firebase Realtime Database

The data within all Firebase Real Time Database reference events.

#### CloudEvent Types:

- `google.firebase.database.ref.v1.created`
- `google.firebase.database.ref.v1.updated`
- `google.firebase.database.ref.v1.deleted`
- `google.firebase.database.ref.v1.written`

#### Sample

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/googleapis/google-cloudevents-go/firebase/database/v1"
)

func main() {
	data := []byte("CloudEvent Data in bytes")

	var e database.ReferenceEventData
	err := json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", e)
}
```

### Firebase Remote Config

The data within all Firebase Remote Config events.

#### CloudEvent Types:

- `google.firebase.remoteconfig.remoteConfig.v1.updated`

#### Sample

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/googleapis/google-cloudevents-go/firebase/remoteconfig/v1"
)

func main() {
	data := []byte("CloudEvent Data in bytes")

	var e remoteconfig.RemoteConfigEventData
	err := json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", e)
}
```

<!-- GENERATED END -->
