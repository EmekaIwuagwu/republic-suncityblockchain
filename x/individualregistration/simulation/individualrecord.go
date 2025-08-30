package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"rsuncitychain/x/individualregistration/keeper"
	"rsuncitychain/x/individualregistration/types"
)

func SimulateMsgCreateIndividualrecord(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Int()
		msg := &types.MsgCreateIndividualrecord{
			Creator: simAccount.Address.String(),
			Index:   strconv.Itoa(i),
		}

		found, err := k.Individualrecord.Has(ctx, msg.Index)
		if err == nil && found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "Individualrecord already exist"), nil, nil
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

func SimulateMsgUpdateIndividualrecord(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount       = simtypes.Account{}
			individualrecord = types.Individualrecord{}
			msg              = &types.MsgUpdateIndividualrecord{}
			found            = false
		)

		var allIndividualrecord []types.Individualrecord
		err := k.Individualrecord.Walk(ctx, nil, func(key string, value types.Individualrecord) (stop bool, err error) {
			allIndividualrecord = append(allIndividualrecord, value)
			return false, nil
		})
		if err != nil {
			panic(err)
		}

		for _, obj := range allIndividualrecord {
			acc, err := ak.AddressCodec().StringToBytes(obj.Creator)
			if err != nil {
				return simtypes.OperationMsg{}, nil, err
			}

			simAccount, found = simtypes.FindAccount(accs, sdk.AccAddress(acc))
			if found {
				individualrecord = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "individualrecord creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.Index = individualrecord.Index

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

func SimulateMsgDeleteIndividualrecord(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount       = simtypes.Account{}
			individualrecord = types.Individualrecord{}
			msg              = &types.MsgUpdateIndividualrecord{}
			found            = false
		)

		var allIndividualrecord []types.Individualrecord
		err := k.Individualrecord.Walk(ctx, nil, func(key string, value types.Individualrecord) (stop bool, err error) {
			allIndividualrecord = append(allIndividualrecord, value)
			return false, nil
		})
		if err != nil {
			panic(err)
		}

		for _, obj := range allIndividualrecord {
			acc, err := ak.AddressCodec().StringToBytes(obj.Creator)
			if err != nil {
				return simtypes.OperationMsg{}, nil, err
			}

			simAccount, found = simtypes.FindAccount(accs, sdk.AccAddress(acc))
			if found {
				individualrecord = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "individualrecord creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.Index = individualrecord.Index

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
