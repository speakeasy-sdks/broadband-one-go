// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PostAppointmentRequestRelatedEntityAppointmentType - Type of appointment.</br>Below are possible values: </br><ul><li>'Standard'- Regular.</li><li>'Expedited' - For priority appointments.</li><li>'Supplier Failed' - For rebooking appointment in case of missed appointments.</li></ul>
type PostAppointmentRequestRelatedEntityAppointmentType string

const (
	PostAppointmentRequestRelatedEntityAppointmentTypeStandard       PostAppointmentRequestRelatedEntityAppointmentType = "Standard"
	PostAppointmentRequestRelatedEntityAppointmentTypeExpedited      PostAppointmentRequestRelatedEntityAppointmentType = "Expedited"
	PostAppointmentRequestRelatedEntityAppointmentTypeSupplierFailed PostAppointmentRequestRelatedEntityAppointmentType = "Supplier Failed"
)

func (e PostAppointmentRequestRelatedEntityAppointmentType) ToPointer() *PostAppointmentRequestRelatedEntityAppointmentType {
	return &e
}

func (e *PostAppointmentRequestRelatedEntityAppointmentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Standard":
		fallthrough
	case "Expedited":
		fallthrough
	case "Supplier Failed":
		*e = PostAppointmentRequestRelatedEntityAppointmentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostAppointmentRequestRelatedEntityAppointmentType: %v", v)
	}
}

// PostAppointmentRequestRelatedEntityProductCharacteristicName - Product characteristic <i>name</i>. </b>Note: This field is passed when <i>characteristic</i> list is passed in request.
type PostAppointmentRequestRelatedEntityProductCharacteristicName string

const (
	PostAppointmentRequestRelatedEntityProductCharacteristicNameDirectoryNumber  PostAppointmentRequestRelatedEntityProductCharacteristicName = "DirectoryNumber"
	PostAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology PostAppointmentRequestRelatedEntityProductCharacteristicName = "AccessTechnology"
	PostAppointmentRequestRelatedEntityProductCharacteristicNameAccessLineID     PostAppointmentRequestRelatedEntityProductCharacteristicName = "AccessLineId"
)

func (e PostAppointmentRequestRelatedEntityProductCharacteristicName) ToPointer() *PostAppointmentRequestRelatedEntityProductCharacteristicName {
	return &e
}

func (e *PostAppointmentRequestRelatedEntityProductCharacteristicName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DirectoryNumber":
		fallthrough
	case "AccessTechnology":
		fallthrough
	case "AccessLineId":
		*e = PostAppointmentRequestRelatedEntityProductCharacteristicName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostAppointmentRequestRelatedEntityProductCharacteristicName: %v", v)
	}
}

type PostAppointmentRequestRelatedEntityProductCharacteristic struct {
	// Product characteristic <i>name</i>. </b>Note: This field is passed when <i>characteristic</i> list is passed in request.
	//
	Name *PostAppointmentRequestRelatedEntityProductCharacteristicName `json:"name,omitempty"`
	// Product characteristic <i>value</i>.</br> Note: This field is passed when <i>characteristic</i> list is passed in request.</br>
	// Below are possible values when 'AccessTechnology' is passed in <i>characteristic.name</i>,
	//  <ul><li>'Generic Ethernet Access'.</li><li>      'Generic Ethernet Access' – FTTP.</li><li>
	//  'SOGEA Existing Line'.</li><li>                       'SOGEA New Line'.</li><li>SOGEA GFAST New Line</li><li>SOGEA GFAST Existing Line</li><li>FTTC</li><li>FTTC GFAST</li><li>FTTC Sim2</li><li>FTTC GFAST Sim2</li><li>FTTP</li></ul>
	//
	Value *string `json:"value,omitempty"`
}

func (o *PostAppointmentRequestRelatedEntityProductCharacteristic) GetName() *PostAppointmentRequestRelatedEntityProductCharacteristicName {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PostAppointmentRequestRelatedEntityProductCharacteristic) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type PostAppointmentRequestRelatedEntityProductPlace struct {
	// Indicates respective place.</br>Note: This field is passed when <i>place</i> list is passed in request.
	//
	AtReferredType *string `json:"@referredType,omitempty"`
	// Unique identifier of the customer place.</br>Note: This field is passed when <i>place</i> list is passed in request. ID value represents address key that we received from the 'Common Geographic Address Management API -GET /geographicAddress API'.
	//
	ID *string `json:"id,omitempty"`
	// Role of the customer place.</br>Note: This field is passed when <i>place</i> list is passed in request.
	//
	Role *string `json:"role,omitempty"`
}

