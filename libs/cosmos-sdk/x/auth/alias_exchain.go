package auth

import (
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth/exported"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth/keeper"
)

type (
	Account       = exported.Account
	ModuleAccount = exported.ModuleAccount
	ObserverI     = keeper.ObserverI
)
