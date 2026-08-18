# go-tool-interface

Small Go SDK for registering external Open Chat tools.

This module remains independent and can be used by any Go project.

In Open Chat, tool definitions are discovered via integrations (`go-integration-interface`), not via a separate `tooldeps.json` manifest.

Migration for existing tool-only subrepos:

1. Keep your tool definitions in this module.
2. Add an integration package that imports those definitions.
3. Register them via `integrationinterface.Definition{ToolDefinitions: ...}`.
4. Add that integration module to `backend/integrationdeps.json`.

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

To expose tools in Open Chat, include the definitions in an integration's `Definition.ToolDefinitions` field.
