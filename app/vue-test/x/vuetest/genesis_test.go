package vuetest_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "vue-test/testutil/keeper"
	"vue-test/testutil/nullify"
	"vue-test/x/vuetest"
	"vue-test/x/vuetest/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VuetestKeeper(t)
	vuetest.InitGenesis(ctx, *k, genesisState)
	got := vuetest.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
