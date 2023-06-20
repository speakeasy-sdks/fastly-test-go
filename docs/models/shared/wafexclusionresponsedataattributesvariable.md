# WafExclusionResponseDataAttributesVariable

The variable to exclude. An optional selector can be specified after the variable separated by a colon (`:`) to restrict the variable to a particular parameter. Required for `exclusion_type=variable`.


## Values

| Name                                                               | Value                                                              |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `WafExclusionResponseDataAttributesVariableReqCookies`             | req.cookies                                                        |
| `WafExclusionResponseDataAttributesVariableReqHeaders`             | req.headers                                                        |
| `WafExclusionResponseDataAttributesVariableReqPost`                | req.post                                                           |
| `WafExclusionResponseDataAttributesVariableReqPostFilename`        | req.post_filename                                                  |
| `WafExclusionResponseDataAttributesVariableReqQs`                  | req.qs                                                             |
| `WafExclusionResponseDataAttributesVariableLessThanNilGreaterThan` | <nil>                                                              |