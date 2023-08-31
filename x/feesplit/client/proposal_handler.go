package client

import (
	"github.com/brc20-collab/brczero/x/feesplit/client/cli"
	"github.com/brc20-collab/brczero/x/feesplit/client/rest"
	govcli "github.com/brc20-collab/brczero/x/gov/client"
)

var (
	// FeeSplitSharesProposalHandler alias gov NewProposalHandler
	FeeSplitSharesProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdFeeSplitSharesProposal,
		rest.FeeSplitSharesProposalRESTHandler,
	)
)
