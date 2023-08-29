# BroadbandOne SDK

## Overview

Appointment Management: List available engineering appointments and reserve an engineering visit.

### Available Operations

* [CreateAppointment](#createappointment) - Create appointment
* [PostSearchAppointment](#postsearchappointment) - Appointment availability

## CreateAppointment

<p>Reserves an engineering appointment and returns an appointment ID for use in subsequent API calls.</p>      	 <p>This API supports TLS mutual authentication.</p>
 <details>
 <summary><b>Sandbox responses</b></summary>

  <p>The sandbox supports a number of success and failure scenarios. Use the parameters below in the sandbox API call to generate the corresponding responses.</p>
  <p>Success scenarios responses are simulated based on the parameters mentioned in the table.</p>

  <p>Back-end error scenarios responses are simulated based on APIGW-Tracking-Header.</p>

  <details>
  <summary><b>Success scenarios</b></summary> 

  | Parameters| Sandbox response invoked | Description |
  | ----------- | ------------------- | ------------------ |
  | <i>name</i>='DirectoryNumber', <i>value</i>='2277264355'. | 201 | Retrieve product appointments based on 'DirectoryNumber'. |
  | <i>id</i>='R0123457890'. | 201 | Retrieve get product appointments based on product place ID. |
  

  </details>

  <details>
  <summary><b>Error scenarios</b></summary>    

  | APIGW-Tracking-Header | Sandbox response invoked | Description |
  | ----------- | ------------ | ------------ | 
  | E40001 | 400, TimeSlotNotSupported | The product does not support appointments in that slot. |
  | E40301 | 403, ExchangeLineAssignedToAnotherCP | The exchange line is in use　by another communication provider. |

  
  </details>
  </details>

### Example Usage

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
    res, err := s.BroadbandOne.CreateAppointment(ctx, operations.CreateAppointmentRequest{
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
        ProductFamily: shared.ProductFamilyEthernet,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PostAppointmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateAppointmentRequest](../../models/operations/createappointmentrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.CreateAppointmentSecurity](../../models/operations/createappointmentsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.CreateAppointmentResponse](../../models/operations/createappointmentresponse.md), error**


## PostSearchAppointment

<p>Find available engineer appointment slots based on a location, product, type of work and 1st available date.</p>
	
<p>This API supports TLS mutual authentication.</p>

 <details>
 <summary><b>Sandbox responses</b></summary>

 <p>The sandbox supports a number of success and failure scenarios. Use the parameters below in the sandbox API call to generate the corresponding responses.</p>
 <p>Success scenarios responses are simulated based on the parameters mentioned in the table.</p>

 <p>Back-end error scenarios responses are simulated based on APIGW-Tracking-Header.</p>

 <details>
 <summary><b>Success scenarios</b></summary> 

 | Parameters| Sandbox response invoked | Description |
 | ----------- | ------------------- | ------------------ |
 | <i>name</i>='DirectoryNumber', <i>value</i>='2277264355'.| 201 | Retrieve product appointments based on 'DirectoryNumber'. |
 | <i>id</i>='R0123457890'. | 201 | Retrieve product appointments based on product place ID. |


 </details>

 <details>
 <summary><b>Error scenarios</b></summary>    

 | APIGW-Tracking-Header | Sandbox response invoked | Description |
 | ----------- | ------------ | ------------ | 
 | E40001 | 400, TimeSlotNotSupported | The product does not support appointments in that slot. |
 | E40301 | 403, ExchangeLineAssignedToAnotherCP | The exchange line is in use by another communication provider. |


 </details>
 </details>

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := broadbandone.New(
        broadbandone.WithSecurity(shared.Security{
            OAuth2: "",
        }),
    )

    ctx := context.Background()
    res, err := s.BroadbandOne.PostSearchAppointment(ctx, operations.PostSearchAppointmentRequest{
        APIGWTrackingHeader: "96bb97fa-b941-46bb-8c4e-86c616c28a13",
        PostSearchAppointmentRequest: &shared.PostSearchAppointmentRequest{
            RelatedEntity: []shared.PostSearchAppointmentRequestRelatedEntity{
                shared.PostSearchAppointmentRequestRelatedEntity{
                    AtBaseType: "RelatedEntity",
                    AtReferredType: "BTProductAppointmentSpecification",
                    AtType: "BTProductAppointmentSpecification",
                    AppointmentType: shared.PostSearchAppointmentRequestRelatedEntityAppointmentTypeStandard.ToPointer(),
                    ID: "1234",
                    NewExchangeLine: broadbandone.Bool(false),
                    Product: shared.PostSearchAppointmentRequestRelatedEntityProduct{
                        Characteristic: []shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristic{
                            shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                        },
                        Place: []shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                            shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                            shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                            shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                            shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                        },
                        ProductSpecification: shared.PostSearchAppointmentRequestRelatedEntityProductProductSpecification{
                            ID: "Etherway Superfast GEA",
                        },
                    },
                    Role: "OrderInformation",
                    SimProvide: broadbandone.Bool(false),
                    SiteVisitReason: shared.PostSearchAppointmentRequestRelatedEntitySiteVisitReasonStandard.ToPointer(),
                },
                shared.PostSearchAppointmentRequestRelatedEntity{
                    AtBaseType: "RelatedEntity",
                    AtReferredType: "BTProductAppointmentSpecification",
                    AtType: "BTProductAppointmentSpecification",
                    AppointmentType: shared.PostSearchAppointmentRequestRelatedEntityAppointmentTypeStandard.ToPointer(),
                    ID: "1234",
                    NewExchangeLine: broadbandone.Bool(false),
                    Product: shared.PostSearchAppointmentRequestRelatedEntityProduct{
                        Characteristic: []shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristic{
                            shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                            shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                        },
                        Place: []shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                            shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                            shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                        },
                        ProductSpecification: shared.PostSearchAppointmentRequestRelatedEntityProductProductSpecification{
                            ID: "Etherway Superfast GEA",
                        },
                    },
                    Role: "OrderInformation",
                    SimProvide: broadbandone.Bool(false),
                    SiteVisitReason: shared.PostSearchAppointmentRequestRelatedEntitySiteVisitReasonStandard.ToPointer(),
                },
                shared.PostSearchAppointmentRequestRelatedEntity{
                    AtBaseType: "RelatedEntity",
                    AtReferredType: "BTProductAppointmentSpecification",
                    AtType: "BTProductAppointmentSpecification",
                    AppointmentType: shared.PostSearchAppointmentRequestRelatedEntityAppointmentTypeStandard.ToPointer(),
                    ID: "1234",
                    NewExchangeLine: broadbandone.Bool(false),
                    Product: shared.PostSearchAppointmentRequestRelatedEntityProduct{
                        Characteristic: []shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristic{
                            shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                            shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                        },
                        Place: []shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                            shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                            shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                            shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                            shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                        },
                        ProductSpecification: shared.PostSearchAppointmentRequestRelatedEntityProductProductSpecification{
                            ID: "Etherway Superfast GEA",
                        },
                    },
                    Role: "OrderInformation",
                    SimProvide: broadbandone.Bool(false),
                    SiteVisitReason: shared.PostSearchAppointmentRequestRelatedEntitySiteVisitReasonStandard.ToPointer(),
                },
                shared.PostSearchAppointmentRequestRelatedEntity{
                    AtBaseType: "RelatedEntity",
                    AtReferredType: "BTProductAppointmentSpecification",
                    AtType: "BTProductAppointmentSpecification",
                    AppointmentType: shared.PostSearchAppointmentRequestRelatedEntityAppointmentTypeStandard.ToPointer(),
                    ID: "1234",
                    NewExchangeLine: broadbandone.Bool(false),
                    Product: shared.PostSearchAppointmentRequestRelatedEntityProduct{
                        Characteristic: []shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristic{
                            shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                            shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                            shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                            shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristic{
                                Name: shared.PostSearchAppointmentRequestRelatedEntityProductCharacteristicNameAccessTechnology.ToPointer(),
                                Value: broadbandone.String("Generic Ethernet Access"),
                            },
                        },
                        Place: []shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                            shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                            shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                            shared.PostSearchAppointmentRequestRelatedEntityProductPlace{
                                AtReferredType: broadbandone.String("OpenreachAddress"),
                                ID: broadbandone.String("Z17090800001"),
                                Role: broadbandone.String("InstallationAddress"),
                            },
                        },
                        ProductSpecification: shared.PostSearchAppointmentRequestRelatedEntityProductProductSpecification{
                            ID: "Etherway Superfast GEA",
                        },
                    },
                    Role: "OrderInformation",
                    SimProvide: broadbandone.Bool(false),
                    SiteVisitReason: shared.PostSearchAppointmentRequestRelatedEntitySiteVisitReasonStandard.ToPointer(),
                },
            },
            RequestedTimeSlot: []shared.PostSearchAppointmentRequestRequestedTimeSlot{
                shared.PostSearchAppointmentRequestRequestedTimeSlot{
                    ValidFor: shared.PostSearchAppointmentRequestRequestedTimeSlotValidFor{
                        StartDateTime: "2022-01-17T09:00:00.000Z",
                    },
                },
                shared.PostSearchAppointmentRequestRequestedTimeSlot{
                    ValidFor: shared.PostSearchAppointmentRequestRequestedTimeSlotValidFor{
                        StartDateTime: "2022-01-17T09:00:00.000Z",
                    },
                },
            },
        },
        ProductFamily: shared.ProductFamilyBroadbandOne,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostSearchAppointmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PostSearchAppointmentRequest](../../models/operations/postsearchappointmentrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.PostSearchAppointmentResponse](../../models/operations/postsearchappointmentresponse.md), error**

