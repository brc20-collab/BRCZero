package client

import (
	"github.com/brc20-collab/brczero/x/staking/client/cli"
	"github.com/brc20-collab/brczero/x/staking/client/rest"
	govcli "github.com/brc20-collab/brczero/x/gov/client"
)

var (
	ProposeValidatorProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdProposeValidatorProposal,
		rest.ProposeValidatorProposalRESTHandler,
	)
)

