module shop

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/envoyproxy/protoc-gen-validate v0.6.3
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20220209030627-9662ef3c213d
	github.com/go-kratos/kratos/v2 v2.1.5
	github.com/google/wire v0.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/hashicorp/consul/api v1.12.0
	go.opentelemetry.io/otel v1.4.0
	go.opentelemetry.io/otel/exporters/jaeger v1.4.0
	go.opentelemetry.io/otel/sdk v1.4.0
	google.golang.org/genproto v0.0.0-20211223182754-3ac035c7e7cb
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
