package main

import (
	"github.com/brc20-collab/brczero/app"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types/tx"
	mintclient "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/mint/client"
	mintrest "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/mint/client/rest"
	evmclient "github.com/brc20-collab/brczero/x/evm/client"

	"github.com/brc20-collab/brczero/app/rpc"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/lcd"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth"
	authrest "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth/client/rest"
	bankrest "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/bank/client/rest"
	supplyrest "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/supply/client/rest"
	brcxrest "github.com/brc20-collab/brczero/x/brcx/client/rest"
	dist "github.com/brc20-collab/brczero/x/distribution"
	distr "github.com/brc20-collab/brczero/x/distribution"
	distrest "github.com/brc20-collab/brczero/x/distribution/client/rest"
	evmrest "github.com/brc20-collab/brczero/x/evm/client/rest"
	govrest "github.com/brc20-collab/brczero/x/gov/client/rest"
	paramsclient "github.com/brc20-collab/brczero/x/params/client"
	stakingclient "github.com/brc20-collab/brczero/x/staking/client"
	stakingrest "github.com/brc20-collab/brczero/x/staking/client/rest"
	"github.com/brc20-collab/brczero/x/token"
	tokensrest "github.com/brc20-collab/brczero/x/token/client/rest"
)

// registerRoutes registers the routes from the different modules for the LCD.
// NOTE: details on the routes added for each module are in the module documentation
// NOTE: If making updates here you also need to update the test helper in client/lcd/test_helper.go
func registerRoutes(rs *lcd.RestServer) {
	registerGrpc(rs)
	rpc.RegisterRoutes(rs)
	registerRoutesV1(rs)
	registerRoutesV2(rs)
}

func registerGrpc(rs *lcd.RestServer) {
	app.ModuleBasics.RegisterGRPCGatewayRoutes(rs.CliCtx, rs.GRPCGatewayRouter)
	app.ModuleBasics.RegisterRPCRouterForGRPC(rs.CliCtx, rs.Mux)
	tx.RegisterGRPCGatewayRoutes(rs.CliCtx, rs.GRPCGatewayRouter)
}

func registerRoutesV1(rs *lcd.RestServer) {
	v1Router := rs.Mux.PathPrefix("/v1").Name("v1").Subrouter()
	client.RegisterRoutes(rs.CliCtx, v1Router)
	authrest.RegisterRoutes(rs.CliCtx, v1Router, auth.StoreKey)
	bankrest.RegisterRoutes(rs.CliCtx, v1Router)
	stakingrest.RegisterRoutes(rs.CliCtx, v1Router)
	distrest.RegisterRoutes(rs.CliCtx, v1Router, dist.StoreKey)

	tokensrest.RegisterRoutes(rs.CliCtx, v1Router, token.StoreKey)
	supplyrest.RegisterRoutes(rs.CliCtx, v1Router)
	evmrest.RegisterRoutes(rs.CliCtx, v1Router)
	govrest.RegisterRoutes(rs.CliCtx, v1Router,
		[]govrest.ProposalRESTHandler{
			paramsclient.ProposalHandler.RESTHandler(rs.CliCtx),
			distr.CommunityPoolSpendProposalHandler.RESTHandler(rs.CliCtx),
			distr.ChangeDistributionTypeProposalHandler.RESTHandler(rs.CliCtx),
			distr.WithdrawRewardEnabledProposalHandler.RESTHandler(rs.CliCtx),
			distr.RewardTruncatePrecisionProposalHandler.RESTHandler(rs.CliCtx),
			evmclient.ManageContractDeploymentWhitelistProposalHandler.RESTHandler(rs.CliCtx),
			evmclient.ManageSysContractAddressProposalHandler.RESTHandler(rs.CliCtx),
			evmclient.ManageContractByteCodeProposalHandler.RESTHandler(rs.CliCtx),
			mintclient.ManageTreasuresProposalHandler.RESTHandler(rs.CliCtx),
			mintclient.ExtraProposalHandler.RESTHandler(rs.CliCtx),
			stakingclient.ProposeValidatorProposalHandler.RESTHandler(rs.CliCtx),
		},
	)
	mintrest.RegisterRoutes(rs.CliCtx, v1Router)
	brcxrest.RegisterRoutes(rs.CliCtx, v1Router)

}

func registerRoutesV2(rs *lcd.RestServer) {
	v2Router := rs.Mux.PathPrefix("/v2").Name("v2").Subrouter()
	client.RegisterRoutes(rs.CliCtx, v2Router)
	authrest.RegisterRoutes(rs.CliCtx, v2Router, auth.StoreKey)
	bankrest.RegisterRoutes(rs.CliCtx, v2Router)
	stakingrest.RegisterRoutes(rs.CliCtx, v2Router)
	distrest.RegisterRoutes(rs.CliCtx, v2Router, dist.StoreKey)
	tokensrest.RegisterRoutesV2(rs.CliCtx, v2Router, token.StoreKey)
}
