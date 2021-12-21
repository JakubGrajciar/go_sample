package greeter

import "fmt"

// Greeter struct
type Greeter struct {
	// Exported field - first letter capital
	Greeting string // greeting string, e.g. hello

	// Private field - first letter lower case
	count int // total calls to Greet method
	// naming scheme uses CammelCase
	errCount int // errors in Greet method
}

// Create new greeter
func NewGreeter(greeting string) (*Greeter, error) {
	if len(greeting) == 0 {
		return nil, fmt.Errorf("missing greeting")
	}

	// Return Greeter as reference
	return &Greeter{
		Greeting: greeting,
		count:    0,
		errCount: 0,
	}, nil
}

// Return number of calls to Greet method
func (g *Greeter) GetCount() int {
	return g.count
}

// Return number of error in Greet method
func (g *Greeter) GetErrCount() int {
	return g.errCount
}

// Greet person e.g.: hello me!
func (g *Greeter) Greet(name string) (string, error) {
	// Increment counter when function exits
	defer func(g *Greeter) {
		g.count++
	}(g)

	if len(name) == 0 {
		g.errCount++
		return "nil", fmt.Errorf("missing name")
	}

	return fmt.Sprintf("%s %s!", g.Greeting, name), nil
}
