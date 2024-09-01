module app

go 1.23

require (
	github.com/alexflint/go-arg v1.5.1
	github.com/grafana/loki v1.6.1
	github.com/klauspost/compress v1.17.9
	github.com/prometheus/prometheus v1.8.2-0.20200727090838-6f296594a852
	github.com/skovtunenko/graterm v1.1.0
	github.com/xhhuango/json v1.19.0
)

exclude k8s.io/client-go v12.0.0+incompatible

// fixes
// github.com/golang/protobuf@v1.5.0/protoc-gen-go/descriptor/descriptor.pb.go:106:61: undefined: descriptorpb.Default_FileOptions_PhpGenericServices
replace google.golang.org/protobuf => google.golang.org/protobuf v1.31.0

require (
	github.com/alexflint/go-scalar v1.2.0 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto v0.0.0-20200724131911-43cab4749ae7 // indirect
	google.golang.org/grpc v1.30.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
