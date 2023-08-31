package rest

import (
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
	govRest "github.com/brc20-collab/brczero/x/gov/client/rest"
)

// ProposeValidatorProposalRESTHandler defines propose validator proposal handler
func ProposeValidatorProposalRESTHandler(context.CLIContext) govRest.ProposalRESTHandler {
	return govRest.ProposalRESTHandler{}
}
