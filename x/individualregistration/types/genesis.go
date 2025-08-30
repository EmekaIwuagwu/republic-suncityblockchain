package types

import

// DefaultGenesis returns the default genesis state
"fmt"

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(), IndividualrecordList: []Individualrecord{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	individualrecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.IndividualrecordList {
		index := fmt.Sprint(elem.Index)
		if _, ok := individualrecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for individualrecord")
		}
		individualrecordIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
