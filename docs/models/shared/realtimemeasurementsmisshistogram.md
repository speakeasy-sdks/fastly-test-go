# RealtimeMeasurementsMissHistogram

A histogram. Each key represents the upper bound of a span of 10 milliseconds and the values represent the number of requests to origin during that 10ms period. Any origin request that takes more than 60 seconds to return will be in the 60000 bucket.


## Fields

| Field       | Type        | Required    | Description |
| ----------- | ----------- | ----------- | ----------- |