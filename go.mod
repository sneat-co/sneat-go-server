module github.com/sneat-co/sneat-go-server

go 1.23.2

//replace github.com/sneat-co/sneat-go-backend => ../sneat-go-backend

require (
	github.com/julienschmidt/httprouter v1.3.0
	github.com/sneat-co/sneat-core-modules v0.15.2
	github.com/sneat-co/sneat-go-backend v0.54.0
	github.com/sneat-co/sneat-go-core v0.37.1
	github.com/sneat-co/sneat-go-firebase v0.6.2
	github.com/strongo/delaying v0.0.1
	github.com/strongo/logus v0.2.0
)

require (
	cel.dev/expr v0.16.2 // indirect
	cloud.google.com/go v0.115.1 // indirect
	cloud.google.com/go/auth v0.9.7 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.4 // indirect
	cloud.google.com/go/compute/metadata v0.5.2 // indirect
	cloud.google.com/go/firestore v1.17.0 // indirect
	cloud.google.com/go/iam v1.2.1 // indirect
	cloud.google.com/go/longrunning v0.6.1 // indirect
	cloud.google.com/go/monitoring v1.21.1 // indirect
	cloud.google.com/go/storage v1.44.0 // indirect
	firebase.google.com/go/v4 v4.14.1 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.24.2 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.48.2 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.48.2 // indirect
	github.com/MicahParks/keyfunc v1.9.0 // indirect
	github.com/alexsergivan/transliterator v1.0.1 // indirect
	github.com/bots-go-framework/bots-api-telegram v0.7.1 // indirect
	github.com/bots-go-framework/bots-fw v0.40.2 // indirect
	github.com/bots-go-framework/bots-fw-store v0.8.2 // indirect
	github.com/bots-go-framework/bots-fw-telegram v0.13.6 // indirect
	github.com/bots-go-framework/bots-fw-telegram-models v0.3.7 // indirect
	github.com/bots-go-framework/bots-fw-telegram-webapp v0.3.1 // indirect
	github.com/bots-go-framework/bots-go-core v0.0.3 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cncf/xds/go v0.0.0-20240905190251-b4127c9b8d78 // indirect
	github.com/crediterra/money v0.3.0 // indirect
	github.com/dal-go/dalgo v0.14.0 // indirect
	github.com/dal-go/dalgo2firestore v0.2.9 // indirect
	github.com/envoyproxy/go-control-plane v0.13.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/googleapis/gax-go/v2 v2.13.0 // indirect
	github.com/gosimple/slug v1.14.0 // indirect
	github.com/gosimple/unidecode v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/pquerna/ffjson v0.0.0-20190930134022-aa0246cd15f7 // indirect
	github.com/strongo/decimal v0.1.1 // indirect
	github.com/strongo/facebook v1.8.1 // indirect
	github.com/strongo/gamp v0.0.1 // indirect
	github.com/strongo/i18n v0.6.1 // indirect
	github.com/strongo/random v0.0.1 // indirect
	github.com/strongo/slice v0.3.0 // indirect
	github.com/strongo/strongoapp v0.25.2 // indirect
	github.com/strongo/validation v0.0.7 // indirect
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/detectors/gcp v1.30.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.55.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.55.0 // indirect
	go.opentelemetry.io/otel v1.30.0 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	go.opentelemetry.io/otel/sdk v1.30.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.30.0 // indirect
	go.opentelemetry.io/otel/trace v1.30.0 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/oauth2 v0.23.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	golang.org/x/time v0.7.0 // indirect
	google.golang.org/api v0.199.0 // indirect
	google.golang.org/appengine/v2 v2.0.6 // indirect
	google.golang.org/genproto v0.0.0-20240930140551-af27646dc61f // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240930140551-af27646dc61f // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240930140551-af27646dc61f // indirect
	google.golang.org/grpc v1.67.1 // indirect
	google.golang.org/grpc/stats/opentelemetry v0.0.0-20241004113128-859602c14c6c // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
