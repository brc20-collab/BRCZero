package app

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/brc20-collab/brczero/app/crypto/ethsecp256k1"
	ethermint "github.com/brc20-collab/brczero/app/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/bank"
	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"
	"github.com/brc20-collab/brczero/libs/tendermint/crypto/ed25519"
	"github.com/brc20-collab/brczero/libs/tendermint/crypto/secp256k1"
	"github.com/brc20-collab/brczero/x/distribution/keeper"
	"github.com/brc20-collab/brczero/x/evm"
	evm_types "github.com/brc20-collab/brczero/x/evm/types"
	"github.com/brc20-collab/brczero/x/staking"
	staking_keeper "github.com/brc20-collab/brczero/x/staking/keeper"
	staking_types "github.com/brc20-collab/brczero/x/staking/types"

	"github.com/stretchr/testify/suite"
)

var (
	coin10  = sdk.NewInt64Coin(sdk.DefaultBondDenom, 10)
	coin20  = sdk.NewInt64Coin(sdk.DefaultBondDenom, 20)
	coin30  = sdk.NewInt64Coin(sdk.DefaultBondDenom, 30)
	coin40  = sdk.NewInt64Coin(sdk.DefaultBondDenom, 40)
	coin50  = sdk.NewInt64Coin(sdk.DefaultBondDenom, 50)
	coin60  = sdk.NewInt64Coin(sdk.DefaultBondDenom, 60)
	coin70  = sdk.NewInt64Coin(sdk.DefaultBondDenom, 70)
	coin80  = sdk.NewInt64Coin(sdk.DefaultBondDenom, 80)
	coin90  = sdk.NewInt64Coin(sdk.DefaultBondDenom, 90)
	coin100 = sdk.NewInt64Coin(sdk.DefaultBondDenom, 100)
	fees    = auth.NewStdFee(21000, sdk.NewCoins(coin10))
)

type InnerTxTestSuite struct {
	suite.Suite

	ctx     sdk.Context
	app     *BRCZeroApp
	stateDB *evm_types.CommitStateDB
	codec   *codec.Codec

	handler sdk.Handler
}

// Note: DefaultMinSelfDelegation was changed to 0 from 10000
func (suite *InnerTxTestSuite) SetupTest() {
	checkTx := false
	chain_id := "ethermint-3"

	suite.app = Setup(checkTx)
	suite.ctx = suite.app.BaseApp.NewContext(checkTx, abci.Header{Height: 1, ChainID: chain_id, Time: time.Now().UTC()})
	suite.ctx.SetDeliverSerial()
	suite.stateDB = evm_types.CreateEmptyCommitStateDB(suite.app.EvmKeeper.GenerateCSDBParams(), suite.ctx)
	suite.codec = codec.New()

	err := ethermint.SetChainId(chain_id)
	suite.Nil(err)

	params := evm_types.DefaultParams()
	params.EnableCreate = true
	params.EnableCall = true
	suite.app.EvmKeeper.SetParams(suite.ctx, params)

	stakingParams := staking_types.DefaultDposParams()
	suite.app.StakingKeeper.SetParams(suite.ctx, stakingParams)
}

func TestInnerTxTestSuite(t *testing.T) {
	suite.Run(t, new(InnerTxTestSuite))
}

