package keeper_test

import (
	"testing"

	"rsuncitychain/testutil/nullify"
	"rsuncitychain/x/taxpayments/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(), TaxpaymentList: []types.Taxpayment{{Index: "0"}, {Index: "1"}},
	}

	f := initFixture(t)
	err := f.keeper.InitGenesis(f.ctx, genesisState)
	require.NoError(t, err)
	got, err := f.keeper.ExportGenesis(f.ctx)
	require.NoError(t, err)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Params, got.Params)
	require.ElementsMatch(t, genesisState.TaxpaymentList, got.TaxpaymentList)

}
