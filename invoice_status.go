package codo

type InvoiceStatus string

func (s InvoiceStatus) String() string {
	return string(s)
}

const (
	InvStatusDraft     InvoiceStatus = "draft"
	InvStatusPrepared  InvoiceStatus = "prepared"
	InvStatusAccepted  InvoiceStatus = "accepted"
	InvStatusFail      InvoiceStatus = "fail"
	InvStatusActive    InvoiceStatus = "active"
	InvStatusHold      InvoiceStatus = "hold"
	InvStatusFulfilled InvoiceStatus = "fulfilled"
)

func InvoiceStatuses() []InvoiceStatus {
	return []InvoiceStatus{
		InvStatusDraft,
		InvStatusPrepared,
		InvStatusAccepted,
		InvStatusFail,
		InvStatusActive,
		InvStatusHold,
		InvStatusFulfilled,
	}
}
