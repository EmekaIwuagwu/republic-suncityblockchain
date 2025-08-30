package types

func NewMsgCreateTaxpayment(
	creator string,
	index string,
	receiptNumber string,
	paymentFromAddress string,
	paymenttoAddress string,
	amount string,
	dateOfPayment string,
	tx string,
	createdAt string,

) *MsgCreateTaxpayment {
	return &MsgCreateTaxpayment{
		Creator:            creator,
		Index:              index,
		ReceiptNumber:      receiptNumber,
		PaymentFromAddress: paymentFromAddress,
		PaymenttoAddress:   paymenttoAddress,
		Amount:             amount,
		DateOfPayment:      dateOfPayment,
		Tx:                 tx,
		CreatedAt:          createdAt,
	}
}

func NewMsgUpdateTaxpayment(
	creator string,
	index string,
	receiptNumber string,
	paymentFromAddress string,
	paymenttoAddress string,
	amount string,
	dateOfPayment string,
	tx string,
	createdAt string,

) *MsgUpdateTaxpayment {
	return &MsgUpdateTaxpayment{
		Creator:            creator,
		Index:              index,
		ReceiptNumber:      receiptNumber,
		PaymentFromAddress: paymentFromAddress,
		PaymenttoAddress:   paymenttoAddress,
		Amount:             amount,
		DateOfPayment:      dateOfPayment,
		Tx:                 tx,
		CreatedAt:          createdAt,
	}
}

func NewMsgDeleteTaxpayment(
	creator string,
	index string,

) *MsgDeleteTaxpayment {
	return &MsgDeleteTaxpayment{
		Creator: creator,
		Index:   index,
	}
}
