package keeper

import (
	"context"
	"errors"
	"fmt"

	"rsuncitychain/x/landregistration/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateLandrecord(ctx context.Context, msg *types.MsgCreateLandrecord) (*types.MsgCreateLandrecordResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.Landrecord.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var landrecord = types.Landrecord{
		Creator:             msg.Creator,
		Index:               msg.Index,
		LandRegNum:          msg.LandRegNum,
		OwnerAddress:        msg.OwnerAddress,
		LandLocationAddress: msg.LandLocationAddress,
		LandOwnerName:       msg.LandOwnerName,
		DateofLandPurchase:  msg.DateofLandPurchase,
		NameOfPreviousOwner: msg.NameOfPreviousOwner,
		LandOwnerTel:        msg.LandOwnerTel,
		LandOwnerEmail:      msg.LandOwnerEmail,
		Tx:                  msg.Tx,
		CreatedAt:           msg.CreatedAt,
	}

	if err := k.Landrecord.Set(ctx, landrecord.Index, landrecord); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateLandrecordResponse{}, nil
}

func (k msgServer) UpdateLandrecord(ctx context.Context, msg *types.MsgUpdateLandrecord) (*types.MsgUpdateLandrecordResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Landrecord.Get(ctx, msg.Index)
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

	var landrecord = types.Landrecord{
		Creator:             msg.Creator,
		Index:               msg.Index,
		LandRegNum:          msg.LandRegNum,
		OwnerAddress:        msg.OwnerAddress,
		LandLocationAddress: msg.LandLocationAddress,
		LandOwnerName:       msg.LandOwnerName,
		DateofLandPurchase:  msg.DateofLandPurchase,
		NameOfPreviousOwner: msg.NameOfPreviousOwner,
		LandOwnerTel:        msg.LandOwnerTel,
		LandOwnerEmail:      msg.LandOwnerEmail,
		Tx:                  msg.Tx,
		CreatedAt:           msg.CreatedAt,
	}

	if err := k.Landrecord.Set(ctx, landrecord.Index, landrecord); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update landrecord")
	}

	return &types.MsgUpdateLandrecordResponse{}, nil
}

func (k msgServer) DeleteLandrecord(ctx context.Context, msg *types.MsgDeleteLandrecord) (*types.MsgDeleteLandrecordResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Landrecord.Get(ctx, msg.Index)
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

	if err := k.Landrecord.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove landrecord")
	}

	return &types.MsgDeleteLandrecordResponse{}, nil
}
