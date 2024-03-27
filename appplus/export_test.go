package appplus

import (
	"encoding/json"
	"os"
	"testing"

	abci "github.com/cometbft/cometbft/abci/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"

	wasmapp "github.com/Finschia/wasmd/app"
	wasmplustypes "github.com/Finschia/wasmd/x/wasmplus/types"
)

func TestZeroHeightGenesis(t *testing.T) {
	db := dbm.NewMemDB()
	gapp := NewWasmApp(log.NewOCLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, DefaultNodeHome, 0, MakeEncodingConfig(), wasmplustypes.EnableAllProposals, wasmapp.EmptyBaseAppOptions{}, emptyWasmOpts)

	genesisState := NewDefaultGenesisState()
	stateBytes, err := json.MarshalIndent(genesisState, "", "  ")
	require.NoError(t, err)

	// Initialize the chain
	gapp.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	gapp.Commit()

	jailAllowedAddress := []string{"linkvaloper12kr02kew9fl73rqekalavuu0xaxcgwr6pz5vt8"}
	_, err = gapp.ExportAppStateAndValidators(true, jailAllowedAddress)
	require.NoError(t, err)
}
