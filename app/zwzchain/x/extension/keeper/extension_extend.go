package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zwzchain/x/extension/types"
)

func (k Keeper) Extend(goCtx context.Context, req *types.ExtendRequest) (*types.ExtendResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	return &types.ExtendResponse{}, nil
}
