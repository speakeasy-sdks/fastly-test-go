# HistoricalResponse

OK


## Fields

| Field                                                                                                                                  | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `Data`                                                                                                                                 | map[string][][Results](../../models/shared/results.md)                                                                                 | :heavy_minus_sign:                                                                                                                     | Contains the results of the query, organized by *service ID*, into arrays where each element describes one service over a *time span*. |
| `Meta`                                                                                                                                 | [*HistoricalResponseMeta](../../models/shared/historicalresponsemeta.md)                                                               | :heavy_minus_sign:                                                                                                                     | Meta information about the scope of the query in a human readable format.                                                              |
| `Msg`                                                                                                                                  | **string*                                                                                                                              | :heavy_minus_sign:                                                                                                                     | If the query was not successful, this will provide a string that explains why.                                                         |
| `Status`                                                                                                                               | **string*                                                                                                                              | :heavy_minus_sign:                                                                                                                     | Whether or not we were able to successfully execute the query.                                                                         |