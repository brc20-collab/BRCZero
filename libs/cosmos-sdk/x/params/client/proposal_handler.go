package client

import (
	govclient "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/gov/client"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/params/client/cli"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/params/client/rest"
)

// ProposalHandler handles param change proposals
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
