# PostSearchAppointmentResponseAvailableTimeSlot


## Fields

| Field                                                                                                                                   | Type                                                                                                                                    | Required                                                                                                                                | Description                                                                                                                             | Example                                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| `Slot`                                                                                                                                  | *string*                                                                                                                                | :heavy_check_mark:                                                                                                                      | Indicates the slot for the given request.                                                                                               | 0900                                                                                                                                    |
| `ValidFor`                                                                                                                              | [PostSearchAppointmentResponseAvailableTimeSlotValidFor](../../models/shared/postsearchappointmentresponseavailabletimeslotvalidfor.md) | :heavy_check_mark:                                                                                                                      | Valid time slot for the given request. Difference between start date time and end date time provides the slot duration.                 |                                                                                                                                         |