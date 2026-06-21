package toolinterface

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

var (
	mu          sync.RWMutex
	definitions = map[string]Definition{}
)

func Register(def Definition) error {
	name := strings.TrimSpace(def.Name)
	if name == "" {
		return fmt.Errorf("tool definition requires a non-empty name")
	}
	if def.Run == nil {
		return fmt.Errorf("tool definition '%s' requires a run function", name)
	}

	mu.Lock()
	defer mu.Unlock()
	if _, exists := definitions[name]; exists {
		return fmt.Errorf("tool definition '%s' already registered", name)
	}

	def.Name = name
	definitions[name] = def
	return nil
}

func MustRegister(def Definition) {
	if err := Register(def); err != nil {
		panic(err)
	}
}

func List() []Definition {
	mu.RLock()
	defer mu.RUnlock()

	names := make([]string, 0, len(definitions))
	for name := range definitions {
		names = append(names, name)
	}
	sort.Strings(names)

	result := make([]Definition, 0, len(names))
	for _, name := range names {
		result = append(result, definitions[name])
	}
	return result
}
