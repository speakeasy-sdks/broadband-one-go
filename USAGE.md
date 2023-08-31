<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/operations"
	"openapi/pkg/models/shared"
)

func main() {
    s := broadbandone.New()
    operationSecurity := operations.CreateAppointmentSecurity{
            OAuth2: "",
        }

    ctx := context.Background()
    res, err := s.Appointments.Schedule(ctx, operations.CreateAppointmentRequest{
        APIGWTrackingHeader: "96bb97fa-b941-46bb-8c4e-86c616c28a13",
        PostAppointmentRequest: &shared.PostAppointmentRequest{
            RelatedEntity: []shared.PostAppointmentRequestRelatedEntity{
                shared.PostAppointmentRequestRelatedEntity{
                    AtBaseType: "RelatedEntity",
                    AtReferredType: "BTProductAppointmentSpecification",
                    AtType: "BTProductAppointmentSpecification",
                    AppointmentType: shared.PostAppointmentRequestRelatedEntityAppointmentTypeStandard.ToPointer(),
                    ID: "1234",
                    NewExchangeLine: broadbandone.Bool(false),
                    Product: shared.PostAppointmentRequestRelatedEntityProduct{
                        Characteristic: []shared.PostAppointmentRequestRelatedEntityProductCharacteristic{
                            shared.PostAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                            shared.PostAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                            shared.PostAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                        },
                        Place: []shared.PostAppointmentRequestRelatedEntityProductPlace{
                            shared.PostAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                            shared.PostAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                            shared.PostAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                        },
                        ProductSpecification: shared.PostAppointmentRequestRelatedEntityProductProductSpecification{
                            ID: "Etherway Superfast GEA",
                        },
                    },
                    Role: "OrderInformation",
                    SimProvide: broadbandone.Bool(true),
                    SiteVisitReason: shared.PostAppointmentRequestRelatedEntitySiteVisitReasonStandard.ToPointer(),
                },
                shared.PostAppointmentRequestRelatedEntity{
                    AtBaseType: "RelatedEntity",
                    AtReferredType: "BTProductAppointmentSpecification",
                    AtType: "BTProductAppointmentSpecification",
                    AppointmentType: shared.PostAppointmentRequestRelatedEntityAppointmentTypeStandard.ToPointer(),
                    ID: "1234",
                    NewExchangeLine: broadbandone.Bool(false),
                    Product: shared.PostAppointmentRequestRelatedEntityProduct{
                        Characteristic: []shared.PostAppointmentRequestRelatedEntityProductCharacteristic{
                            shared.PostAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                            shared.PostAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                            shared.PostAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                            shared.PostAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                        },
                        Place: []shared.PostAppointmentRequestRelatedEntityProductPlace{
                            shared.PostAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                            shared.PostAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                            shared.PostAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                        },
                        ProductSpecification: shared.PostAppointmentRequestRelatedEntityProductProductSpecification{
                            ID: "Etherway Superfast GEA",
                        },
                    },
                    Role: "OrderInformation",
                    SimProvide: broadbandone.Bool(true),
                    SiteVisitReason: shared.PostAppointmentRequestRelatedEntitySiteVisitReasonStandard.ToPointer(),
                },
                shared.PostAppointmentRequestRelatedEntity{
                    AtBaseType: "RelatedEntity",
                    AtReferredType: "BTProductAppointmentSpecification",
                    AtType: "BTProductAppointmentSpecification",
                    AppointmentType: shared.PostAppointmentRequestRelatedEntityAppointmentTypeStandard.ToPointer(),
                    ID: "1234",
                    NewExchangeLine: broadbandone.Bool(false),
                    Product: shared.PostAppointmentRequestRelatedEntityProduct{
                        Characteristic: []shared.PostAppointmentRequestRelatedEntityProductCharacteristic{
                            shared.PostAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                            shared.PostAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                            shared.PostAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                            shared.PostAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                        },
                        Place: []shared.PostAppointmentRequestRelatedEntityProductPlace{
                            shared.PostAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                            shared.PostAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                            shared.PostAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                        },
                        ProductSpecification: shared.PostAppointmentRequestRelatedEntityProductProductSpecification{
                            ID: "Etherway Superfast GEA",
                        },
                    },
                    Role: "OrderInformation",
                    SimProvide: broadbandone.Bool(true),
                    SiteVisitReason: shared.PostAppointmentRequestRelatedEntitySiteVisitReasonStandard.ToPointer(),
                },
            },
            ValidFor: shared.PostAppointmentRequestValidFor{
                EndDateTime: "2022-03-17T11:00:00Z",
                StartDateTime: "2022-03-17T09:00:00Z",
            },
        },
        ProductFamily: shared.ProductFamilyBb1Hub,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PostAppointmentResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->