# HistoricalRegionsResponse


## Fields

| Field                                                                                                 | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `Data`                                                                                                | []*string*                                                                                            | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `Meta`                                                                                                | [*components.HistoricalRegionsResponseMeta](../../models/components/historicalregionsresponsemeta.md) | :heavy_minus_sign:                                                                                    | Meta information about the scope of the query in a human readable format.                             |
| `Msg`                                                                                                 | **string*                                                                                             | :heavy_minus_sign:                                                                                    | If the query was not successful, this will provide a string that explains why.                        |
| `Status`                                                                                              | **string*                                                                                             | :heavy_minus_sign:                                                                                    | Whether or not we were able to successfully execute the query.                                        |