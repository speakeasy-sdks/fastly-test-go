# HTTPResponseFormat

Payload format for delivering to subscribers of whole HTTP responses (`response` hold mode). One of `body` or `body-bin` must be specified.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Body`                                                               | **string*                                                            | :heavy_minus_sign:                                                   | The response body as a string.                                       |
| `BodyBin`                                                            | **string*                                                            | :heavy_minus_sign:                                                   | The response body as a base64-encoded binary blob.                   |
| `Code`                                                               | **int64*                                                             | :heavy_minus_sign:                                                   | The HTTP status code.                                                |
| `Headers`                                                            | map[string]*string*                                                  | :heavy_minus_sign:                                                   | A map of arbitrary HTTP response headers to include in the response. |
| `Reason`                                                             | **string*                                                            | :heavy_minus_sign:                                                   | The HTTP status string. Defaults to a string appropriate for `code`. |