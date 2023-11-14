# TLSBulkCertificateDataAttributes


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `AllowUntrustedRoot`                                     | **bool*                                                  | :heavy_minus_sign:                                       | Allow certificates that chain to untrusted roots.        |
| `CertBlob`                                               | **string*                                                | :heavy_minus_sign:                                       | The PEM-formatted certificate blob. Required.            |
| `IntermediatesBlob`                                      | **string*                                                | :heavy_minus_sign:                                       | The PEM-formatted chain of intermediate blobs. Required. |