func (o *PostAppointmentRequestRelatedEntityProductPlace) GetAtReferredType() *string {
	if o == nil {
		return nil
	}
	return o.AtReferredType
}

func (o *PostAppointmentRequestRelatedEntityProductPlace) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PostAppointmentRequestRelatedEntityProductPlace) GetRole() *string {
	if o == nil {
		return nil
	}
	return o.Role
}

// PostAppointmentRequestRelatedEntityProductProductSpecification - Place holder for capturing product details against which appointment is being booked.
type PostAppointmentRequestRelatedEntityProductProductSpecification struct {
	// Product name for which appointment is being requested.</br>Etherway Superfast GEA for ethernet request and BB1 for broadband hubco user.
	ID string `json:"id"`
}

func (o *PostAppointmentRequestRelatedEntityProductProductSpecification) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// PostAppointmentRequestRelatedEntityProduct - Indicates <i>product</i> details.
type PostAppointmentRequestRelatedEntityProduct struct {
	// Indicates list of product characteristics. Characteristics of the product to instantiate or to modify.</br>Note: This list is passed when directory number search is made.
	//
	Characteristic []PostAppointmentRequestRelatedEntityProductCharacteristic `json:"characteristic,omitempty"`
	// List of places.</br>Note: This list is passed when search is made based on address key.
	//
	Place []PostAppointmentRequestRelatedEntityProductPlace `json:"place,omitempty"`
	// Place holder for capturing product details against which appointment is being booked.
	ProductSpecification PostAppointmentRequestRelatedEntityProductProductSpecification `json:"productSpecification"`
}

func (o *PostAppointmentRequestRelatedEntityProduct) GetCharacteristic() []PostAppointmentRequestRelatedEntityProductCharacteristic {
	if o == nil {
		return nil
	}
	return o.Characteristic
}

func (o *PostAppointmentRequestRelatedEntityProduct) GetPlace() []PostAppointmentRequestRelatedEntityProductPlace {
	if o == nil {
		return nil
	}
	return o.Place
}

func (o *PostAppointmentRequestRelatedEntityProduct) GetProductSpecification() PostAppointmentRequestRelatedEntityProductProductSpecification {
	if o == nil {
		return PostAppointmentRequestRelatedEntityProductProductSpecification{}
	}
	return o.ProductSpecification
}

// PostAppointmentRequestRelatedEntitySiteVisitReason - Indicates site visit reason for appointment.
// Note: Below are the default values,
// <ul><li>FTTC - 'Standard'.</li>
// <li>FTTP - 'Standard'. </li>
// <li>SOGFast - 'Premium'.</li>
// <li>SOGEA- 'Standard'.</li></ul>
type PostAppointmentRequestRelatedEntitySiteVisitReason string

const (
	PostAppointmentRequestRelatedEntitySiteVisitReasonStandard PostAppointmentRequestRelatedEntitySiteVisitReason = "Standard"
	PostAppointmentRequestRelatedEntitySiteVisitReasonPremium  PostAppointmentRequestRelatedEntitySiteVisitReason = "Premium"
)

func (e PostAppointmentRequestRelatedEntitySiteVisitReason) ToPointer() *PostAppointmentRequestRelatedEntitySiteVisitReason {
	return &e
}

func (e *PostAppointmentRequestRelatedEntitySiteVisitReason) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Standard":
		fallthrough
	case "Premium":
		*e = PostAppointmentRequestRelatedEntitySiteVisitReason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostAppointmentRequestRelatedEntitySiteVisitReason: %v", v)
	}
}

