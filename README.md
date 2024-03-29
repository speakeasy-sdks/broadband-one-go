<div align="center">
    <img src="https://github.com/speakeasy-sdks/broadband-one-go/assets/6267663/22e4e028-de1d-4942-a3dd-48abd278f1ea" width="100">
    <h1>Go SDK</h1>
   <p>Place orders for Broadband One from your systems to The Hub.</p>
   <a href="https://developer.bt.com/products/broadband-one-order-api"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=5514b4&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/broadband-one-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/broadband-one-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

![bb1-diagram-h_0 (1)](https://github.com/speakeasy-sdks/broadband-one-go/assets/6267663/bfc886eb-2ab0-4c64-8423-ef6b3e7d21f7)

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/broadband-one-go
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Appointments](docs/sdks/appointments/README.md)

* [Schedule](docs/sdks/appointments/README.md#schedule) - Create appointment
* [Search](docs/sdks/appointments/README.md#search) - Appointment availability
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
