module github.com/pangum/mqtt

go 1.16

require (
	github.com/eclipse/paho.mqtt.golang v1.3.5
	github.com/google/uuid v1.3.0
	github.com/rs/xid v1.3.0
	github.com/storezhang/glog v1.0.8
	github.com/storezhang/gox v1.6.4
	github.com/storezhang/pangu v1.3.9
	github.com/vmihailenco/msgpack/v5 v5.3.4
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	google.golang.org/protobuf v1.27.1
)

// replace github.com/storezhang/gox => ../gox
// replace github.com/storezhang/echox/v2 => ../../storezhang/echox
