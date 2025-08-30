package keeper

import (
	"context"
	"errors"

	"rsuncitychain/x/taxpayments/types"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListTaxpayment(ctx context.Context, req *types.QueryAllTaxpaymentRequest) (*types.QueryAllTaxpaymentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	taxpayments, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Taxpayment,
		req.Pagination,
		func(_ string, value types.Taxpayment) (types.Taxpayment, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTaxpaymentResponse{Taxpayment: taxpayments, Pagination: pageRes}, nil
}

func (q queryServer) GetTaxpayment(ctx context.Context, req *types.QueryGetTaxpaymentRequest) (*types.QueryGetTaxpaymentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.Taxpayment.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetTaxpaymentResponse{Taxpayment: val}, nil
}
