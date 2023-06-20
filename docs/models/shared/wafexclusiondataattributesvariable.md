# WafExclusionDataAttributesVariable

The variable to exclude. An optional selector can be specified after the variable separated by a colon (`:`) to restrict the variable to a particular parameter. Required for `exclusion_type=variable`.


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `WafExclusionDataAttributesVariableReqCookies`             | req.cookies                                                |
| `WafExclusionDataAttributesVariableReqHeaders`             | req.headers                                                |
| `WafExclusionDataAttributesVariableReqPost`                | req.post                                                   |
| `WafExclusionDataAttributesVariableReqPostFilename`        | req.post_filename                                          |
| `WafExclusionDataAttributesVariableReqQs`                  | req.qs                                                     |
| `WafExclusionDataAttributesVariableLessThanNilGreaterThan` | <nil>                                                      |