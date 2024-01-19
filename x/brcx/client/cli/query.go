package cli

import (
	"fmt"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"strings"

	"github.com/spf13/cobra"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/flags"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	"github.com/brc20-collab/brczero/x/brcx/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group brcx queries under a subcommand
	brcQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the brcx module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	brcQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdQueryProtocol(queryRoute, cdc),
		)...,
	)

	return brcQueryCmd

}

// GetCmdQuerySigningInfo implements the command to query signing info.
func GetCmdQueryProtocol(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		//todo
		Use:   "protocol [name]",
		Short: "Query addresses of protocol",
		Long: strings.TrimSpace(`Query addresses of protocol:

$ <brczerocli> query brcx protocol brc-20
`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryProtocol)
			bz, _, err := cliCtx.QueryWithData(route, []byte(args[0]))
			if err != nil {
				return err
			}
			var result []sdk.AccAddress
			cdc.MustUnmarshalJSON(bz, &result)

			addrs := make([]string, 0)
			for i := range result {
				addrs = append(addrs, common.BytesToAddress(result[i].Bytes()).Hex())
			}
			return cliCtx.PrintOutput(addrs)
		},
	}
}
