# GetInvoiceByIDResponse


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ContentType`                                                     | *string*                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `StatusCode`                                                      | *int*                                                             | :heavy_check_mark:                                                | N/A                                                               |
| `RawResponse`                                                     | [*http.Response](https://pkg.go.dev/net/http#Response)            | :heavy_minus_sign:                                                | N/A                                                               |
| `BillingResponse`                                                 | [*shared.BillingResponse](../../models/shared/billingresponse.md) | :heavy_minus_sign:                                                | OK                                                                |
| `GetInvoiceByID200ApplicationPdfBinaryString`                     | *[]byte*                                                          | :heavy_minus_sign:                                                | OK                                                                |
| `GetInvoiceByID200TextCsvCsvString`                               | **string*                                                         | :heavy_minus_sign:                                                | OK                                                                |