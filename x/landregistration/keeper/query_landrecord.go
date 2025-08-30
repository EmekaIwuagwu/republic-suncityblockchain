package keeper

import (
	"context"
	"errors"

	"rsuncitychain/x/landregistration/types"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListLandrecord(ctx context.Context, req *types.QueryAllLandrecordRequest) (*types.QueryAllLandrecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	landrecords, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Landrecord,
		req.Pagination,
		func(_ string, value types.Landrecord) (types.Landrecord, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLandrecordResponse{Landrecord: landrecords, Pagination: pageRes}, nil
}

func (q queryServer) GetLandrecord(ctx context.Context, req *types.QueryGetLandrecordRequest) (*types.QueryGetLandrecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.Landrecord.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetLandrecordResponse{Landrecord: val}, nil
}
