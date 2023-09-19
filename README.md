# traceparent

Generates a traceparent value for usage in things like curl

## Usage

`go install github.com/supershabam/traceparent/cmd/traceparent@v1.0.0`

curl -v -H "Traceparent: $(traceparent)" example.com
