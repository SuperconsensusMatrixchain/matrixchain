/*
 * Copyright (c) 2021. Baidu Inc. All Rights Reserved.
 */

package cmd

import (
	"context"
	"fmt"
	"github.com/superconsensus-chain/xupercore/lib/utils"
	"github.com/xuperchain/xuperchain/service/pb"
	"io/ioutil"

	"github.com/spf13/cobra"

	"github.com/superconsensus-chain/xupercore/bcs/ledger/xledger/state/utxo"
)

// ProposalProposeCommand proposal a proposal struct
type ProposalProposeCommand struct {
	cli *Cli
	cmd *cobra.Command

	proposal string
	fee      string
}

// NewProposalProposeCommand propose a proposal cmd
func NewProposalProposeCommand(cli *Cli) *cobra.Command {
	t := new(ProposalProposeCommand)
	t.cli = cli
	t.cmd = &cobra.Command{
		Use:   "propose",
		Short: "Propose a proposal.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			return t.proposeProposal(ctx)
		},
	}
	t.addFlags()

	return t.cmd
}

func (c *ProposalProposeCommand) addFlags() {
	c.cmd.Flags().StringVarP(&c.proposal, "proposal", "p", "", "proposal.")
	c.cmd.Flags().StringVar(&c.fee, "fee", "0", "The fee to propose a proposal.")
}

func (c *ProposalProposeCommand) proposeProposal(ctx context.Context) error {
	ct := &CommTrans{
		Amount:       "0",
		Fee:          c.fee,
		FrozenHeight: 0,
		Version:      utxo.TxVersion,

		MethodName: "Propose",
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

	proposal, err := c.getProposal()
	if err != nil {
		return err
	}

	ct.ModuleName = "xkernel"
	ct.ContractName = "$proposal"
	ct.Args["proposal"] = proposal

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
	// 参照了consensus_invoke，向提案命令注入高度参数，不过不需要减3
	ct.Args["height"] = []byte(fmt.Sprintf("%d", status.GetMeta().TrunkHeight))

	err = ct.Transfer(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *ProposalProposeCommand) getProposal() ([]byte, error) {
	if c.proposal == "" {
		return []byte("no proposal"), nil
	}
	return ioutil.ReadFile(c.proposal)
}
