package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"rsuncitychain/x/evoting/keeper"
	"rsuncitychain/x/evoting/types"
)

func SimulateMsgCreateProposal(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Int()
		msg := &types.MsgCreateProposal{
			Creator: simAccount.Address.String(),
			Index:   strconv.Itoa(i),
		}

		found, err := k.Proposal.Has(ctx, msg.Index)
		if err == nil && found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "Proposal already exist"), nil, nil
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateProposal(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			proposal   = types.Proposal{}
			msg        = &types.MsgUpdateProposal{}
			found      = false
		)

		var allProposal []types.Proposal
		err := k.Proposal.Walk(ctx, nil, func(key string, value types.Proposal) (stop bool, err error) {
			allProposal = append(allProposal, value)
			return false, nil
		})
		if err != nil {
			panic(err)
		}

		for _, obj := range allProposal {
			acc, err := ak.AddressCodec().StringToBytes(obj.Creator)
			if err != nil {
				return simtypes.OperationMsg{}, nil, err
			}

			simAccount, found = simtypes.FindAccount(accs, sdk.AccAddress(acc))
			if found {
				proposal = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "proposal creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.Index = proposal.Index

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteProposal(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			proposal   = types.Proposal{}
			msg        = &types.MsgUpdateProposal{}
			found      = false
		)

		var allProposal []types.Proposal
		err := k.Proposal.Walk(ctx, nil, func(key string, value types.Proposal) (stop bool, err error) {
			allProposal = append(allProposal, value)
			return false, nil
		})
		if err != nil {
			panic(err)
		}

		for _, obj := range allProposal {
			acc, err := ak.AddressCodec().StringToBytes(obj.Creator)
			if err != nil {
				return simtypes.OperationMsg{}, nil, err
			}

			simAccount, found = simtypes.FindAccount(accs, sdk.AccAddress(acc))
			if found {
				proposal = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "proposal creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.Index = proposal.Index

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
