package types

func NewMsgCreateVote(
	creator string,
	index string,
	voteId string,
	proposalId string,
	voterAddress string,
	selectedOption string,
	voterName string,
	timestamp string,
	tx string,
	createdAt string,

) *MsgCreateVote {
	return &MsgCreateVote{
		Creator:        creator,
		Index:          index,
		VoteId:         voteId,
		ProposalId:     proposalId,
		VoterAddress:   voterAddress,
		SelectedOption: selectedOption,
		VoterName:      voterName,
		Timestamp:      timestamp,
		Tx:             tx,
		CreatedAt:      createdAt,
	}
}

func NewMsgUpdateVote(
	creator string,
	index string,
	voteId string,
	proposalId string,
	voterAddress string,
	selectedOption string,
	voterName string,
	timestamp string,
	tx string,
	createdAt string,

) *MsgUpdateVote {
	return &MsgUpdateVote{
		Creator:        creator,
		Index:          index,
		VoteId:         voteId,
		ProposalId:     proposalId,
		VoterAddress:   voterAddress,
		SelectedOption: selectedOption,
		VoterName:      voterName,
		Timestamp:      timestamp,
		Tx:             tx,
		CreatedAt:      createdAt,
	}
}

func NewMsgDeleteVote(
	creator string,
	index string,

) *MsgDeleteVote {
	return &MsgDeleteVote{
		Creator: creator,
		Index:   index,
	}
}
