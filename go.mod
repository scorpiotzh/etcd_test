module github.com/scorpiotzh/etcd_test

go 1.15

require (
	github.com/coreos/bbolt v1.3.6 // indirect
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/scorpiotzh/mylog v1.0.8
	github.com/scorpiotzh/toolib v1.0.5
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20201229170055-e5319fda7802 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	go.uber.org/zap v1.19.0 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/genproto v0.0.0-20210825212027-de86158e7fda // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace github.com/coreos/bbolt v1.3.6 => go.etcd.io/bbolt v1.3.6

replace go.etcd.io/bbolt v1.3.6 => github.com/coreos/bbolt v1.3.6

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
