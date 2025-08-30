package individualregistration

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"rsuncitychain/x/individualregistration/types"
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
					RpcMethod: "ListIndividualrecord",
					Use:       "list-individualrecord",
					Short:     "List all individualrecord",
				},
				{
					RpcMethod:      "GetIndividualrecord",
					Use:            "get-individualrecord [id]",
					Short:          "Gets a individualrecord",
					Alias:          []string{"show-individualrecord"},
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
					RpcMethod:      "CreateIndividualrecord",
					Use:            "create-individualrecord [index] [personalRegnum] [ownerAddress] [locationAddress] [dateOfBirth] [gender] [email] [tx] [telephone] [createdAt]",
					Short:          "Create a new individualrecord",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "personalRegnum"}, {ProtoField: "ownerAddress"}, {ProtoField: "locationAddress"}, {ProtoField: "dateOfBirth"}, {ProtoField: "gender"}, {ProtoField: "email"}, {ProtoField: "tx"}, {ProtoField: "telephone"}, {ProtoField: "createdAt"}},
				},
				{
					RpcMethod:      "UpdateIndividualrecord",
					Use:            "update-individualrecord [index] [personalRegnum] [ownerAddress] [locationAddress] [dateOfBirth] [gender] [email] [tx] [telephone] [createdAt]",
					Short:          "Update individualrecord",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "personalRegnum"}, {ProtoField: "ownerAddress"}, {ProtoField: "locationAddress"}, {ProtoField: "dateOfBirth"}, {ProtoField: "gender"}, {ProtoField: "email"}, {ProtoField: "tx"}, {ProtoField: "telephone"}, {ProtoField: "createdAt"}},
				},
				{
					RpcMethod:      "DeleteIndividualrecord",
					Use:            "delete-individualrecord [index]",
					Short:          "Delete individualrecord",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
