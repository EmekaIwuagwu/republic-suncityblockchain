package keeper

import (
	"context"
	"errors"

	"rsuncitychain/x/individualregistration/types"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListIndividualrecord(ctx context.Context, req *types.QueryAllIndividualrecordRequest) (*types.QueryAllIndividualrecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	individualrecords, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Individualrecord,
		req.Pagination,
		func(_ string, value types.Individualrecord) (types.Individualrecord, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIndividualrecordResponse{Individualrecord: individualrecords, Pagination: pageRes}, nil
}

func (q queryServer) GetIndividualrecord(ctx context.Context, req *types.QueryGetIndividualrecordRequest) (*types.QueryGetIndividualrecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.Individualrecord.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetIndividualrecordResponse{Individualrecord: val}, nil
}
