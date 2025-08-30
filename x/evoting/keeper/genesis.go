package keeper

import (
	"context"

	"rsuncitychain/x/evoting/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	for _, elem := range genState.ProposalList {
		if err := k.Proposal.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}
	for _, elem := range genState.VoteList {
		if err := k.Vote.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}

	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	if err := k.Proposal.Walk(ctx, nil, func(_ string, val types.Proposal) (stop bool, err error) {
		genesis.ProposalList = append(genesis.ProposalList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	if err := k.Vote.Walk(ctx, nil, func(_ string, val types.Vote) (stop bool, err error) {
		genesis.VoteList = append(genesis.VoteList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	return genesis, nil
}
