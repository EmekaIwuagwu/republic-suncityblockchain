package keeper

import (
	"context"
	"errors"
	"fmt"

	"rsuncitychain/x/evoting/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateProposal(ctx context.Context, msg *types.MsgCreateProposal) (*types.MsgCreateProposalResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.Proposal.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var proposal = types.Proposal{
		Creator:       msg.Creator,
		Index:         msg.Index,
		ProposalId:    msg.ProposalId,
		Title:         msg.Title,
		Description:   msg.Description,
		VotingOptions: msg.VotingOptions,
		StartTime:     msg.StartTime,
		EndTime:       msg.EndTime,
		Status:        msg.Status,
		TotalVotes:    msg.TotalVotes,
		CreatedAt:     msg.CreatedAt,
	}

	if err := k.Proposal.Set(ctx, proposal.Index, proposal); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateProposalResponse{}, nil
}

func (k msgServer) UpdateProposal(ctx context.Context, msg *types.MsgUpdateProposal) (*types.MsgUpdateProposalResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Proposal.Get(ctx, msg.Index)
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

	var proposal = types.Proposal{
		Creator:       msg.Creator,
		Index:         msg.Index,
		ProposalId:    msg.ProposalId,
		Title:         msg.Title,
		Description:   msg.Description,
		VotingOptions: msg.VotingOptions,
		StartTime:     msg.StartTime,
		EndTime:       msg.EndTime,
		Status:        msg.Status,
		TotalVotes:    msg.TotalVotes,
		CreatedAt:     msg.CreatedAt,
	}

	if err := k.Proposal.Set(ctx, proposal.Index, proposal); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update proposal")
	}

	return &types.MsgUpdateProposalResponse{}, nil
}

func (k msgServer) DeleteProposal(ctx context.Context, msg *types.MsgDeleteProposal) (*types.MsgDeleteProposalResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Proposal.Get(ctx, msg.Index)
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

	if err := k.Proposal.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove proposal")
	}

	return &types.MsgDeleteProposalResponse{}, nil
}
