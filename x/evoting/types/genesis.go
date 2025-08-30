package types

import

// DefaultGenesis returns the default genesis state
"fmt"

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(), ProposalList: []Proposal{}, VoteList: []Vote{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	proposalIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProposalList {
		index := fmt.Sprint(elem.Index)
		if _, ok := proposalIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for proposal")
		}
		proposalIndexMap[index] = struct{}{}
	}
	voteIndexMap := make(map[string]struct{})

	for _, elem := range gs.VoteList {
		index := fmt.Sprint(elem.Index)
		if _, ok := voteIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for vote")
		}
		voteIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
