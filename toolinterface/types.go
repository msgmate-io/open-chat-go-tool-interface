package toolinterface

type RunFunc func(input interface{}, init map[string]interface{}) (string, error)

// ConfirmationResolver resolves a tool's confirmation requirement at call
// time. Unlike RequiresConfirmation (a static flag which may be persisted as
// user config), a resolver is supplied by the tool's own code and can react to
// runtime state (eg the bound session's interaction mode). When it returns
// requiresConfirmation=false together with a non-empty resolvedResult, the core
// uses resolvedResult as the tool output and continues without pausing for the
// user; when resolvedResult is empty the tool's Run is invoked normally.
type ConfirmationResolver func(input interface{}, init map[string]interface{}) (requiresConfirmation bool, resolvedResult string)

type Definition struct {
	Name                           string
	FunctionName                   string
	Description                    string
	AdminOnly                      bool
	RequiresInit                   bool
	InitSchema                     map[string]interface{}
	RequiresConfirmation           bool
	StopOnFirstConfirmableToolCall bool
	ConfirmationBlockMessage       string
	InputType                      interface{}
	RequiredParams                 []string
	Parameters                     map[string]interface{}
	Run                            RunFunc
	// ResolveConfirmation, when set, overrides RequiresConfirmation at call
	// time (code-defined, never user config).
	ResolveConfirmation ConfirmationResolver
}
