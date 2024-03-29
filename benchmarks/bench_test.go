package benchmarks

import (
	"encoding/hex"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/brc20-collab/brczero/libs/system"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb/opt"

	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"
	"github.com/brc20-collab/brczero/libs/tendermint/crypto/secp256k1"
	tmtypes "github.com/brc20-collab/brczero/libs/tendermint/types"
	dbm "github.com/brc20-collab/brczero/libs/tm-db"
	evmtypes "github.com/brc20-collab/brczero/x/evm/types"
	token "github.com/brc20-collab/brczero/x/token/types"
)

func BenchmarkTxSending(b *testing.B) {
	cases := []struct {
		name        string
		db          func(*testing.B) dbm.DB
		txBuilder   func(int, *AppInfo) []tmtypes.Tx
		blockSize   int
		numAccounts int
	}{
		{
			name:        "basic send - memdb",
			db:          buildMemDB,
			blockSize:   20,
			txBuilder:   buildTxFromMsg(bankSendMsg),
			numAccounts: 50,
		},
		{
			name:        "oip20 transfer - memdb",
			db:          buildMemDB,
			blockSize:   20,
			txBuilder:   buildOip20Transfer,
			numAccounts: 50,
		},
		{
			name:        "basic send - leveldb",
			db:          buildLevelDB,
			blockSize:   20,
			txBuilder:   buildTxFromMsg(bankSendMsg),
			numAccounts: 50,
		},
		{
			name:        "oip20 transfer - leveldb",
			db:          buildLevelDB,
			blockSize:   20,
			txBuilder:   buildOip20Transfer,
			numAccounts: 50,
		},
		{
			name:        "basic send - leveldb - 8k accounts",
			db:          buildLevelDB,
			blockSize:   20,
			txBuilder:   buildTxFromMsg(bankSendMsg),
			numAccounts: 8000,
		},
		{
			name:        "oip20 transfer - leveldb - 8k accounts",
			db:          buildLevelDB,
			blockSize:   20,
			txBuilder:   buildOip20Transfer,
			numAccounts: 8000,
		},
		{
			name:        "basic send - leveldb - 8k accounts - huge blocks",
			db:          buildLevelDB,
			blockSize:   1000,
			txBuilder:   buildTxFromMsg(bankSendMsg),
			numAccounts: 8000,
		},
		{
			name:        "oip20 transfer - leveldb - 8k accounts - huge blocks",
			db:          buildLevelDB,
			blockSize:   1000,
			txBuilder:   buildOip20Transfer,
			numAccounts: 8000,
		},
		{
			name:        "basic send - leveldb - 80k accounts - huge blocks",
			db:          buildLevelDB,
			blockSize:   1000,
			txBuilder:   buildTxFromMsg(bankSendMsg),
			numAccounts: 80000,
		},
		{
			name:        "oip20 transfer - leveldb - 80k accounts - huge blocks",
			db:          buildLevelDB,
			blockSize:   1000,
			txBuilder:   buildOip20Transfer,
			numAccounts: 80000,
		},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			db := tc.db(b)
			defer func() {
				_ = db.Close()
				_ = os.RemoveAll("./data")
			}()
			appInfo := InitializeOKXApp(b, db, tc.numAccounts)
			err := deployOip20(&appInfo)
			require.NoError(b, err)
			err = deployCw20(&appInfo)
			require.NoError(b, err)
			txs := tc.txBuilder(b.N, &appInfo)

			// number of Tx per block for the benchmarks
			blockSize := tc.blockSize
			height := appInfo.height + 1

			b.ResetTimer()

			for i := 0; i < b.N; {
				if i%blockSize == 0 {
					appInfo.App.BeginBlock(abci.RequestBeginBlock{Header: abci.Header{ChainID: system.Chain + "-67", Height: height, Time: time.Now()}})
				}
				//res := appInfo.App.CheckTx(abci.RequestCheckTx{
				//	Tx: txs[idx],
				//})
				//require.True(b, res.IsOK())

				res2 := appInfo.App.DeliverTx(abci.RequestDeliverTx{
					Tx: txs[i],
				})
				require.True(b, res2.IsOK())
				i++
				if i%blockSize == 0 {
					appInfo.App.EndBlock(abci.RequestEndBlock{Height: height})
					appInfo.App.Commit(abci.RequestCommit{})
					height++
				}
			}
		})
	}
}

func bankSendMsg(info *AppInfo) ([]sdk.Msg, error) {
	// Precompute all txs
	return tokenSendMsg(info)
}

func tokenSendMsg(info *AppInfo) ([]sdk.Msg, error) {
	// Precompute all txs
	rcpt := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	coins := sdk.Coins{sdk.NewInt64Coin(info.Denom, 1)}
	sendMsg := token.NewMsgTokenSend(info.MinterAddr, rcpt, coins)
	return []sdk.Msg{sendMsg}, nil
}

func buildTxFromMsg(builder func(info *AppInfo) ([]sdk.Msg, error)) func(n int, info *AppInfo) []tmtypes.Tx {
	return func(n int, info *AppInfo) []tmtypes.Tx {
		return GenSequenceOfTxs(info, builder, n)
	}
}

func buildOip20Transfer(n int, info *AppInfo) []tmtypes.Tx {
	txs := make([]tmtypes.Tx, n)
	// call oip20 transfer
	OipBytes, err := hex.DecodeString(Oip20TransferPayload)
	if err != nil {
		panic(err)
	}
	for i := range txs {
		oipTransferTx := evmtypes.NewMsgEthereumTx(info.Nonce, &info.ContractAddr, nil, GasLimit, big.NewInt(GasPrice), OipBytes)
		if err := oipTransferTx.Sign(big.NewInt(ChainId), info.evmMintKey); err != nil {
			panic(err)
		}
		info.Nonce++
		tx, err := rlp.EncodeToBytes(oipTransferTx)
		if err != nil {
			panic(err)
		}
		txs[i] = tx
	}
	return txs
}

func buildCw20Transfer(n int, info *AppInfo) []tmtypes.Tx {
	txs := make([]tmtypes.Tx, n)

	return txs
}

func buildMemDB(b *testing.B) dbm.DB {
	return dbm.NewMemDB()
}

func buildLevelDB(b *testing.B) dbm.DB {
	levelDB, err := dbm.NewGoLevelDBWithOpts("testing", b.TempDir(), &opt.Options{BlockCacher: opt.NoCacher})
	require.NoError(b, err)
	return levelDB
}