func (suite *InnerTxTestSuite) TestMsgSend() {
	var (
		tx          sdk.Tx
		privFrom, _ = ethsecp256k1.GenerateKey()
		//ethFrom     = common.HexToAddress(privFrom.PubKey().Address().String())
		cmFrom = sdk.AccAddress(privFrom.PubKey().Address())
		privTo = secp256k1.GenPrivKeySecp256k1([]byte("private key to"))
		ethTo  = common.HexToAddress(privTo.PubKey().Address().String())
		cmTo   = sdk.AccAddress(privTo.PubKey().Address())

		valPriv      = ed25519.GenPrivKeyFromSecret([]byte("ed25519 private key"))
		valpub       = valPriv.PubKey()
		valopaddress = sdk.ValAddress(valpub.Address())
		valcmaddress = sdk.AccAddress(valpub.Address())

		privFrom1 = secp256k1.GenPrivKeySecp256k1([]byte("from1"))
		cmFrom1   = sdk.AccAddress(privFrom1.PubKey().Address())
		privTo1   = secp256k1.GenPrivKeySecp256k1([]byte("to1"))
		cmTo1     = sdk.AccAddress(privTo1.PubKey().Address())
	)
	normal := func() {
		err := suite.app.BankKeeper.SetCoins(suite.ctx, cmFrom, sdk.NewCoins(coin100))
		suite.Require().NoError(err)
	}
	testCases := []struct {
		msg        string
		prepare    func()
		expPass    bool
		expectfunc func()
	}{
		{
			"send msg(bank)",
			func() {
				suite.handler = bank.NewHandler(suite.app.BankKeeper)

				msg := bank.NewMsgSend(cmFrom, cmTo, sdk.NewCoins(coin10))
				tx = auth.NewStdTx([]sdk.Msg{msg}, fees, nil, "")
			},
			true,
			func() {
				fromBalance := suite.app.AccountKeeper.GetAccount(suite.ctx, cmFrom).GetCoins()
				suite.Require().True(fromBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(coin90))))

				toBalance := suite.app.AccountKeeper.GetAccount(suite.ctx, cmTo).GetCoins()
				suite.Require().True(toBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(coin10))))
			},
		},
		{
			"send msgs(bank)",
			func() {
				suite.handler = bank.NewHandler(suite.app.BankKeeper)

				msg := bank.NewMsgSend(cmFrom, cmTo, sdk.NewCoins(coin10))
				tx = auth.NewStdTx([]sdk.Msg{msg, msg}, fees, nil, "")
			},
			true,
			func() {
				fromBalance := suite.app.AccountKeeper.GetAccount(suite.ctx, cmFrom).GetCoins()
				suite.Require().True(fromBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(coin80))))

				toBalance := suite.app.AccountKeeper.GetAccount(suite.ctx, cmTo).GetCoins()
				suite.Require().True(toBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(coin20))))
			},
		},
		{
			"multi msg(bank)",
			func() {
				suite.handler = bank.NewHandler(suite.app.BankKeeper)
				suite.app.BankKeeper.SetCoins(suite.ctx, cmFrom, sdk.NewCoins(coin100))
				suite.app.BankKeeper.SetCoins(suite.ctx, cmFrom1, sdk.NewCoins(coin100))
				inputCoin1 := sdk.NewCoins(coin20)
				inputCoin2 := sdk.NewCoins(coin10)
				outputCoin1 := sdk.NewCoins(coin10)
				outputCoin2 := sdk.NewCoins(coin20)
				input1 := bank.NewInput(cmFrom, inputCoin1)
				input2 := bank.NewInput(cmFrom1, inputCoin2)
				output1 := bank.NewOutput(cmTo, outputCoin1)
				output2 := bank.NewOutput(cmTo1, outputCoin2)

				msg := bank.NewMsgMultiSend([]bank.Input{input1, input2}, []bank.Output{output1, output2})
				tx = auth.NewStdTx([]sdk.Msg{msg}, fees, nil, "")
			},
			true,
			func() {
				fromBalance := suite.app.AccountKeeper.GetAccount(suite.ctx, cmFrom).GetCoins()
				suite.Require().True(fromBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(coin80))))
				fromBalance = suite.app.AccountKeeper.GetAccount(suite.ctx, cmFrom1).GetCoins()
				suite.Require().True(fromBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(coin90))))

				toBalance := suite.app.AccountKeeper.GetAccount(suite.ctx, cmTo).GetCoins()
				suite.Require().True(toBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(coin10))))
				toBalance = suite.app.AccountKeeper.GetAccount(suite.ctx, cmTo1).GetCoins()
				suite.Require().True(toBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(coin20))))
			},
		},
		{
			"evm send msg(evm)",
			func() {
				suite.handler = evm.NewHandler(suite.app.EvmKeeper)
				tx = evm_types.NewMsgEthereumTx(0, &ethTo, coin10.Amount.BigInt(), 3000000, big.NewInt(0), nil)

				// parse context chain ID to big.Int
				chainID, err := ethermint.ParseChainID(suite.ctx.ChainID())
				suite.Require().NoError(err)

				// sign transaction
				ethTx, ok := tx.(*evm_types.MsgEthereumTx)
				suite.Require().True(ok)

				err = ethTx.Sign(chainID, privFrom.ToECDSA())
				suite.Require().NoError(err)
				tx = ethTx
			},
			true,
			func() {
				fromBalance := suite.app.AccountKeeper.GetAccount(suite.ctx, cmFrom).GetCoins()
				suite.Require().True(fromBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(coin90))))

				toBalance := suite.app.AccountKeeper.GetAccount(suite.ctx, cmTo).GetCoins()
				suite.Require().True(toBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(coin10))))
			},
		},
		{
			"create validator(staking)",
			func() {
				suite.handler = staking.NewHandler(suite.app.StakingKeeper)

				err := suite.app.BankKeeper.SetCoins(suite.ctx, valcmaddress, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 20000)))
				suite.Require().NoError(err)

				msg := staking_keeper.NewTestMsgCreateValidator(valopaddress, valpub, coin10.Amount)
				tx = auth.NewStdTx([]sdk.Msg{msg}, fees, nil, "")
			},
			true,
			func() {
				fromBalance := suite.app.AccountKeeper.GetAccount(suite.ctx, valcmaddress).GetCoins()
				suite.Require().True(fromBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000)))))

				suite.app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(suite.ctx)
				val, ok := suite.app.StakingKeeper.GetValidator(suite.ctx, valopaddress)
				suite.Require().True(ok)
				suite.Require().Equal(valopaddress, val.OperatorAddress)
				suite.Require().True(val.MinSelfDelegation.Equal(sdk.NewDec(10000)))
			},
		},
		{
			"destroy validator(staking)",
			func() {
				suite.handler = staking.NewHandler(suite.app.StakingKeeper)

				err := suite.app.BankKeeper.SetCoins(suite.ctx, valcmaddress, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 20000)))
				suite.Require().NoError(err)

				msg := staking_keeper.NewTestMsgCreateValidator(valopaddress, valpub, coin10.Amount)

				destroyValMsg := staking_types.NewMsgDestroyValidator([]byte(valopaddress))
				tx = auth.NewStdTx([]sdk.Msg{msg, destroyValMsg}, fees, nil, "")
			},
			true,
			func() {
				fromBalance := suite.app.AccountKeeper.GetAccount(suite.ctx, valcmaddress).GetCoins()
				suite.Require().True(fromBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000)))))

				suite.app.EndBlocker(suite.ctx.WithBlockTime(time.Now().Add(staking_types.DefaultUnbondingTime)), abci.RequestEndBlock{Height: 2})
				_, ok := suite.app.StakingKeeper.GetValidator(suite.ctx, valopaddress)
				suite.Require().False(ok)
				fromBalance = suite.app.AccountKeeper.GetAccount(suite.ctx, valcmaddress).GetCoins()
				suite.Require().True(fromBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(sdk.NewInt64Coin(sdk.DefaultBondDenom, 20000)))))

			},
		},
		{
			"deposit msg(staking)",
			func() {
				suite.handler = staking.NewHandler(suite.app.StakingKeeper)
				err := suite.app.BankKeeper.SetCoins(suite.ctx, cmFrom, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000)))
				suite.Require().NoError(err)
				err = suite.app.BankKeeper.SetCoins(suite.ctx, valcmaddress, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 20000)))
				suite.Require().NoError(err)

				msg := staking_keeper.NewTestMsgCreateValidator(valopaddress, valpub, coin10.Amount)

				depositMsg := staking_types.NewMsgDeposit(cmFrom, keeper.NewTestSysCoin(10000, 0))
				tx = auth.NewStdTx([]sdk.Msg{msg, depositMsg}, fees, nil, "")
			},
			true,
			func() {
				fromBalance := suite.app.AccountKeeper.GetAccount(suite.ctx, cmFrom).GetCoins()
				suite.Require().True(fromBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)))))

			},
		},
		{
			"withdraw msg(staking)",
			func() {
				suite.handler = staking.NewHandler(suite.app.StakingKeeper)
				err := suite.app.BankKeeper.SetCoins(suite.ctx, cmFrom, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000)))
				suite.Require().NoError(err)
				err = suite.app.BankKeeper.SetCoins(suite.ctx, valcmaddress, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 20000)))
				suite.Require().NoError(err)

				msg := staking_keeper.NewTestMsgCreateValidator(valopaddress, valpub, coin10.Amount)

				depositMsg := staking_types.NewMsgDeposit(cmFrom, keeper.NewTestSysCoin(10000, 0))

				withdrawMsg := staking_types.NewMsgWithdraw(cmFrom, keeper.NewTestSysCoin(10000, 0))
				tx = auth.NewStdTx([]sdk.Msg{msg, depositMsg, withdrawMsg}, fees, nil, "")
			},
			true,
			func() {
				fromBalance := suite.app.AccountKeeper.GetAccount(suite.ctx, cmFrom).GetCoins()
				suite.Require().True(fromBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)))))
				suite.app.EndBlocker(suite.ctx.WithBlockTime(time.Now().Add(staking_types.DefaultUnbondingTime)), abci.RequestEndBlock{Height: 2})
				fromBalance = suite.app.AccountKeeper.GetAccount(suite.ctx, cmFrom).GetCoins()
				suite.Require().True(fromBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000)))))
			},
		},
		{
			"addshare msg(staking)",
			func() {
				suite.handler = staking.NewHandler(suite.app.StakingKeeper)
				err := suite.app.BankKeeper.SetCoins(suite.ctx, cmFrom, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000)))
				suite.Require().NoError(err)
				err = suite.app.BankKeeper.SetCoins(suite.ctx, valcmaddress, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 20000)))
				suite.Require().NoError(err)

				msg := staking_keeper.NewTestMsgCreateValidator(valopaddress, valpub, coin10.Amount)

				depositMsg := staking_types.NewMsgDeposit(cmFrom, keeper.NewTestSysCoin(10000, 0))
				addShareMsg := staking_types.NewMsgAddShares(cmFrom, []sdk.ValAddress{valopaddress})
				tx = auth.NewStdTx([]sdk.Msg{msg, depositMsg, addShareMsg}, fees, nil, "")
			},
			true,
			func() {
				fromBalance := suite.app.AccountKeeper.GetAccount(suite.ctx, cmFrom).GetCoins()
				suite.Require().True(fromBalance.IsEqual(sdk.NewDecCoins(sdk.NewDecCoinFromCoin(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)))))

			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.msg, func() {
			suite.SetupTest() // reset
			normal()
			//nolint
			tc.prepare()
			suite.ctx.SetGasMeter(sdk.NewInfiniteGasMeter())
			msgs := tx.GetMsgs()
			for _, msg := range msgs {
				_, err := suite.handler(suite.ctx, msg)

				//nolint
				if tc.expPass {
					suite.Require().NoError(err)
				} else {
					suite.Require().Error(err)
				}
			}
			tc.expectfunc()
		})
	}
}
