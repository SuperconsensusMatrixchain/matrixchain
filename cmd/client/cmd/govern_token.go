/*
 * Copyright (c) 2021. Baidu Inc. All Rights Reserved.
 */

package cmd

import (
	"github.com/spf13/cobra"
)

// GovernTokenCommand govern token cmd entrance
type GovernTokenCommand struct {
	cli *Cli
	cmd *cobra.Command
}

// NewGovernTokenCommand new govern token cmd
func NewGovernTokenCommand(cli *Cli) *cobra.Command {
	c := new(ProposalCommand)
	c.cli = cli
	c.cmd = &cobra.Command{
		Use:   "governToken",
		Short: "governToken: total|bonusQuery|bonusObtain|query|buy|sell.",
	}
	//c.cmd.AddCommand(NewGovernInitCommand(cli))
	//c.cmd.AddCommand(NewGovernTransferCommand(cli))
	c.cmd.AddCommand(NewTotalCommand(cli))
	c.cmd.AddCommand(NewGovernTokenQueryCommand(cli))
	c.cmd.AddCommand(NewGovernBuyTokenCommand(cli))
	c.cmd.AddCommand(NewGovernSellTokenCommand(cli))
	c.cmd.AddCommand(NewBonusObtainCommand(cli))
	c.cmd.AddCommand(NewBonusQueryCommand(cli))
	return c.cmd
}

func init() {
	AddCommand(NewGovernTokenCommand)
}
