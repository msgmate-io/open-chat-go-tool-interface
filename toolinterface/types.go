package toolinterface

type RunFunc func(input interface{}, init map[string]interface{}) (string, error)

type Definition struct {
	Name                           string
	FunctionName                   string
	Description                    string
	AdminOnly                      bool
	RequiresInit                   bool
	RequiresConfirmation           bool
	StopOnFirstConfirmableToolCall bool
	ConfirmationBlockMessage       string
	InputType                      interface{}
	RequiredParams                 []string
	Parameters                     map[string]interface{}
	Run                            RunFunc
}
