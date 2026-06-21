# go-tool-interface

Small Go SDK for registering external Open Chat tools.

## Usage

```go
package mytool

import "github.com/msgmate-io/go-tool-interface/toolinterface"

func init() {
	toolinterface.MustRegister(toolinterface.Definition{
		Name:        "my_tool",
		Description: "My external tool",
		InputType: struct {
			Message string `json:"message"`
		}{},
		RequiredParams: []string{"message"},
		Parameters: map[string]interface{}{
			"message": map[string]interface{}{"type": "string"},
		},
		Run: func(input interface{}, init map[string]interface{}) (string, error) {
			in := input.(struct {
				Message string `json:"message"`
			})
			return "echo: " + in.Message, nil
		},
	})
}
```

At build time, add your package import path in `backend/tooldeps.json`.
