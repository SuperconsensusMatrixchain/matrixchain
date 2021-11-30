package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

type TotalCommand struct {
	cli *Cli
	cmd *cobra.Command
	// 查询方式 exchange|total，默认前者
	mode string
}

func NewTotalCommand(cli *Cli) *cobra.Command {
	t := new(TotalCommand)
	t.cli = cli
	t.cmd = &cobra.Command{
		Use:   "total",
		Short: "Query the amount of govern token in the whole network. exchange|total",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			return t.Total(ctx)
		},
	}
	t.addFlags()

	return t.cmd
}

func (c *TotalCommand) addFlags() {
	c.cmd.Flags().StringVarP(&c.mode, "mode", "m", "exchange", "exchange: 全网已有多少UTXO兑换成治理代币|total: 全网UTXO总量如果换算为治理代币会有多少")
}

func (c *TotalCommand) Total(ctx context.Context) error {
	ct := &CommTrans{
		//Fee:	c.fee,
		//Descfile: c.desc,
		Args:       make(map[string][]byte),
		//IsQuick: false,
		ChainName:    c.cli.RootOptions.Name,
		Keys:         c.cli.RootOptions.Keys,
		XchainClient: c.cli.XchainClient(),
		//CryptoType:   c.cli.RootOptions.Crypto,
		//RootOptions:  c.cli.RootOptions,
		ModuleName: "xkernel",
		ContractName: "$govern_token",
		MethodName: "TotalSupply",
	}

	if c.mode == "exchange" {
		ct.MethodName = "TotalSupply"
	}else if c.mode == "total" {
		ct.MethodName = "AllToken"
	}else {
		return fmt.Errorf("不存在的查询方式: %s", c.mode)
	}

	response, _, err := ct.GenPreExeRes(ctx)
	_ = response
	//fmt.Println("response", response)
	//response header:<logid:"1633663466_7034442070621071" from_node:"vv_server" > bcname:"xuper" response:<response:"hello bonus balance" gas_used:1000000 requests:<module_name:"xkernel" contract_name:"$govern_token" method_name:"Bonus" args:<key:"account" value:"a" > resource_limits:<> resource_limits:<type:MEMORY > resource_limits:<type:DISK > resource_limits:<type:XFEE > > responses:<status:200 message:"success" body:"hello bonus balance" > >
	//for _, req := range response.GetResponse().GetRequests() {
	//	limits := req.GetResourceLimits()
	//	for _, limit := range limits {
	//		fmt.Println(limit.Type.String(), ": ", limit.Limit)
	//	}
	//}
	return err

}
