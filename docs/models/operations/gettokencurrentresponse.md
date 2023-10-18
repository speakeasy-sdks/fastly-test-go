# GetTokenCurrentResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ContentType`                                                         | *string*                                                              | :heavy_check_mark:                                                    | HTTP response content type for this operation                         |
| `StatusCode`                                                          | *int*                                                                 | :heavy_check_mark:                                                    | HTTP response status code for this operation                          |
| `RawResponse`                                                         | [*http.Response](https://pkg.go.dev/net/http#Response)                | :heavy_minus_sign:                                                    | Raw HTTP response; suitable for custom response parsing               |
| `GenericTokenError`                                                   | [*shared.GenericTokenError](../../models/shared/generictokenerror.md) | :heavy_minus_sign:                                                    | Missing or expired token.                                             |
| `TokenResponse`                                                       | [*shared.TokenResponse](../../models/shared/tokenresponse.md)         | :heavy_minus_sign:                                                    | OK                                                                    |