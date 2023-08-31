package types

import dbm "github.com/brc20-collab/brczero/libs/tm-db"

// DBBackend This is set at compile time.
var DBBackend = string(dbm.GoLevelDBBackend)
