package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/certikfoundation/shentu/v2/x/interview/types"
)

// NewTxCmd returns the transaction commands for the interview module.
func NewTxCmd() *cobra.Command {
	txCmds := &cobra.Command{
		Use:   "interview",
		Short: "Interview module's commands",
	}

	txCmds.AddCommand(
		GetCmdLockUser(),
	)

	return txCmds
}

// GetCmdLockUser returns the interview module's transaction command.
func GetCmdLockUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lock-user <user id>",
		Short: "Lock a user with given user id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(cliCtx, cmd.Flags()).WithTxConfig(cliCtx.TxConfig).WithAccountRetriever(cliCtx.AccountRetriever)

			from := cliCtx.GetFromAddress()
			if err := txf.AccountRetriever().EnsureExists(cliCtx, from); err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgLockUser(from.String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxWithFactory(cliCtx, txf, msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
