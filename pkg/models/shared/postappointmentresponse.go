// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PostAppointmentResponseRelatedEntityAppointmentType - Type of an appointment. </br>Below are possible values:</br><ul><li>'Standard' - Regular.</li><li>'Expedited' - For priority appointments.</li><li>'Supplier Failed' - For rebooking appointment in case of missed appointments.</li></ul>
type PostAppointmentResponseRelatedEntityAppointmentType string

const (
	PostAppointmentResponseRelatedEntityAppointmentTypeStandard       PostAppointmentResponseRelatedEntityAppointmentType = "Standard"
	PostAppointmentResponseRelatedEntityAppointmentTypeExpedited      PostAppointmentResponseRelatedEntityAppointmentType = "Expedited"
	PostAppointmentResponseRelatedEntityAppointmentTypeSupplierFailed PostAppointmentResponseRelatedEntityAppointmentType = "Supplier Failed"
)

func (e PostAppointmentResponseRelatedEntityAppointmentType) ToPointer() *PostAppointmentResponseRelatedEntityAppointmentType {
	return &e
}

func (e *PostAppointmentResponseRelatedEntityAppointmentType) UnmarshalJSON(data []byte) error {
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
		*e = PostAppointmentResponseRelatedEntityAppointmentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostAppointmentResponseRelatedEntityAppointmentType: %v", v)
	}
}

// PostAppointmentResponseRelatedEntityProductCharacteristicName - Product characteristic <i>name</i>. </b>Note: This field is passed when <i>characteristic</i> list is passed in request.
type PostAppointmentResponseRelatedEntityProductCharacteristicName string

const (
	PostAppointmentResponseRelatedEntityProductCharacteristicNameDirectoryNumber  PostAppointmentResponseRelatedEntityProductCharacteristicName = "DirectoryNumber"
	PostAppointmentResponseRelatedEntityProductCharacteristicNameAccessTechnology PostAppointmentResponseRelatedEntityProductCharacteristicName = "AccessTechnology"
	PostAppointmentResponseRelatedEntityProductCharacteristicNameAccessLineID     PostAppointmentResponseRelatedEntityProductCharacteristicName = "AccessLineId"
)

func (e PostAppointmentResponseRelatedEntityProductCharacteristicName) ToPointer() *PostAppointmentResponseRelatedEntityProductCharacteristicName {
	return &e
}

func (e *PostAppointmentResponseRelatedEntityProductCharacteristicName) UnmarshalJSON(data []byte) error {
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
		*e = PostAppointmentResponseRelatedEntityProductCharacteristicName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostAppointmentResponseRelatedEntityProductCharacteristicName: %v", v)
	}
}

type PostAppointmentResponseRelatedEntityProductCharacteristic struct {
	// Product characteristic <i>name</i>. </b>Note: This field is passed when <i>characteristic</i> list is passed in request.
	//
	Name *PostAppointmentResponseRelatedEntityProductCharacteristicName `json:"name,omitempty"`
	// Product characteristic <i>value</i>.</br> Note: This field is passed when <i>characteristic</i> list is passed in request.</br>
	//  Below are possible values when 'AccessTechnology' is passed in <i>characteristic.name</i>,
	//  <ul><li>'Generic Ethernet Access'.</li><li>'Generic Ethernet Access' – FTTP.</li><li>
	//  'SOGEA Existing Line'.</li><li>'SOGEA New Line'</li><li>SOGEA GFAST New Line</li><li>SOGEA GFAST Existing Line</li><li>FTTC</li><li>FTTC GFAST</li><li>FTTC Sim2</li><li>FTTC GFAST Sim2</li><li>FTTP</li></ul>
	//
	Value *string `json:"value,omitempty"`
}

func (o *PostAppointmentResponseRelatedEntityProductCharacteristic) GetName() *PostAppointmentResponseRelatedEntityProductCharacteristicName {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PostAppointmentResponseRelatedEntityProductCharacteristic) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type PostAppointmentResponseRelatedEntityProductPlace struct {
	// Indicates the referred type of the respective place.</br>Note: This field is passed when <i>place</i> list is passed in request.
	//
	AtReferredType *string `json:"@referredType,omitempty"`
	// Unique identifier of the respective place.</br>Note: This field is passed when <i>place</i> list is passed in request. ID value represents address key that we received from the 'Common Geographic Address Management API -GET /geographicAddress API'.
	//
	ID *string `json:"id,omitempty"`
	// Role of the respective place.</br>Note: This field is passed when <i>place</i> list is passed in request.
	//
	Role *string `json:"role,omitempty"`
}

func (o *PostAppointmentResponseRelatedEntityProductPlace) GetAtReferredType() *string {
	if o == nil {
		return nil
	}
	return o.AtReferredType
}

func (o *PostAppointmentResponseRelatedEntityProductPlace) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PostAppointmentResponseRelatedEntityProductPlace) GetRole() *string {
	if o == nil {
		return nil
	}
	return o.Role
}

