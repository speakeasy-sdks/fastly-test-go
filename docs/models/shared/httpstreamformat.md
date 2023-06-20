# HTTPStreamFormat

Payload format for delivering to subscribers of HTTP streaming response bodies (`stream` hold mode). One of `content` or `content-bin` must be specified.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Content`                                                | **string*                                                | :heavy_minus_sign:                                       | A fragment of body data as a string.                     |
| `ContentBin`                                             | **string*                                                | :heavy_minus_sign:                                       | A fragment of body data as a base64-encoded binary blob. |