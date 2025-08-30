package types_test

import (
	"testing"

	"rsuncitychain/x/evoting/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc:     "valid genesis state",
			genState: &types.GenesisState{ProposalList: []types.Proposal{{Index: "0"}, {Index: "1"}}, VoteList: []types.Vote{{Index: "0"}, {Index: "1"}}},
			valid:    true,
		}, {desc: "duplicated proposal",

			genState: &types.GenesisState{ProposalList: []types.Proposal{

				{
					Index: "0"}, {Index: "0"}}, VoteList: []types.Vote{{Index: "0"}, {Index: "1"}},
			}, valid: false,
		}, {desc: "duplicated vote",

			genState: &types.GenesisState{

				VoteList: []types.Vote{{Index: "0"}, {Index: "0"}}}, valid: false},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
