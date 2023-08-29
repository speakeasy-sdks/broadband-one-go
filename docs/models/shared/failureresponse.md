# FailureResponse

Unauthorized

| Code | Message | Description |
| ---- | ------- | ----------- |
| 40 | Missing credentials | The <i>Authorization</i> parameter is missing. |
| 41 | Invalid credentials | The <i>Authorization</i> parameter is not valid. |
| 42 | Expired credentials | Renew the access token using the OAuth API and try again. |



## Fields

| Field                         | Type                          | Required                      | Description                   |
| ----------------------------- | ----------------------------- | ----------------------------- | ----------------------------- |
| `Code`                        | *string*                      | :heavy_check_mark:            | Represents the error code.    |
| `Message`                     | *string*                      | :heavy_check_mark:            | Represents the error message. |