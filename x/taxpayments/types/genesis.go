package types

import

// DefaultGenesis returns the default genesis state
"fmt"

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(), TaxpaymentList: []Taxpayment{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	taxpaymentIndexMap := make(map[string]struct{})

	for _, elem := range gs.TaxpaymentList {
		index := fmt.Sprint(elem.Index)
		if _, ok := taxpaymentIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for taxpayment")
		}
		taxpaymentIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
