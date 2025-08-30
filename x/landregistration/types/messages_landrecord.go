package types

func NewMsgCreateLandrecord(
	creator string,
	index string,
	landRegNum string,
	ownerAddress string,
	landLocationAddress string,
	landOwnerName string,
	dateofLandPurchase string,
	nameOfPreviousOwner string,
	landOwnerTel string,
	landOwnerEmail string,
	tx string,
	createdAt string,

) *MsgCreateLandrecord {
	return &MsgCreateLandrecord{
		Creator:             creator,
		Index:               index,
		LandRegNum:          landRegNum,
		OwnerAddress:        ownerAddress,
		LandLocationAddress: landLocationAddress,
		LandOwnerName:       landOwnerName,
		DateofLandPurchase:  dateofLandPurchase,
		NameOfPreviousOwner: nameOfPreviousOwner,
		LandOwnerTel:        landOwnerTel,
		LandOwnerEmail:      landOwnerEmail,
		Tx:                  tx,
		CreatedAt:           createdAt,
	}
}

func NewMsgUpdateLandrecord(
	creator string,
	index string,
	landRegNum string,
	ownerAddress string,
	landLocationAddress string,
	landOwnerName string,
	dateofLandPurchase string,
	nameOfPreviousOwner string,
	landOwnerTel string,
	landOwnerEmail string,
	tx string,
	createdAt string,

) *MsgUpdateLandrecord {
	return &MsgUpdateLandrecord{
		Creator:             creator,
		Index:               index,
		LandRegNum:          landRegNum,
		OwnerAddress:        ownerAddress,
		LandLocationAddress: landLocationAddress,
		LandOwnerName:       landOwnerName,
		DateofLandPurchase:  dateofLandPurchase,
		NameOfPreviousOwner: nameOfPreviousOwner,
		LandOwnerTel:        landOwnerTel,
		LandOwnerEmail:      landOwnerEmail,
		Tx:                  tx,
		CreatedAt:           createdAt,
	}
}

func NewMsgDeleteLandrecord(
	creator string,
	index string,

) *MsgDeleteLandrecord {
	return &MsgDeleteLandrecord{
		Creator: creator,
		Index:   index,
	}
}