// PostAppointmentResponseRelatedEntityProductProductSpecification - Indicates detailed description of a tangible or intangible object made available externally in the form of a 'productOffering' to customers or other parties playing a party role.
type PostAppointmentResponseRelatedEntityProductProductSpecification struct {
	// Unique identifier of the product specification.</br>Note: It is the 'SCODE' of the product, there is no length validation on the 'SCODE'. The passed 'SCODE' must be valid.
	//
	ID string `json:"id"`
}

func (o *PostAppointmentResponseRelatedEntityProductProductSpecification) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// PostAppointmentResponseRelatedEntityProduct - Indicates <i>product</i> details.
type PostAppointmentResponseRelatedEntityProduct struct {
	// Indicates list of product characteristics. Characteristics of the product to instantiate or to modify.</br>Note: This list is passed when directory number search is made.
	//
	Characteristic []PostAppointmentResponseRelatedEntityProductCharacteristic `json:"characteristic,omitempty"`
	// List of places.</br>Note: This list is passed when search is made based on address key.
	//
	Place []PostAppointmentResponseRelatedEntityProductPlace `json:"place,omitempty"`
	// Indicates detailed description of a tangible or intangible object made available externally in the form of a 'productOffering' to customers or other parties playing a party role.
	ProductSpecification PostAppointmentResponseRelatedEntityProductProductSpecification `json:"productSpecification"`
}

func (o *PostAppointmentResponseRelatedEntityProduct) GetCharacteristic() []PostAppointmentResponseRelatedEntityProductCharacteristic {
	if o == nil {
		return nil
	}
	return o.Characteristic
}

func (o *PostAppointmentResponseRelatedEntityProduct) GetPlace() []PostAppointmentResponseRelatedEntityProductPlace {
	if o == nil {
		return nil
	}
	return o.Place
}

func (o *PostAppointmentResponseRelatedEntityProduct) GetProductSpecification() PostAppointmentResponseRelatedEntityProductProductSpecification {
	if o == nil {
		return PostAppointmentResponseRelatedEntityProductProductSpecification{}
	}
	return o.ProductSpecification
}

// PostAppointmentResponseRelatedEntitySiteVisitReason - Indicates site visit reason. This field required to have values either 'Standard or Premium', refer below note for default values.
// Note: Below are the default values,
// <ul><li>FTTC - 'Standard'.</li>
// <li>FTTP - 'Standard'. </li>
// <li>SOGFast - 'Premium'.</li>
// <li>SOGEA- 'Standard'.</li></ul>
type PostAppointmentResponseRelatedEntitySiteVisitReason string

const (
	PostAppointmentResponseRelatedEntitySiteVisitReasonStandard PostAppointmentResponseRelatedEntitySiteVisitReason = "Standard"
	PostAppointmentResponseRelatedEntitySiteVisitReasonPremium  PostAppointmentResponseRelatedEntitySiteVisitReason = "Premium"
)

func (e PostAppointmentResponseRelatedEntitySiteVisitReason) ToPointer() *PostAppointmentResponseRelatedEntitySiteVisitReason {
	return &e
}

func (e *PostAppointmentResponseRelatedEntitySiteVisitReason) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Standard":
		fallthrough
	case "Premium":
		*e = PostAppointmentResponseRelatedEntitySiteVisitReason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostAppointmentResponseRelatedEntitySiteVisitReason: %v", v)
	}
}

type PostAppointmentResponseRelatedEntity struct {
	// Base type of the respective entity.
	AtBaseType string `json:"@baseType"`
	// Referred type of the respective entity.
	AtReferredType string `json:"@referredType"`
	// Type of the appointment resource.
	AtType string `json:"@type"`
	// Type of an appointment. </br>Below are possible values:</br><ul><li>'Standard' - Regular.</li><li>'Expedited' - For priority appointments.</li><li>'Supplier Failed' - For rebooking appointment in case of missed appointments.</li></ul>
	//
	AppointmentType *PostAppointmentResponseRelatedEntityAppointmentType `json:"appointmentType,omitempty"`
	// Identifier of the respective entity.
	ID string `json:"id"`
	// Indicates new exchange for SOGEA.
	NewExchangeLine *bool `json:"newExchangeLine,omitempty"`
	// Indicates <i>product</i> details.
	Product PostAppointmentResponseRelatedEntityProduct `json:"product"`
	// Indicates the role of the respective entity.
	Role string `json:"role"`
	// The value can be 'true' or 'false', by default this value will be false if not received in input parameter list.
	SimProvide *bool `json:"simProvide,omitempty"`
	// Indicates site visit reason. This field required to have values either 'Standard or Premium', refer below note for default values.
	// Note: Below are the default values,
	// <ul><li>FTTC - 'Standard'.</li>
	// <li>FTTP - 'Standard'. </li>
	// <li>SOGFast - 'Premium'.</li>
	// <li>SOGEA- 'Standard'.</li></ul>
	//
	SiteVisitReason *PostAppointmentResponseRelatedEntitySiteVisitReason `json:"siteVisitReason,omitempty"`
}

