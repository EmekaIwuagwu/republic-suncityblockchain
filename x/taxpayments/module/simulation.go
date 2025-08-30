package taxpayments

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"rsuncitychain/testutil/sample"
	taxpaymentssimulation "rsuncitychain/x/taxpayments/simulation"
	"rsuncitychain/x/taxpayments/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	taxpaymentsGenesis := types.GenesisState{
		Params: types.DefaultParams(), TaxpaymentList: []types.Taxpayment{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}},
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&taxpaymentsGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreateTaxpayment          = "op_weight_msg_taxpayments"
		defaultWeightMsgCreateTaxpayment int = 100
	)

	var weightMsgCreateTaxpayment int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateTaxpayment, &weightMsgCreateTaxpayment, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTaxpayment = defaultWeightMsgCreateTaxpayment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTaxpayment,
		taxpaymentssimulation.SimulateMsgCreateTaxpayment(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateTaxpayment          = "op_weight_msg_taxpayments"
		defaultWeightMsgUpdateTaxpayment int = 100
	)

	var weightMsgUpdateTaxpayment int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateTaxpayment, &weightMsgUpdateTaxpayment, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTaxpayment = defaultWeightMsgUpdateTaxpayment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTaxpayment,
		taxpaymentssimulation.SimulateMsgUpdateTaxpayment(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteTaxpayment          = "op_weight_msg_taxpayments"
		defaultWeightMsgDeleteTaxpayment int = 100
	)

	var weightMsgDeleteTaxpayment int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteTaxpayment, &weightMsgDeleteTaxpayment, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTaxpayment = defaultWeightMsgDeleteTaxpayment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTaxpayment,
		taxpaymentssimulation.SimulateMsgDeleteTaxpayment(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
