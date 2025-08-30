package types

import

// DefaultGenesis returns the default genesis state
"fmt"

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(), LandrecordList: []Landrecord{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	landrecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.LandrecordList {
		index := fmt.Sprint(elem.Index)
		if _, ok := landrecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for landrecord")
		}
		landrecordIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
