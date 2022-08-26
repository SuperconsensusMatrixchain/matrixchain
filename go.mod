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
	github.com/superconsensus/matrixcore v1.0.1
	github.com/xuperchain/crypto v0.0.0-20201028025054-4d560674bcd6
	golang.org/x/net v0.0.0-20201021035429-f5854403a974
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.35.0
)

replace github.com/hyperledger/burrow => github.com/xuperchain/burrow v0.30.6-0.20211229032028-fbee6a05ab0f
// 由原先仓库依赖改为最新（6.28 gas up）
replace github.com/superconsensus/matrixcore v1.0.1 => github.com/SuperconsensusMatrixchain/matrixcore v0.0.0-20220628084951-564785ea7345
