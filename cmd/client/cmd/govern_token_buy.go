package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/superconsensus-chain/xupercore/bcs/ledger/xledger/state/utxo"
)

type GovernBuyTokenCommand struct {
	cli *Cli
	cmd *cobra.Command

	amount   string
	fee      string
	desc     string
}

func NewGovernBuyTokenCommand(cli *Cli) *cobra.Command {
	t := new(GovernBuyTokenCommand)
	t.cli = cli
	t.cmd = &cobra.Command{
		Use:   "buy",
		Short: "buy govern token.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			return t.BuyToken(ctx)
		},
	}
	t.addFlags()

	return t.cmd
}

func (c *GovernBuyTokenCommand) addFlags() {
	c.cmd.Flags().StringVar(&c.fee, "fee", "0", "The fee to initialize govern token.")
	c.cmd.Flags().StringVar(&c.amount, "amount", "0", "The amount to buy govern token.")
	c.cmd.Flags().StringVar(&c.desc, "desc", "0", "transaction description.")
}

func (c *GovernBuyTokenCommand) BuyToken(ctx context.Context) error {
	ct := &CommTrans{
		Amount: c.amount,
		Fee:	c.fee,
		Descfile: c.desc,
		FrozenHeight: 0,
		Version:      utxo.TxVersion,
		MethodName: "Buy",
		Args:       make(map[string][]byte),
		IsQuick: false,

		ChainName:    c.cli.RootOptions.Name,
		Keys:         c.cli.RootOptions.Keys,
		XchainClient: c.cli.XchainClient(),
		CryptoType:   c.cli.RootOptions.Crypto,
		RootOptions:  c.cli.RootOptions,
	}
	var err error
	//ct.To, err = readAddress(ct.Keys)
	//if err != nil {
	//	return err
	//}
	ct.To = "testa"

	ct.ModuleName = "xkernel"
	ct.ContractName = "$govern_token"
	ct.Args["amount"] = []byte(c.amount)
	ct.Args["desc"] = []byte(c.desc)

	err = ct.Transfer(ctx)
	if err != nil {
		return err
	}

	return nil

}