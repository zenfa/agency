package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/zenfa/agency/x/agency/types"
)

func GetCmdListAgency(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-agency",
		Short: "list all agency",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryListAgency), nil)
			if err != nil {
				fmt.Printf("could not list Commit\n%s\n", err.Error())
				return nil
			}
			var out []types.Agency
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetAgency(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-agency [principal] [agent]",
		Short: "Get a agency by principal and agent",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			principal := args[0]
			agent := args[1]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s/%s", queryRoute, types.QueryGetAgency, principal, agent), nil)

			if err != nil {
				fmt.Printf("could not get Commit with principal %s and agent %s\n%s\n", principal, agent, err.Error())
				return nil
			}
			var out types.Agency
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
