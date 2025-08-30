package types

func NewMsgCreateIndividualrecord(
	creator string,
	index string,
	personalRegnum string,
	ownerAddress string,
	locationAddress string,
	dateOfBirth string,
	gender string,
	email string,
	tx string,
	telephone string,
	createdAt string,

) *MsgCreateIndividualrecord {
	return &MsgCreateIndividualrecord{
		Creator:         creator,
		Index:           index,
		PersonalRegnum:  personalRegnum,
		OwnerAddress:    ownerAddress,
		LocationAddress: locationAddress,
		DateOfBirth:     dateOfBirth,
		Gender:          gender,
		Email:           email,
		Tx:              tx,
		Telephone:       telephone,
		CreatedAt:       createdAt,
	}
}

func NewMsgUpdateIndividualrecord(
	creator string,
	index string,
	personalRegnum string,
	ownerAddress string,
	locationAddress string,
	dateOfBirth string,
	gender string,
	email string,
	tx string,
	telephone string,
	createdAt string,

) *MsgUpdateIndividualrecord {
	return &MsgUpdateIndividualrecord{
		Creator:         creator,
		Index:           index,
		PersonalRegnum:  personalRegnum,
		OwnerAddress:    ownerAddress,
		LocationAddress: locationAddress,
		DateOfBirth:     dateOfBirth,
		Gender:          gender,
		Email:           email,
		Tx:              tx,
		Telephone:       telephone,
		CreatedAt:       createdAt,
	}
}

func NewMsgDeleteIndividualrecord(
	creator string,
	index string,

) *MsgDeleteIndividualrecord {
	return &MsgDeleteIndividualrecord{
		Creator: creator,
		Index:   index,
	}
}
