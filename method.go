package codo_cash

type MethodType string

func (t MethodType) String() string {
	return string(t)
}

const (
	MethodByCode    MethodType = "by_code"
	MethodByPhone   MethodType = "by_phone"
	MethodForPerson MethodType = "for_person"
	MethodToPhone   MethodType = "to_phone"
	MethodToCard    MethodType = "to_card"
	MethodToAccount MethodType = "to_account"
	MethodDelivery  MethodType = "delivery"
)