type PostAppointmentRequestRelatedEntity struct {
	// Base type of the respective entity.
	AtBaseType string `json:"@baseType"`
	// Referred type of the respective entity.
	AtReferredType string `json:"@referredType"`
	// Type of the respective entity.
	AtType string `json:"@type"`
	// Type of appointment.</br>Below are possible values: </br><ul><li>'Standard'- Regular.</li><li>'Expedited' - For priority appointments.</li><li>'Supplier Failed' - For rebooking appointment in case of missed appointments.</li></ul>
	//
	AppointmentType *PostAppointmentRequestRelatedEntityAppointmentType `json:"appointmentType,omitempty"`
	// Identifier of the respective entity.
	ID string `json:"id"`
	// It is mandatory for 'SOGEA' request. <br/>Note: When request is for new exchange line value should be 'true' or else 'false'.
	//
	NewExchangeLine *bool `json:"newExchangeLine,omitempty"`
	// Indicates <i>product</i> details.
	Product PostAppointmentRequestRelatedEntityProduct `json:"product"`
	// Role of the respective entity.
	Role string `json:"role"`
	// The value can be 'true or false', by default this value will be false if not received in input parameter list. Below are the possible value:<br/>
	// 'false'- Standalone BB Journey.<br/>
	// 'true' -SIM2 journey.<br/>
	// Note: <i>simProvide</i> field is not applicable for 'SOGEA request and Ethernet' request.
	//
	SimProvide *bool `json:"simProvide,omitempty"`
	// Indicates site visit reason for appointment.
	// Note: Below are the default values,
	// <ul><li>FTTC - 'Standard'.</li>
	// <li>FTTP - 'Standard'. </li>
	// <li>SOGFast - 'Premium'.</li>
	// <li>SOGEA- 'Standard'.</li></ul>
	//
	SiteVisitReason *PostAppointmentRequestRelatedEntitySiteVisitReason `json:"siteVisitReason,omitempty"`
}

func (o *PostAppointmentRequestRelatedEntity) GetAtBaseType() string {
	if o == nil {
		return ""
	}
	return o.AtBaseType
}

func (o *PostAppointmentRequestRelatedEntity) GetAtReferredType() string {
	if o == nil {
		return ""
	}
	return o.AtReferredType
}

func (o *PostAppointmentRequestRelatedEntity) GetAtType() string {
	if o == nil {
		return ""
	}
	return o.AtType
}

func (o *PostAppointmentRequestRelatedEntity) GetAppointmentType() *PostAppointmentRequestRelatedEntityAppointmentType {
	if o == nil {
		return nil
	}
	return o.AppointmentType
}

func (o *PostAppointmentRequestRelatedEntity) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PostAppointmentRequestRelatedEntity) GetNewExchangeLine() *bool {
	if o == nil {
		return nil
	}
	return o.NewExchangeLine
}

func (o *PostAppointmentRequestRelatedEntity) GetProduct() PostAppointmentRequestRelatedEntityProduct {
	if o == nil {
		return PostAppointmentRequestRelatedEntityProduct{}
	}
	return o.Product
}

func (o *PostAppointmentRequestRelatedEntity) GetRole() string {
	if o == nil {
		return ""
	}
	return o.Role
}

func (o *PostAppointmentRequestRelatedEntity) GetSimProvide() *bool {
	if o == nil {
		return nil
	}
	return o.SimProvide
}

func (o *PostAppointmentRequestRelatedEntity) GetSiteVisitReason() *PostAppointmentRequestRelatedEntitySiteVisitReason {
	if o == nil {
		return nil
	}
	return o.SiteVisitReason
}

// PostAppointmentRequestValidFor - List of requested time slots.
type PostAppointmentRequestValidFor struct {
	// End time for the given request.
	EndDateTime string `json:"endDateTime"`
	// Start time for the given request.
	StartDateTime string `json:"startDateTime"`
}

func (o *PostAppointmentRequestValidFor) GetEndDateTime() string {
	if o == nil {
		return ""
	}
	return o.EndDateTime
}

func (o *PostAppointmentRequestValidFor) GetStartDateTime() string {
	if o == nil {
		return ""
	}
	return o.StartDateTime
}

type PostAppointmentRequest struct {
	// List of related entities.
	RelatedEntity []PostAppointmentRequestRelatedEntity `json:"relatedEntity"`
	// List of requested time slots.
	ValidFor PostAppointmentRequestValidFor `json:"validFor"`
}

func (o *PostAppointmentRequest) GetRelatedEntity() []PostAppointmentRequestRelatedEntity {
	if o == nil {
		return []PostAppointmentRequestRelatedEntity{}
	}
	return o.RelatedEntity
}

func (o *PostAppointmentRequest) GetValidFor() PostAppointmentRequestValidFor {
	if o == nil {
		return PostAppointmentRequestValidFor{}
	}
	return o.ValidFor
}
