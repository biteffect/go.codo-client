package codo

import (
	"fmt"
	"strings"
)

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

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *MethodType) UnmarshalJSON(bytes []byte) error {
	str := strings.Trim(string(bytes), "\" ")
	str = strings.ToLower(str)
	for _, v := range TransferMethods() {
		if string(v) == str {
			*t = v
			return nil
		}
	}
	return fmt.Errorf("unsupported transfer method: %v", str)
}

func TransferMethods() []MethodType {
	return []MethodType{
		MethodByCode,
		MethodByPhone,
		MethodForPerson,
		MethodToPhone,
		MethodToCard,
		MethodToAccount,
		MethodDelivery,
	}
}
