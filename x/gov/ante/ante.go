package ante

import (
	"fmt"
	ethermint "github.com/brc20-collab/brczero/app/types"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/mint"
	evmtypes "github.com/brc20-collab/brczero/x/evm/types"
	"github.com/brc20-collab/brczero/x/gov/types"
	"github.com/brc20-collab/brczero/x/params"
	paramstypes "github.com/brc20-collab/brczero/x/params/types"
	stakingkeeper "github.com/brc20-collab/brczero/x/staking/exported"
	stakingtypes "github.com/brc20-collab/brczero/x/staking/types"
	ethcmn "github.com/ethereum/go-ethereum/common"
)

type AnteDecorator struct {
	sk stakingkeeper.Keeper
	ak auth.AccountKeeper
	pk params.Keeper
}

func NewAnteDecorator(k stakingkeeper.Keeper, ak auth.AccountKeeper, pk params.Keeper) AnteDecorator {
	return AnteDecorator{sk: k, ak: ak, pk: pk}
}

func (ad AnteDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, m := range tx.GetMsgs() {
		switch msg := m.(type) {
		case types.MsgSubmitProposal:
			switch proposalType := msg.Content.(type) {
			case evmtypes.ManageContractByteCodeProposal:
				if !ad.sk.IsValidator(ctx, msg.Proposer) {
					return ctx, evmtypes.ErrCodeProposerMustBeValidator()
				}

				// check operation contract
				contract := ad.ak.GetAccount(ctx, proposalType.Contract)
				contractAcc, ok := contract.(*ethermint.EthAccount)
				if !ok || !contractAcc.IsContract() {
					return ctx, evmtypes.ErrNotContracAddress(fmt.Errorf(ethcmn.BytesToAddress(proposalType.Contract).String()))
				}

				//check substitute contract
				substitute := ad.ak.GetAccount(ctx, proposalType.SubstituteContract)
				substituteAcc, ok := substitute.(*ethermint.EthAccount)
				if !ok || !substituteAcc.IsContract() {
					return ctx, evmtypes.ErrNotContracAddress(fmt.Errorf(ethcmn.BytesToAddress(proposalType.SubstituteContract).String()))
				}
			case stakingtypes.ProposeValidatorProposal:
				if !ad.sk.IsValidator(ctx, msg.Proposer) {
					return ctx, stakingtypes.ErrCodeProposerMustBeValidator
				}
			case paramstypes.UpgradeProposal:
				if err := ad.pk.CheckMsgSubmitProposal(ctx, msg); err != nil {
					return ctx, err
				}
			case mint.ExtraProposal:
				if !ad.sk.IsValidator(ctx, msg.Proposer) {
					return ctx, mint.ErrProposerMustBeValidator
				}

			}
		}
	}

	return next(ctx, tx, simulate)
}
