package feesplit_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/brc20-collab/brczero/app"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"
	"github.com/brc20-collab/brczero/x/feesplit"
	"github.com/brc20-collab/brczero/x/feesplit/types"
)

type GenesisTestSuite struct {
	suite.Suite

	ctx     sdk.Context
	app     *app.BRCZeroApp
	genesis types.GenesisState
}

func (suite *GenesisTestSuite) SetupTest() {
	checkTx := false

	suite.app = app.Setup(false)
	suite.ctx = suite.app.BaseApp.NewContext(checkTx, abci.Header{Height: 1, ChainID: "ethermint-3", Time: time.Now().UTC()})
	suite.app.FeeSplitKeeper.SetParams(suite.ctx, types.DefaultParams())
	suite.genesis = types.DefaultGenesisState()
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}

func (suite *GenesisTestSuite) TestFeeSplitInitGenesis() {
	testCases := []struct {
		name     string
		genesis  types.GenesisState
		expPanic bool
	}{
		{
			"default genesis",
			suite.genesis,
			false,
		},
		{
			"custom genesis - feesplit disabled",
			types.GenesisState{
				Params: types.Params{
					EnableFeeSplit:           false,
					DeveloperShares:          types.DefaultDeveloperShares,
					AddrDerivationCostCreate: types.DefaultAddrDerivationCostCreate,
				},
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			if tc.expPanic {
				suite.Require().Panics(func() {
					feesplit.InitGenesis(suite.ctx, suite.app.FeeSplitKeeper, tc.genesis)
				})
			} else {
				suite.Require().NotPanics(func() {
					feesplit.InitGenesis(suite.ctx, suite.app.FeeSplitKeeper, tc.genesis)
				})

				params := suite.app.FeeSplitKeeper.GetParams(suite.ctx)
				suite.Require().Equal(tc.genesis.Params, params)
			}
		})
	}
}

func (suite *GenesisTestSuite) TestFeeSplitExportGenesis() {
	feesplit.InitGenesis(suite.ctx, suite.app.FeeSplitKeeper, suite.genesis)

	genesisExported := feesplit.ExportGenesis(suite.ctx, suite.app.FeeSplitKeeper)
	suite.Require().Equal(genesisExported.Params, suite.genesis.Params)
}
