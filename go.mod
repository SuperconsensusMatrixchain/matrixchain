module github.com/superconsensus/matrixchain

go 1.14

require (
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/hyperledger/burrow v0.30.5
	github.com/manifoldco/promptui v0.7.0
	github.com/prometheus/client_golang v1.11.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.2
	github.com/superconsensus/matrixcore v1.0.3-0.20220408071608-b2a219e65ee1
	github.com/xuperchain/crypto v0.0.0-20211221122406-302ac826ac90
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.35.0
)

replace github.com/hyperledger/burrow => github.com/xuperchain/burrow v0.30.6-0.20211229032028-fbee6a05ab0f
