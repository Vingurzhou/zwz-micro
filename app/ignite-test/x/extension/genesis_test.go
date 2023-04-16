package extension_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ignite-test/testutil/keeper"
	"ignite-test/testutil/nullify"
	"ignite-test/x/extension"
	"ignite-test/x/extension/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ExtensionKeeper(t)
	extension.InitGenesis(ctx, *k, genesisState)
	got := extension.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
