package keeper

import (
	"zwzchain/x/extension/types"
)

var _ types.QueryServer = Keeper{}
