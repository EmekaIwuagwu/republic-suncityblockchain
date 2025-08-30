package landregistration

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"rsuncitychain/x/landregistration/types"
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
					RpcMethod: "ListLandrecord",
					Use:       "list-landrecord",
					Short:     "List all landrecord",
				},
				{
					RpcMethod:      "GetLandrecord",
					Use:            "get-landrecord [id]",
					Short:          "Gets a landrecord",
					Alias:          []string{"show-landrecord"},
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
					RpcMethod:      "CreateLandrecord",
					Use:            "create-landrecord [index] [landRegNum] [ownerAddress] [landLocationAddress] [landOwnerName] [dateofLandPurchase] [nameOfPreviousOwner] [landOwnerTel] [landOwnerEmail] [tx] [createdAt]",
					Short:          "Create a new landrecord",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "landRegNum"}, {ProtoField: "ownerAddress"}, {ProtoField: "landLocationAddress"}, {ProtoField: "landOwnerName"}, {ProtoField: "dateofLandPurchase"}, {ProtoField: "nameOfPreviousOwner"}, {ProtoField: "landOwnerTel"}, {ProtoField: "landOwnerEmail"}, {ProtoField: "tx"}, {ProtoField: "createdAt"}},
				},
				{
					RpcMethod:      "UpdateLandrecord",
					Use:            "update-landrecord [index] [landRegNum] [ownerAddress] [landLocationAddress] [landOwnerName] [dateofLandPurchase] [nameOfPreviousOwner] [landOwnerTel] [landOwnerEmail] [tx] [createdAt]",
					Short:          "Update landrecord",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "landRegNum"}, {ProtoField: "ownerAddress"}, {ProtoField: "landLocationAddress"}, {ProtoField: "landOwnerName"}, {ProtoField: "dateofLandPurchase"}, {ProtoField: "nameOfPreviousOwner"}, {ProtoField: "landOwnerTel"}, {ProtoField: "landOwnerEmail"}, {ProtoField: "tx"}, {ProtoField: "createdAt"}},
				},
				{
					RpcMethod:      "DeleteLandrecord",
					Use:            "delete-landrecord [index]",
					Short:          "Delete landrecord",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
