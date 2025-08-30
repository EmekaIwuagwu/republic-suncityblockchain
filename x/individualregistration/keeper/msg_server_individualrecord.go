package keeper

import (
	"context"
	"errors"
	"fmt"

	"rsuncitychain/x/individualregistration/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateIndividualrecord(ctx context.Context, msg *types.MsgCreateIndividualrecord) (*types.MsgCreateIndividualrecordResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.Individualrecord.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var individualrecord = types.Individualrecord{
		Creator:         msg.Creator,
		Index:           msg.Index,
		PersonalRegnum:  msg.PersonalRegnum,
		OwnerAddress:    msg.OwnerAddress,
		LocationAddress: msg.LocationAddress,
		DateOfBirth:     msg.DateOfBirth,
		Gender:          msg.Gender,
		Email:           msg.Email,
		Tx:              msg.Tx,
		Telephone:       msg.Telephone,
		CreatedAt:       msg.CreatedAt,
	}

	if err := k.Individualrecord.Set(ctx, individualrecord.Index, individualrecord); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateIndividualrecordResponse{}, nil
}

func (k msgServer) UpdateIndividualrecord(ctx context.Context, msg *types.MsgUpdateIndividualrecord) (*types.MsgUpdateIndividualrecordResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Individualrecord.Get(ctx, msg.Index)
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

	var individualrecord = types.Individualrecord{
		Creator:         msg.Creator,
		Index:           msg.Index,
		PersonalRegnum:  msg.PersonalRegnum,
		OwnerAddress:    msg.OwnerAddress,
		LocationAddress: msg.LocationAddress,
		DateOfBirth:     msg.DateOfBirth,
		Gender:          msg.Gender,
		Email:           msg.Email,
		Tx:              msg.Tx,
		Telephone:       msg.Telephone,
		CreatedAt:       msg.CreatedAt,
	}

	if err := k.Individualrecord.Set(ctx, individualrecord.Index, individualrecord); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update individualrecord")
	}

	return &types.MsgUpdateIndividualrecordResponse{}, nil
}

func (k msgServer) DeleteIndividualrecord(ctx context.Context, msg *types.MsgDeleteIndividualrecord) (*types.MsgDeleteIndividualrecordResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Individualrecord.Get(ctx, msg.Index)
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

	if err := k.Individualrecord.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove individualrecord")
	}

	return &types.MsgDeleteIndividualrecordResponse{}, nil
}
