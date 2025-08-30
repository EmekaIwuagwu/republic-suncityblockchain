package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"rsuncitychain/testutil/nullify"
	"rsuncitychain/x/evoting/keeper"
	"rsuncitychain/x/evoting/types"
)

func createNProposal(keeper keeper.Keeper, ctx context.Context, n int) []types.Proposal {
	items := make([]types.Proposal, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		_ = keeper.Proposal.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestProposalQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNProposal(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetProposalRequest
		response *types.QueryGetProposalResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetProposalRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetProposalResponse{Proposal: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetProposalRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetProposalResponse{Proposal: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetProposalRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetProposal(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestProposalQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNProposal(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllProposalRequest {
		return &types.QueryAllProposalRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListProposal(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Proposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Proposal),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListProposal(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Proposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Proposal),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListProposal(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Proposal),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListProposal(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
