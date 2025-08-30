package evoting

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"rsuncitychain/testutil/sample"
	evotingsimulation "rsuncitychain/x/evoting/simulation"
	"rsuncitychain/x/evoting/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	evotingGenesis := types.GenesisState{
		Params: types.DefaultParams(), ProposalList: []types.Proposal{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}, VoteList: []types.Vote{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}},
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&evotingGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreateProposal          = "op_weight_msg_evoting"
		defaultWeightMsgCreateProposal int = 100
	)

	var weightMsgCreateProposal int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateProposal, &weightMsgCreateProposal, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProposal = defaultWeightMsgCreateProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProposal,
		evotingsimulation.SimulateMsgCreateProposal(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateProposal          = "op_weight_msg_evoting"
		defaultWeightMsgUpdateProposal int = 100
	)

	var weightMsgUpdateProposal int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateProposal, &weightMsgUpdateProposal, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProposal = defaultWeightMsgUpdateProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProposal,
		evotingsimulation.SimulateMsgUpdateProposal(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteProposal          = "op_weight_msg_evoting"
		defaultWeightMsgDeleteProposal int = 100
	)

	var weightMsgDeleteProposal int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteProposal, &weightMsgDeleteProposal, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProposal = defaultWeightMsgDeleteProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteProposal,
		evotingsimulation.SimulateMsgDeleteProposal(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCreateVote          = "op_weight_msg_evoting"
		defaultWeightMsgCreateVote int = 100
	)

	var weightMsgCreateVote int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateVote, &weightMsgCreateVote, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVote = defaultWeightMsgCreateVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVote,
		evotingsimulation.SimulateMsgCreateVote(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateVote          = "op_weight_msg_evoting"
		defaultWeightMsgUpdateVote int = 100
	)

	var weightMsgUpdateVote int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateVote, &weightMsgUpdateVote, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateVote = defaultWeightMsgUpdateVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateVote,
		evotingsimulation.SimulateMsgUpdateVote(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteVote          = "op_weight_msg_evoting"
		defaultWeightMsgDeleteVote int = 100
	)

	var weightMsgDeleteVote int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteVote, &weightMsgDeleteVote, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteVote = defaultWeightMsgDeleteVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteVote,
		evotingsimulation.SimulateMsgDeleteVote(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
