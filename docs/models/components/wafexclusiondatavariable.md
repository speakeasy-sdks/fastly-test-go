# WafExclusionDataVariable

The variable to exclude. An optional selector can be specified after the variable separated by a colon (`:`) to restrict the variable to a particular parameter. Required for `exclusion_type=variable`.


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `WafExclusionDataVariableReqCookies`      | req.cookies                               |
| `WafExclusionDataVariableReqHeaders`      | req.headers                               |
| `WafExclusionDataVariableReqPost`         | req.post                                  |
| `WafExclusionDataVariableReqPostFilename` | req.post_filename                         |
| `WafExclusionDataVariableReqQs`           | req.qs                                    |