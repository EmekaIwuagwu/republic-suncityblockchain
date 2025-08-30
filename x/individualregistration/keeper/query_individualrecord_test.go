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
	"rsuncitychain/x/individualregistration/keeper"
	"rsuncitychain/x/individualregistration/types"
)

func createNIndividualrecord(keeper keeper.Keeper, ctx context.Context, n int) []types.Individualrecord {
	items := make([]types.Individualrecord, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		_ = keeper.Individualrecord.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestIndividualrecordQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNIndividualrecord(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetIndividualrecordRequest
		response *types.QueryGetIndividualrecordResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetIndividualrecordRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetIndividualrecordResponse{Individualrecord: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetIndividualrecordRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetIndividualrecordResponse{Individualrecord: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetIndividualrecordRequest{
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
			response, err := qs.GetIndividualrecord(f.ctx, tc.request)
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

func TestIndividualrecordQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNIndividualrecord(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllIndividualrecordRequest {
		return &types.QueryAllIndividualrecordRequest{
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
			resp, err := qs.ListIndividualrecord(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Individualrecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Individualrecord),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListIndividualrecord(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Individualrecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Individualrecord),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListIndividualrecord(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Individualrecord),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListIndividualrecord(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
