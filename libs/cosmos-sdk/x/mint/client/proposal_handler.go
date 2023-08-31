package client

import (
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/mint/client/cli"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/mint/client/rest"
	govcli "github.com/brc20-collab/brczero/x/gov/client"
)

var (
	ManageTreasuresProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdManageTreasuresProposal,
		rest.ManageTreasuresProposalRESTHandler,
	)

	ExtraProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdExtraProposal,
		rest.ExtraProposalRESTHandler,
	)
)
