package codo_cash

type CodoFieldName string

const (
	FieldBearerCode           CodoFieldName = "bearer_code"
	FieldInitiatorTransferId  CodoFieldName = "initiator_transfer_id"
	FieldInitiatorCallbackUrl CodoFieldName = "initiator_callback_url"
	FieldInitiatorStash       CodoFieldName = "initiator_stash"
	FieldAcceptorTransferId   CodoFieldName = "acceptor_transfer_id"
	FieldAcceptorCallbackUrl  CodoFieldName = "acceptor_callback_url"
	FieldAcceptorStash        CodoFieldName = "acceptor_stash"
)

/*
	fund_method
	fund_amount
	fund_currency
	fund_card_mask
	payout_method
	payout_amount
	payout_currency
	payout_card_mask
*/
