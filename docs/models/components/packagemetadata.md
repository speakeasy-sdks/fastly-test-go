# PackageMetadata

[Package metadata](#metadata-model) that has been extracted from the uploaded package.



## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Authors`                                          | []*string*                                         | :heavy_minus_sign:                                 | A list of package authors' email addresses.        |                                                    |
| `Description`                                      | **string*                                          | :heavy_minus_sign:                                 | Description of the Compute@Edge package.           |                                                    |
| `FilesHash`                                        | **string*                                          | :heavy_minus_sign:                                 | Hash of the files within the Compute@Edge package. |                                                    |
| `Hashsum`                                          | **string*                                          | :heavy_minus_sign:                                 | Hash of the Compute@Edge package.                  |                                                    |
| `Language`                                         | **string*                                          | :heavy_minus_sign:                                 | The language of the Compute@Edge package.          | rust                                               |
| `Name`                                             | **string*                                          | :heavy_minus_sign:                                 | Name of the Compute@Edge package.                  |                                                    |
| `Size`                                             | **int64*                                           | :heavy_minus_sign:                                 | Size of the Compute@Edge package in bytes.         |                                                    |