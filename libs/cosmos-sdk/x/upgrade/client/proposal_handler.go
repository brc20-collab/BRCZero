package client

import (
	govclient "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/gov/client"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/upgrade/client/cli"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
