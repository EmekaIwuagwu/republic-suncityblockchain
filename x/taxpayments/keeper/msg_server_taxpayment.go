package keeper

import (
	"context"
	"errors"
	"fmt"

	"rsuncitychain/x/taxpayments/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateTaxpayment(ctx context.Context, msg *types.MsgCreateTaxpayment) (*types.MsgCreateTaxpaymentResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.Taxpayment.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var taxpayment = types.Taxpayment{
		Creator:            msg.Creator,
		Index:              msg.Index,
		ReceiptNumber:      msg.ReceiptNumber,
		PaymentFromAddress: msg.PaymentFromAddress,
		PaymenttoAddress:   msg.PaymenttoAddress,
		Amount:             msg.Amount,
		DateOfPayment:      msg.DateOfPayment,
		Tx:                 msg.Tx,
		CreatedAt:          msg.CreatedAt,
	}

	if err := k.Taxpayment.Set(ctx, taxpayment.Index, taxpayment); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateTaxpaymentResponse{}, nil
}

func (k msgServer) UpdateTaxpayment(ctx context.Context, msg *types.MsgUpdateTaxpayment) (*types.MsgUpdateTaxpaymentResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Taxpayment.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var taxpayment = types.Taxpayment{
		Creator:            msg.Creator,
		Index:              msg.Index,
		ReceiptNumber:      msg.ReceiptNumber,
		PaymentFromAddress: msg.PaymentFromAddress,
		PaymenttoAddress:   msg.PaymenttoAddress,
		Amount:             msg.Amount,
		DateOfPayment:      msg.DateOfPayment,
		Tx:                 msg.Tx,
		CreatedAt:          msg.CreatedAt,
	}

	if err := k.Taxpayment.Set(ctx, taxpayment.Index, taxpayment); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update taxpayment")
	}

	return &types.MsgUpdateTaxpaymentResponse{}, nil
}

func (k msgServer) DeleteTaxpayment(ctx context.Context, msg *types.MsgDeleteTaxpayment) (*types.MsgDeleteTaxpaymentResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Taxpayment.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Taxpayment.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove taxpayment")
	}

	return &types.MsgDeleteTaxpaymentResponse{}, nil
}
