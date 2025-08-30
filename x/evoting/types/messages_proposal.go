package types

func NewMsgCreateProposal(
	creator string,
	index string,
	proposalId string,
	title string,
	description string,
	votingOptions string,
	startTime string,
	endTime string,
	status string,
	totalVotes string,
	createdAt string,

) *MsgCreateProposal {
	return &MsgCreateProposal{
		Creator:       creator,
		Index:         index,
		ProposalId:    proposalId,
		Title:         title,
		Description:   description,
		VotingOptions: votingOptions,
		StartTime:     startTime,
		EndTime:       endTime,
		Status:        status,
		TotalVotes:    totalVotes,
		CreatedAt:     createdAt,
	}
}

func NewMsgUpdateProposal(
	creator string,
	index string,
	proposalId string,
	title string,
	description string,
	votingOptions string,
	startTime string,
	endTime string,
	status string,
	totalVotes string,
	createdAt string,

) *MsgUpdateProposal {
	return &MsgUpdateProposal{
		Creator:       creator,
		Index:         index,
		ProposalId:    proposalId,
		Title:         title,
		Description:   description,
		VotingOptions: votingOptions,
		StartTime:     startTime,
		EndTime:       endTime,
		Status:        status,
		TotalVotes:    totalVotes,
		CreatedAt:     createdAt,
	}
}

func NewMsgDeleteProposal(
	creator string,
	index string,

) *MsgDeleteProposal {
	return &MsgDeleteProposal{
		Creator: creator,
		Index:   index,
	}
}
