# Action

If set, will cause vcl_fetch to terminate after processing this rule with the return state specified. If not set, other configuration logic in vcl_fetch with a lower priority will run after this rule.



## Values

| Name            | Value           |
| --------------- | --------------- |
| `ActionPass`    | pass            |
| `ActionCache`   | cache           |
| `ActionRestart` | restart         |