func (o *PostAppointmentResponseRelatedEntity) GetAtBaseType() string {
	if o == nil {
		return ""
	}
	return o.AtBaseType
}

func (o *PostAppointmentResponseRelatedEntity) GetAtReferredType() string {
	if o == nil {
		return ""
	}
	return o.AtReferredType
}

func (o *PostAppointmentResponseRelatedEntity) GetAtType() string {
	if o == nil {
		return ""
	}
	return o.AtType
}

func (o *PostAppointmentResponseRelatedEntity) GetAppointmentType() *PostAppointmentResponseRelatedEntityAppointmentType {
	if o == nil {
		return nil
	}
	return o.AppointmentType
}

func (o *PostAppointmentResponseRelatedEntity) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PostAppointmentResponseRelatedEntity) GetNewExchangeLine() *bool {
	if o == nil {
		return nil
	}
	return o.NewExchangeLine
}

func (o *PostAppointmentResponseRelatedEntity) GetProduct() PostAppointmentResponseRelatedEntityProduct {
	if o == nil {
		return PostAppointmentResponseRelatedEntityProduct{}
	}
	return o.Product
}

func (o *PostAppointmentResponseRelatedEntity) GetRole() string {
	if o == nil {
		return ""
	}
	return o.Role
}

func (o *PostAppointmentResponseRelatedEntity) GetSimProvide() *bool {
	if o == nil {
		return nil
	}
	return o.SimProvide
}

func (o *PostAppointmentResponseRelatedEntity) GetSiteVisitReason() *PostAppointmentResponseRelatedEntitySiteVisitReason {
	if o == nil {
		return nil
	}
	return o.SiteVisitReason
}

// PostAppointmentResponseStatus - Indicates the <i>status</i> of the requested slot.
type PostAppointmentResponseStatus string

const (
	PostAppointmentResponseStatusConfirmed   PostAppointmentResponseStatus = "confirmed"
	PostAppointmentResponseStatusInProgress  PostAppointmentResponseStatus = "inProgress"
	PostAppointmentResponseStatusCancelled   PostAppointmentResponseStatus = "cancelled"
	PostAppointmentResponseStatusInitialized PostAppointmentResponseStatus = "initialized"
	PostAppointmentResponseStatusCompleted   PostAppointmentResponseStatus = "completed"
	PostAppointmentResponseStatusFailed      PostAppointmentResponseStatus = "failed"
)

func (e PostAppointmentResponseStatus) ToPointer() *PostAppointmentResponseStatus {
	return &e
}

func (e *PostAppointmentResponseStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "confirmed":
		fallthrough
	case "inProgress":
		fallthrough
	case "cancelled":
		fallthrough
	case "initialized":
		fallthrough
	case "completed":
		fallthrough
	case "failed":
		*e = PostAppointmentResponseStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostAppointmentResponseStatus: %v", v)
	}
}

// PostAppointmentResponseValidFor - List of requested time slots.
type PostAppointmentResponseValidFor struct {
	// End time for the given request. Format is TMF date format.
	EndDateTime string `json:"endDateTime"`
	// Start time for the given request. Format is TMF date format.
	StartDateTime string `json:"startDateTime"`
}

func (o *PostAppointmentResponseValidFor) GetEndDateTime() string {
	if o == nil {
		return ""
	}
	return o.EndDateTime
}

func (o *PostAppointmentResponseValidFor) GetStartDateTime() string {
	if o == nil {
		return ""
	}
	return o.StartDateTime
}

// PostAppointmentResponse - created
type PostAppointmentResponse struct {
	// Indicates the creation date for the given request. Format is TMF date.
	CreationDate *string `json:"creationDate,omitempty"`
	// Appointment ID for the confirmed appointment.</br>This ID should be used in the request for the order journey.
	ID string `json:"id"`
	// List of related entities.
	RelatedEntity []PostAppointmentResponseRelatedEntity `json:"relatedEntity"`
	// Indicates the <i>status</i> of the requested slot.
	Status PostAppointmentResponseStatus `json:"status"`
	// List of requested time slots.
	ValidFor PostAppointmentResponseValidFor `json:"validFor"`
}

func (o *PostAppointmentResponse) GetCreationDate() *string {
	if o == nil {
		return nil
	}
	return o.CreationDate
}

func (o *PostAppointmentResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PostAppointmentResponse) GetRelatedEntity() []PostAppointmentResponseRelatedEntity {
	if o == nil {
		return []PostAppointmentResponseRelatedEntity{}
	}
	return o.RelatedEntity
}

func (o *PostAppointmentResponse) GetStatus() PostAppointmentResponseStatus {
	if o == nil {
		return PostAppointmentResponseStatus("")
	}
	return o.Status
}

func (o *PostAppointmentResponse) GetValidFor() PostAppointmentResponseValidFor {
	if o == nil {
		return PostAppointmentResponseValidFor{}
	}
	return o.ValidFor
}
