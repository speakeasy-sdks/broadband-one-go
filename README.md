# openapi

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/broadband-one-ts
```
<!-- End SDK Installation -->

## SDK Example Usage
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
    res, err := s.CreateAppointment(ctx, operations.CreateAppointmentRequest{
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations

### [BroadbandOne SDK](docs/sdks/broadbandone/README.md)

* [CreateAppointment](docs/sdks/broadbandone/README.md#createappointment) - Create appointment
* [PostSearchAppointment](docs/sdks/broadbandone/README.md#postsearchappointment) - Appointment availability
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
