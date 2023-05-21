package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "vue-test/testutil/keeper"
	"vue-test/x/vuetest/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VuetestKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
