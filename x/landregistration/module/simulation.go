package landregistration

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"rsuncitychain/testutil/sample"
	landregistrationsimulation "rsuncitychain/x/landregistration/simulation"
	"rsuncitychain/x/landregistration/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	landregistrationGenesis := types.GenesisState{
		Params: types.DefaultParams(), LandrecordList: []types.Landrecord{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}},
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&landregistrationGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreateLandrecord          = "op_weight_msg_landregistration"
		defaultWeightMsgCreateLandrecord int = 100
	)

	var weightMsgCreateLandrecord int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateLandrecord, &weightMsgCreateLandrecord, nil,
		func(_ *rand.Rand) {
			weightMsgCreateLandrecord = defaultWeightMsgCreateLandrecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateLandrecord,
		landregistrationsimulation.SimulateMsgCreateLandrecord(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateLandrecord          = "op_weight_msg_landregistration"
		defaultWeightMsgUpdateLandrecord int = 100
	)

	var weightMsgUpdateLandrecord int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateLandrecord, &weightMsgUpdateLandrecord, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateLandrecord = defaultWeightMsgUpdateLandrecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateLandrecord,
		landregistrationsimulation.SimulateMsgUpdateLandrecord(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteLandrecord          = "op_weight_msg_landregistration"
		defaultWeightMsgDeleteLandrecord int = 100
	)

	var weightMsgDeleteLandrecord int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteLandrecord, &weightMsgDeleteLandrecord, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteLandrecord = defaultWeightMsgDeleteLandrecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteLandrecord,
		landregistrationsimulation.SimulateMsgDeleteLandrecord(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
