package evoting

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"rsuncitychain/x/evoting/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ListProposal",
					Use:       "list-proposal",
					Short:     "List all proposal",
				},
				{
					RpcMethod:      "GetProposal",
					Use:            "get-proposal [id]",
					Short:          "Gets a proposal",
					Alias:          []string{"show-proposal"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ListVote",
					Use:       "list-vote",
					Short:     "List all vote",
				},
				{
					RpcMethod:      "GetVote",
					Use:            "get-vote [id]",
					Short:          "Gets a vote",
					Alias:          []string{"show-vote"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateProposal",
					Use:            "create-proposal [index] [proposalId] [title] [description] [votingOptions] [startTime] [endTime] [status] [totalVotes] [createdAt]",
					Short:          "Create a new proposal",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "proposalId"}, {ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "votingOptions"}, {ProtoField: "startTime"}, {ProtoField: "endTime"}, {ProtoField: "status"}, {ProtoField: "totalVotes"}, {ProtoField: "createdAt"}},
				},
				{
					RpcMethod:      "UpdateProposal",
					Use:            "update-proposal [index] [proposalId] [title] [description] [votingOptions] [startTime] [endTime] [status] [totalVotes] [createdAt]",
					Short:          "Update proposal",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "proposalId"}, {ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "votingOptions"}, {ProtoField: "startTime"}, {ProtoField: "endTime"}, {ProtoField: "status"}, {ProtoField: "totalVotes"}, {ProtoField: "createdAt"}},
				},
				{
					RpcMethod:      "DeleteProposal",
					Use:            "delete-proposal [index]",
					Short:          "Delete proposal",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateVote",
					Use:            "create-vote [index] [voteId] [proposalId] [voterAddress] [selectedOption] [voterName] [timestamp] [tx] [createdAt]",
					Short:          "Create a new vote",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "voteId"}, {ProtoField: "proposalId"}, {ProtoField: "voterAddress"}, {ProtoField: "selectedOption"}, {ProtoField: "voterName"}, {ProtoField: "timestamp"}, {ProtoField: "tx"}, {ProtoField: "createdAt"}},
				},
				{
					RpcMethod:      "UpdateVote",
					Use:            "update-vote [index] [voteId] [proposalId] [voterAddress] [selectedOption] [voterName] [timestamp] [tx] [createdAt]",
					Short:          "Update vote",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "voteId"}, {ProtoField: "proposalId"}, {ProtoField: "voterAddress"}, {ProtoField: "selectedOption"}, {ProtoField: "voterName"}, {ProtoField: "timestamp"}, {ProtoField: "tx"}, {ProtoField: "createdAt"}},
				},
				{
					RpcMethod:      "DeleteVote",
					Use:            "delete-vote [index]",
					Short:          "Delete vote",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
