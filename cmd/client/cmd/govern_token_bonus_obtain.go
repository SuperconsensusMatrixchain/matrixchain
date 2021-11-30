package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/superconsensus-chain/xupercore/bcs/ledger/xledger/state/utxo"
	"github.com/superconsensus-chain/xupercore/lib/utils"
	"github.com/xuperchain/xuperchain/service/pb"
)

type BonusObtainCommand struct {
	cli *Cli
	cmd *cobra.Command

	amount   string
	fee      string
	//desc     string
}

func NewBonusObtainCommand(cli *Cli) *cobra.Command {
	t := new(BonusObtainCommand)
	t.cli = cli
	t.cmd = &cobra.Command{
		Use:   "bonusObtain",
		Short: "obtain bonus reward(UTXO).",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			return t.Obtain(ctx)
		},
	}
	t.addFlags()

	return t.cmd
}

func (c *BonusObtainCommand) addFlags() {
	c.cmd.Flags().StringVar(&c.fee, "fee", "0", "The fee to initialize govern token.")
	c.cmd.Flags().StringVar(&c.amount, "amount", "0", "The amount to obtain UTXO from bonus reward.")
	//c.cmd.Flags().StringVar(&c.desc, "desc", "0", "transaction description.")
}

func (c *BonusObtainCommand) Obtain(ctx context.Context) error {
	ct := &CommTrans{
		Amount: "0",
		Fee:	c.fee,
		//Descfile: c.desc,
		FrozenHeight: 0,
		Version:      utxo.TxVersion,
		MethodName: "BonusObtain",
		Args:       make(map[string][]byte),
		IsQuick: false,

		ChainName:    c.cli.RootOptions.Name,
		Keys:         c.cli.RootOptions.Keys,
		XchainClient: c.cli.XchainClient(),
		CryptoType:   c.cli.RootOptions.Crypto,
		RootOptions:  c.cli.RootOptions,
	}
	var err error
	ct.To, err = readAddress(ct.Keys)
	if err != nil {
		return err
	}

	ct.ModuleName = "xkernel"
	ct.ContractName = "$govern_token"
	ct.Args["amount"] = []byte(c.amount)

	bcStatus := &pb.BCStatus{
		Header: &pb.Header{
			Logid: utils.GenLogId(),
		},
		Bcname: c.cli.RootOptions.Name,
	}
	status, err := c.cli.XchainClient().GetBlockChainStatus(ctx, bcStatus)
	if err != nil {
		return fmt.Errorf("get chain status error.\n")
	}
	// 参照了consensus_invoke，向分红提现命令注入高度参数，不过不需要减3，反而需要+2，+1是本次交易通过后上链的高度，真正到账高度应该为本次交易的下一个块，即+2
	ct.Args["height"] = []byte(fmt.Sprintf("%d", status.GetMeta().TrunkHeight + 2))

	err = ct.Transfer(ctx)
	if err != nil {
		return err
	}

	return nil

}