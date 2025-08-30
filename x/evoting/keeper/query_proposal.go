package keeper

import (
	"context"
	"errors"

	"rsuncitychain/x/evoting/types"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListProposal(ctx context.Context, req *types.QueryAllProposalRequest) (*types.QueryAllProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	proposals, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Proposal,
		req.Pagination,
		func(_ string, value types.Proposal) (types.Proposal, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProposalResponse{Proposal: proposals, Pagination: pageRes}, nil
}

func (q queryServer) GetProposal(ctx context.Context, req *types.QueryGetProposalRequest) (*types.QueryGetProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.Proposal.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetProposalResponse{Proposal: val}, nil
}
