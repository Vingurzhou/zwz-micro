package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "ignite-test/testutil/keeper"
	"ignite-test/x/extension/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ExtensionKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
