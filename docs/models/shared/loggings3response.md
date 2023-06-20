# LoggingS3Response

OK


## Fields

| Field                                                                                                                                                                                                                                                                            | Type                                                                                                                                                                                                                                                                             | Required                                                                                                                                                                                                                                                                         | Description                                                                                                                                                                                                                                                                      | Example                                                                                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `AccessKey`                                                                                                                                                                                                                                                                      | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The access key for your S3 account. Not required if `iam_role` is provided.                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                  |
| `ACL`                                                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The access control list (ACL) specific request header. See the AWS documentation for [Access Control List (ACL) Specific Request Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/mpUploadInitiate.html#initiate-mpu-acl-specific-request-headers) for more information. |                                                                                                                                                                                                                                                                                  |
| `BucketName`                                                                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The bucket name for S3 account.                                                                                                                                                                                                                                                  |                                                                                                                                                                                                                                                                                  |
| `CompressionCodec`                                                                                                                                                                                                                                                               | [*LoggingS3ResponseCompressionCodec](../../models/shared/loggings3responsecompressioncodec.md)                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.                                                                                   |                                                                                                                                                                                                                                                                                  |
| `CreatedAt`                                                                                                                                                                                                                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                                               | Date and time in ISO 8601 format.                                                                                                                                                                                                                                                | 2020-04-09T18:14:30Z                                                                                                                                                                                                                                                             |
| `DeletedAt`                                                                                                                                                                                                                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                                               | Date and time in ISO 8601 format.                                                                                                                                                                                                                                                | 2020-04-09T18:14:30Z                                                                                                                                                                                                                                                             |
| `Domain`                                                                                                                                                                                                                                                                         | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The domain of the Amazon S3 endpoint.                                                                                                                                                                                                                                            |                                                                                                                                                                                                                                                                                  |
| `Format`                                                                                                                                                                                                                                                                         | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).                                                                                                                                                                                              | %h %l %u %t "%r" %&gt;s %b                                                                                                                                                                                                                                                       |
| `FormatVersion`                                                                                                                                                                                                                                                                  | [*LoggingS3ResponseFormatVersion](../../models/shared/loggings3responseformatversion.md)                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.<br/>                                                   | 2                                                                                                                                                                                                                                                                                |
| `GzipLevel`                                                                                                                                                                                                                                                                      | **int64*                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.                                                                                                | 0                                                                                                                                                                                                                                                                                |
| `IamRole`                                                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The Amazon Resource Name (ARN) for the IAM role granting Fastly access to S3. Not required if `access_key` and `secret_key` are provided.                                                                                                                                        |                                                                                                                                                                                                                                                                                  |
| `MessageType`                                                                                                                                                                                                                                                                    | [*LoggingS3ResponseMessageType](../../models/shared/loggings3responsemessagetype.md)                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                               | How the message should be formatted.                                                                                                                                                                                                                                             | classic                                                                                                                                                                                                                                                                          |
| `Name`                                                                                                                                                                                                                                                                           | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The name for the real-time logging configuration.                                                                                                                                                                                                                                | test-log-endpoint                                                                                                                                                                                                                                                                |
| `Path`                                                                                                                                                                                                                                                                           | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The path to upload logs to.                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                  |
| `Period`                                                                                                                                                                                                                                                                         | **int64*                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                               | How frequently log files are finalized so they can be available for reading (in seconds).                                                                                                                                                                                        | 3600                                                                                                                                                                                                                                                                             |
| `Placement`                                                                                                                                                                                                                                                                      | [*LoggingS3ResponsePlacement](../../models/shared/loggings3responseplacement.md)                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                               | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.<br/>                                                                    | null                                                                                                                                                                                                                                                                             |
| `PublicKey`                                                                                                                                                                                                                                                                      | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | A PGP public key that Fastly will use to encrypt your log files before writing them to disk.                                                                                                                                                                                     | -----BEGIN PRIVATE KEY-----<br/>...<br/>-----END PRIVATE KEY-----<br/>                                                                                                                                                                                                           |
| `Redundancy`                                                                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The S3 redundancy level.                                                                                                                                                                                                                                                         |                                                                                                                                                                                                                                                                                  |
| `ResponseCondition`                                                                                                                                                                                                                                                              | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The name of an existing condition in the configured endpoint, or leave blank to always execute.                                                                                                                                                                                  | null                                                                                                                                                                                                                                                                             |
| `SecretKey`                                                                                                                                                                                                                                                                      | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The secret key for your S3 account. Not required if `iam_role` is provided.                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                  |
| `ServerSideEncryption`                                                                                                                                                                                                                                                           | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | Set this to `AES256` or `aws:kms` to enable S3 Server Side Encryption.                                                                                                                                                                                                           |                                                                                                                                                                                                                                                                                  |
| `ServerSideEncryptionKmsKeyID`                                                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | Optional server-side KMS Key Id. Must be set if `server_side_encryption` is set to `aws:kms` or `AES256`.                                                                                                                                                                        |                                                                                                                                                                                                                                                                                  |
| `ServiceID`                                                                                                                                                                                                                                                                      | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | N/A                                                                                                                                                                                                                                                                              | SU1Z0isxPaozGVKXdv0eY                                                                                                                                                                                                                                                            |
| `TimestampFormat`                                                                                                                                                                                                                                                                | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | A timestamp format                                                                                                                                                                                                                                                               | %Y-%m-%dT%H:%M:%S.000                                                                                                                                                                                                                                                            |
| `UpdatedAt`                                                                                                                                                                                                                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                                               | Date and time in ISO 8601 format.                                                                                                                                                                                                                                                | 2020-04-09T18:14:30Z                                                                                                                                                                                                                                                             |
| `Version`                                                                                                                                                                                                                                                                        | **int64*                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                               | N/A                                                                                                                                                                                                                                                                              | 1                                                                                                                                                                                                                                                                                |