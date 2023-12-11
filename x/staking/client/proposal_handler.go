package client

import (
	govcli "github.com/brc20-collab/brczero/x/gov/client"
	"github.com/brc20-collab/brczero/x/staking/client/cli"
	"github.com/brc20-collab/brczero/x/staking/client/rest"
)

var (
	ProposeValidatorProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdProposeValidatorProposal,
		rest.ProposeValidatorProposalRESTHandler,
	)
)
