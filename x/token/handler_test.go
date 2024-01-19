package token_test

import (
	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	chain "github.com/brc20-collab/brczero/app"
	app "github.com/brc20-collab/brczero/app/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/mock"
	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"
	"github.com/brc20-collab/brczero/libs/tendermint/crypto/secp256k1"
	"github.com/brc20-collab/brczero/libs/tendermint/libs/log"
	dbm "github.com/brc20-collab/brczero/libs/tm-db"
)

// Setup initializes a new BRCZeroApp. A Nop logger is set in BRCZeroApp.
func initApp(isCheckTx bool) *chain.BRCZeroApp {
	db := dbm.NewMemDB()
	app := chain.NewBRCZeroApp(log.NewNopLogger(), db, nil, true, 0)

	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		genesisState := chain.NewDefaultGenesisState()
		stateBytes, err := codec.MarshalJSONIndent(app.Codec(), genesisState)
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		app.InitChain(
			abci.RequestInitChain{
				Validators:    []abci.ValidatorUpdate{},
				AppStateBytes: stateBytes,
			},
		)
		app.EndBlock(abci.RequestEndBlock{})
		app.Commit(abci.RequestCommit{})
	}

	return app
}

func CreateEthAccounts(numAccs int, genCoins sdk.SysCoins) (genAccs []app.EthAccount) {
	for i := 0; i < numAccs; i++ {
		privKey := secp256k1.GenPrivKey()
		pubKey := privKey.PubKey()
		addr := sdk.AccAddress(pubKey.Address())

		ak := mock.NewAddrKeys(addr, pubKey, privKey)
		testAccount := app.EthAccount{
			BaseAccount: &auth.BaseAccount{
				Address: ak.Address,
				Coins:   genCoins,
			},
			CodeHash: ethcrypto.Keccak256(nil),
		}
		genAccs = append(genAccs, testAccount)
	}
	return
}
