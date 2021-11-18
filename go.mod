module github.com/pangum/mqtt

go 1.16

require (
	github.com/eclipse/paho.mqtt.golang v1.3.5
	github.com/google/uuid v1.3.0
	github.com/pangum/logging v0.0.7
	github.com/pangum/pangu v0.0.9
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/xid v1.3.0
	github.com/storezhang/gox v1.8.1
	github.com/vmihailenco/msgpack/v5 v5.3.4
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

// replace github.com/storezhang/gox => ../gox
// replace github.com/storezhang/echox/v2 => ../../storezhang/echox
