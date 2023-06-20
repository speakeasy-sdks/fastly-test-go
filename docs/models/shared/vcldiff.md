# VclDiff

OK


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `Diff`                                                                                        | **string*                                                                                     | :heavy_minus_sign:                                                                            | The differences between two specified versions.                                               |
| `Format`                                                                                      | [*VclDiffFormat](../../models/shared/vcldiffformat.md)                                        | :heavy_minus_sign:                                                                            | The format in which compared VCL changes are being returned in.                               |
| `From`                                                                                        | **int64*                                                                                      | :heavy_minus_sign:                                                                            | The version number of the service to which changes in the generated VCL are being compared.   |
| `To`                                                                                          | **int64*                                                                                      | :heavy_minus_sign:                                                                            | The version number of the service from which changes in the generated VCL are being compared. |