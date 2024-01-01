package cli

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	clictx "github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/flags"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/server"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types/module"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth/client/utils"
	tmtypes "github.com/brc20-collab/brczero/libs/tendermint/types"
	"github.com/brc20-collab/brczero/x/brcx"
)

const (
	flagInscription        = "inscription"
	flagInscriptionContext = "inscription_context"
)

// GenTxCmd builds the application's gentx command
func GenMsgInscriptionCmd(ctx *server.Context, cdc *codec.Codec, mbm module.BasicManager, defaultNodeHome, defaultCLIHome string) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "geninscription",
		Short: "Generate a genesis tx for inscription",
		Args:  cobra.NoArgs,
		Long:  fmt.Sprintf(`This command is uesd for init inscription`),

		RunE: func(cmd *cobra.Command, args []string) error {

			config := ctx.Config
			config.SetRoot(viper.GetString(flags.FlagHome))

			genDoc, err := tmtypes.GenesisDocFromFile(config.GenesisFile())
			if err != nil {
				return err
			}

			var genesisState map[string]json.RawMessage
			if err = cdc.UnmarshalJSON(genDoc.AppState, &genesisState); err != nil {
				return err
			}

			if err = mbm.ValidateGenesis(genesisState); err != nil {
				return err
			}

			inBuf := bufio.NewReader(cmd.InOrStdin())

			// Set flags for creating gentx
			viper.Set(flags.FlagHome, viper.GetString(flagClientHome))

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := clictx.NewCLIContext().WithCodec(cdc)

			// Set the generate-only flag here after the CLI context has
			// been created. This allows the from name/key to be correctly populated.
			//
			// TODO: Consider removing the manual setting of generate-only in
			// favor of a 'gentx' flag in the create-validator command.
			viper.Set(flags.FlagGenerateOnly, true)
			inscription := viper.GetString(flagInscription)
			inscriptionContextStr := viper.GetString(flagInscriptionContext)
			if len(inscription) == 0 || len(inscriptionContextStr) == 0 {
				return fmt.Errorf("the inscription or inscription_context is empty")
			}
			var inscriptionContext brcx.InscriptionContext
			if err := json.Unmarshal([]byte(inscriptionContextStr), &inscriptionContext); err != nil {
				return err
			}
			msg := brcx.NewMsgInscription(inscription, inscriptionContext)

			// write the unsigned transaction to the buffer
			w := bytes.NewBuffer([]byte{})
			cliCtx = cliCtx.WithOutput(w)

			if err = utils.PrintUnsignedStdTx(txBldr, cliCtx, []sdk.Msg{msg}); err != nil {
				return err
			}

			// read the transaction
			stdTx, err := readUnsignedGenTxFile(cdc, w)
			if err != nil {
				return err
			}

			// Fetch output file name
			outputDocument := viper.GetString(flags.FlagOutputDocument)
			if outputDocument == "" {
				outputDocument, err = makeOutputFilepath(config.RootDir, hex.EncodeToString(stdTx.Hash))
				if err != nil {
					return err
				}
			}

			if err := writeSignedGenTx(cdc, outputDocument, stdTx); err != nil {
				return err
			}

			if _, err := fmt.Fprintf(os.Stderr, "Genesis transaction written to %q\n", outputDocument); err != nil {
				ctx.Logger.Error(err.Error())
			}

			return nil

		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "node's home directory")
	cmd.Flags().String(flagClientHome, defaultCLIHome, "client's home directory")
	cmd.Flags().String(flagInscription, "", "inscription for msginscription")
	cmd.Flags().String(flagInscriptionContext, "", "inscriptionContext for msginscription")
	cmd.Flags().String(flags.FlagOutputDocument, "",
		"write the genesis transaction JSON document to the given file instead of the default location")
	cmd.Flags().String(flagChainId, "", "inscriptionContext for msginscription")
	if err := cmd.MarkFlagRequired(flagInscription); err != nil {
		ctx.Logger.Error(err.Error())
	}
	if err := cmd.MarkFlagRequired(flagInscriptionContext); err != nil {
		ctx.Logger.Error(err.Error())
	}
	return cmd
}

