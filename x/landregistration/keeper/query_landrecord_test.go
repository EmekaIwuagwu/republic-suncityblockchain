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
	"rsuncitychain/x/landregistration/keeper"
	"rsuncitychain/x/landregistration/types"
)

func createNLandrecord(keeper keeper.Keeper, ctx context.Context, n int) []types.Landrecord {
	items := make([]types.Landrecord, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		_ = keeper.Landrecord.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestLandrecordQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNLandrecord(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetLandrecordRequest
		response *types.QueryGetLandrecordResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetLandrecordRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetLandrecordResponse{Landrecord: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetLandrecordRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetLandrecordResponse{Landrecord: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetLandrecordRequest{
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
			response, err := qs.GetLandrecord(f.ctx, tc.request)
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

func TestLandrecordQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNLandrecord(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllLandrecordRequest {
		return &types.QueryAllLandrecordRequest{
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
			resp, err := qs.ListLandrecord(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Landrecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Landrecord),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListLandrecord(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Landrecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Landrecord),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListLandrecord(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Landrecord),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListLandrecord(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
