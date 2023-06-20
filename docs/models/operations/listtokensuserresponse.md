# ListTokensUserResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ContentType`                                                         | *string*                                                              | :heavy_check_mark:                                                    | N/A                                                                   |
| `StatusCode`                                                          | *int*                                                                 | :heavy_check_mark:                                                    | N/A                                                                   |
| `RawResponse`                                                         | [*http.Response](https://pkg.go.dev/net/http#Response)                | :heavy_minus_sign:                                                    | N/A                                                                   |
| `GenericTokenError`                                                   | [*shared.GenericTokenError](../../models/shared/generictokenerror.md) | :heavy_minus_sign:                                                    | Missing or expired token.                                             |
| `TokenResponses`                                                      | [][shared.TokenResponse](../../models/shared/tokenresponse.md)        | :heavy_minus_sign:                                                    | OK                                                                    |