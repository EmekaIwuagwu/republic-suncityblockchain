package individualregistration

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"rsuncitychain/testutil/sample"
	individualregistrationsimulation "rsuncitychain/x/individualregistration/simulation"
	"rsuncitychain/x/individualregistration/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	individualregistrationGenesis := types.GenesisState{
		Params: types.DefaultParams(), IndividualrecordList: []types.Individualrecord{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}},
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&individualregistrationGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreateIndividualrecord          = "op_weight_msg_individualregistration"
		defaultWeightMsgCreateIndividualrecord int = 100
	)

	var weightMsgCreateIndividualrecord int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateIndividualrecord, &weightMsgCreateIndividualrecord, nil,
		func(_ *rand.Rand) {
			weightMsgCreateIndividualrecord = defaultWeightMsgCreateIndividualrecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateIndividualrecord,
		individualregistrationsimulation.SimulateMsgCreateIndividualrecord(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateIndividualrecord          = "op_weight_msg_individualregistration"
		defaultWeightMsgUpdateIndividualrecord int = 100
	)

	var weightMsgUpdateIndividualrecord int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateIndividualrecord, &weightMsgUpdateIndividualrecord, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateIndividualrecord = defaultWeightMsgUpdateIndividualrecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateIndividualrecord,
		individualregistrationsimulation.SimulateMsgUpdateIndividualrecord(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteIndividualrecord          = "op_weight_msg_individualregistration"
		defaultWeightMsgDeleteIndividualrecord int = 100
	)

	var weightMsgDeleteIndividualrecord int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteIndividualrecord, &weightMsgDeleteIndividualrecord, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteIndividualrecord = defaultWeightMsgDeleteIndividualrecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteIndividualrecord,
		individualregistrationsimulation.SimulateMsgDeleteIndividualrecord(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
