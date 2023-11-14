# Pop


## Fields

| Field                                                                                                                                                     | Type                                                                                                                                                      | Required                                                                                                                                                  | Description                                                                                                                                               |
| --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `BillingRegion`                                                                                                                                           | [components.BillingRegion](../../models/components/billingregion.md)                                                                                      | :heavy_check_mark:                                                                                                                                        | the region used for billing                                                                                                                               |
| `Code`                                                                                                                                                    | *string*                                                                                                                                                  | :heavy_check_mark:                                                                                                                                        | the three-letter code for the [POP](https://developer.fastly.com/learning/concepts/pop/)                                                                  |
| `Coordinates`                                                                                                                                             | [*components.Coordinates](../../models/components/coordinates.md)                                                                                         | :heavy_minus_sign:                                                                                                                                        | the geographic location of the POP                                                                                                                        |
| `Group`                                                                                                                                                   | *string*                                                                                                                                                  | :heavy_check_mark:                                                                                                                                        | N/A                                                                                                                                                       |
| `Name`                                                                                                                                                    | *string*                                                                                                                                                  | :heavy_check_mark:                                                                                                                                        | the name of the POP                                                                                                                                       |
| `Region`                                                                                                                                                  | [components.PopRegion](../../models/components/popregion.md)                                                                                              | :heavy_check_mark:                                                                                                                                        | N/A                                                                                                                                                       |
| `Shield`                                                                                                                                                  | **string*                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                        | the name of the [shield code](https://developer.fastly.com/learning/concepts/shielding/#choosing-a-shield-location) if this POP is suitable for shielding |
| `StatsRegion`                                                                                                                                             | [components.StatsRegion](../../models/components/statsregion.md)                                                                                          | :heavy_check_mark:                                                                                                                                        | the region used for stats reporting                                                                                                                       |