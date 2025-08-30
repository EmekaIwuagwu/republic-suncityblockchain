package taxpayments

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"rsuncitychain/x/taxpayments/types"
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
					RpcMethod: "ListTaxpayment",
					Use:       "list-taxpayment",
					Short:     "List all taxpayment",
				},
				{
					RpcMethod:      "GetTaxpayment",
					Use:            "get-taxpayment [id]",
					Short:          "Gets a taxpayment",
					Alias:          []string{"show-taxpayment"},
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
					RpcMethod:      "CreateTaxpayment",
					Use:            "create-taxpayment [index] [receiptNumber] [paymentFromAddress] [paymenttoAddress] [amount] [dateOfPayment] [tx] [createdAt]",
					Short:          "Create a new taxpayment",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "receiptNumber"}, {ProtoField: "paymentFromAddress"}, {ProtoField: "paymenttoAddress"}, {ProtoField: "amount"}, {ProtoField: "dateOfPayment"}, {ProtoField: "tx"}, {ProtoField: "createdAt"}},
				},
				{
					RpcMethod:      "UpdateTaxpayment",
					Use:            "update-taxpayment [index] [receiptNumber] [paymentFromAddress] [paymenttoAddress] [amount] [dateOfPayment] [tx] [createdAt]",
					Short:          "Update taxpayment",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "receiptNumber"}, {ProtoField: "paymentFromAddress"}, {ProtoField: "paymenttoAddress"}, {ProtoField: "amount"}, {ProtoField: "dateOfPayment"}, {ProtoField: "tx"}, {ProtoField: "createdAt"}},
				},
				{
					RpcMethod:      "DeleteTaxpayment",
					Use:            "delete-taxpayment [index]",
					Short:          "Delete taxpayment",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