// GenTxCmd builds the application's gentx command
func GenMsgBasicxCmd(ctx *server.Context, cdc *codec.Codec, mbm module.BasicManager, defaultNodeHome, defaultCLIHome string) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "gen-basicx",
		Short: "Generate a genesis tx for basicx",
		Args:  cobra.NoArgs,
		Long:  fmt.Sprintf(`This command is uesd for init basicx`),

		RunE: func(cmd *cobra.Command, args []string) error {

			config := ctx.Config
			config.SetRoot(viper.GetString(flags.FlagHome))

			genDoc, err := tmtypes.GenesisDocFromFile(config.GenesisFile())
			if err != nil {
				return err
			}

			var genesisState map[string]json.RawMessage
			if err = cdc.UnmarshalJSON(genDoc.AppState, &genesisState); err != nil {
				return err
			}

			if err = mbm.ValidateGenesis(genesisState); err != nil {
				return err
			}

			inBuf := bufio.NewReader(cmd.InOrStdin())

			// Set flags for creating gentx
			viper.Set(flags.FlagHome, viper.GetString(flagClientHome))

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := clictx.NewCLIContext().WithCodec(cdc)

			// Set the generate-only flag here after the CLI context has
			// been created. This allows the from name/key to be correctly populated.
			//
			// TODO: Consider removing the manual setting of generate-only in
			// favor of a 'gentx' flag in the create-validator command.
			viper.Set(flags.FlagGenerateOnly, true)
			inscription := viper.GetString(flagInscription)
			inscriptionContextStr := viper.GetString(flagInscriptionContext)
			if len(inscription) == 0 || len(inscriptionContextStr) == 0 {
				return fmt.Errorf("the inscription or inscription_context is empty")
			}

			msg := brcx.NewMsgBasicProtocolOp(brcx.ManageContractProtocolName, inscription, "1111111111111111111111111111111111111111111111111111111111111111", "", inscriptionContextStr)

			// write the unsigned transaction to the buffer
			w := bytes.NewBuffer([]byte{})
			cliCtx = cliCtx.WithOutput(w)

			if err = utils.PrintUnsignedStdTx(txBldr, cliCtx, []sdk.Msg{msg}); err != nil {
				return err
			}

			// read the transaction
			stdTx, err := readUnsignedGenTxFile(cdc, w)
			if err != nil {
				return err
			}

			// Fetch output file name
			outputDocument := viper.GetString(flags.FlagOutputDocument)
			if outputDocument == "" {
				outputDocument, err = makeOutputFilepath(config.RootDir, hex.EncodeToString(stdTx.Hash))
				if err != nil {
					return err
				}
			}

			if err := writeSignedGenTx(cdc, outputDocument, stdTx); err != nil {
				return err
			}

			if _, err := fmt.Fprintf(os.Stderr, "Genesis transaction written to %q\n", outputDocument); err != nil {
				ctx.Logger.Error(err.Error())
			}

			return nil

		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "node's home directory")
	cmd.Flags().String(flagClientHome, defaultCLIHome, "client's home directory")
	cmd.Flags().String(flagInscription, "", "inscription for msgBasicX")
	cmd.Flags().String(flagInscriptionContext, "", "inscriptionContext for msgBasicX")
	cmd.Flags().String(flags.FlagOutputDocument, "",
		"write the genesis transaction JSON document to the given file instead of the default location")
	cmd.Flags().String(flagChainId, "", "inscriptionContext for msginscription")
	if err := cmd.MarkFlagRequired(flagInscription); err != nil {
		ctx.Logger.Error(err.Error())
	}
	if err := cmd.MarkFlagRequired(flagInscriptionContext); err != nil {
		ctx.Logger.Error(err.Error())
	}
	return cmd
}
