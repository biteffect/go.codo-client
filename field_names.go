package codo

type CodoFieldName string

const (
	FieldBearerCode           CodoFieldName = "bearer_code"
	FieldInitiatorTransferId  CodoFieldName = "initiator_transfer_id"
	FieldInitiatorCallbackUrl CodoFieldName = "initiator_callback_url"
	FieldInitiatorStash       CodoFieldName = "initiator_stash"
	FieldAcceptorTransferId   CodoFieldName = "acceptor_transfer_id"
	FieldAcceptorCallbackUrl  CodoFieldName = "acceptor_callback_url"
	FieldAcceptorStash        CodoFieldName = "acceptor_stash"
	FieldFundMethod           CodoFieldName = "fund_method"
	FieldFundAmount           CodoFieldName = "fund_amount"
	FieldFundCurrency         CodoFieldName = "fund_currency"
	FieldFundCardMask         CodoFieldName = "fund_card_mask"
	FieldPayoutCardMask       CodoFieldName = "payout_card_mask"
	FieldPayoutAmount         CodoFieldName = "payout_amount"
	FieldPayoutCurrency       CodoFieldName = "payout_currency"
	FieldPayoutMethod         CodoFieldName = "payout_method"
)
