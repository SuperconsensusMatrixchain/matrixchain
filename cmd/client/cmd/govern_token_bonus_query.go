package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

type BonusQueryCommand struct {
	cli *Cli
	cmd *cobra.Command

	account   string
	//fee      string
	//desc     string
}

func NewBonusQueryCommand(cli *Cli) *cobra.Command {
	t := new(BonusQueryCommand)
	t.cli = cli
	t.cmd = &cobra.Command{
		Use:   "bonusQuery",
		Short: "query bonus reward(UTXO).",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			return t.Query(ctx)
		},
	}
	t.addFlags()

	return t.cmd
}

func (c *BonusQueryCommand) addFlags() {
	c.cmd.Flags().StringVarP(&c.account, "account", "A", "", "govern token account")
	//c.cmd.Flags().StringVar(&c.fee, "fee", "0", "The fee to initialize govern token.")
	//c.cmd.Flags().StringVar(&c.desc, "desc", "0", "transaction description.")
}

func (c *BonusQueryCommand) Query(ctx context.Context) error {
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
		MethodName: "BonusQuery",
	}
	//var err error
	//ct.To, err = readAddress(ct.Keys)
	//if err != nil {
	//	return err
	//}
	if c.account == "" {
		return fmt.Errorf("请输入查询账户")
	}

	ct.Args["account"] = []byte(c.account)
	//ct.Args["amount"] = []byte(c.amount)
	//ct.Args["desc"] = []byte(c.desc)

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
