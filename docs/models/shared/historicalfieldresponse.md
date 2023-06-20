# HistoricalFieldResponse

OK


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Data`                                                                             | map[string][]map[string]*string*                                                   | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `Meta`                                                                             | [*HistoricalFieldResponseMeta](../../models/shared/historicalfieldresponsemeta.md) | :heavy_minus_sign:                                                                 | Meta information about the scope of the query in a human readable format.          |
| `Msg`                                                                              | **string*                                                                          | :heavy_minus_sign:                                                                 | If the query was not successful, this will provide a string that explains why.     |
| `Status`                                                                           | **string*                                                                          | :heavy_minus_sign:                                                                 | Whether or not we were able to successfully execute the query.                     |