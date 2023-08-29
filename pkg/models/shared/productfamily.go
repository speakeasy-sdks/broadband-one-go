// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ProductFamily - Indicates the product family to which product belongs to.
type ProductFamily string

const (
	ProductFamilyEthernet     ProductFamily = "Ethernet"
	ProductFamilyBroadbandOne ProductFamily = "Broadband-One"
	ProductFamilyBb1Hub       ProductFamily = "BB1Hub"
)

func (e ProductFamily) ToPointer() *ProductFamily {
	return &e
}

func (e *ProductFamily) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Ethernet":
		fallthrough
	case "Broadband-One":
		fallthrough
	case "BB1Hub":
		*e = ProductFamily(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProductFamily: %v", v)
	}
}