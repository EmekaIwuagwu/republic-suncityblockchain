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
	"rsuncitychain/x/taxpayments/keeper"
	"rsuncitychain/x/taxpayments/types"
)

func createNTaxpayment(keeper keeper.Keeper, ctx context.Context, n int) []types.Taxpayment {
	items := make([]types.Taxpayment, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		_ = keeper.Taxpayment.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestTaxpaymentQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNTaxpayment(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetTaxpaymentRequest
		response *types.QueryGetTaxpaymentResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetTaxpaymentRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetTaxpaymentResponse{Taxpayment: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetTaxpaymentRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetTaxpaymentResponse{Taxpayment: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetTaxpaymentRequest{
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
			response, err := qs.GetTaxpayment(f.ctx, tc.request)
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

func TestTaxpaymentQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNTaxpayment(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllTaxpaymentRequest {
		return &types.QueryAllTaxpaymentRequest{
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
			resp, err := qs.ListTaxpayment(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Taxpayment), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Taxpayment),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListTaxpayment(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Taxpayment), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Taxpayment),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListTaxpayment(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Taxpayment),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListTaxpayment(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
