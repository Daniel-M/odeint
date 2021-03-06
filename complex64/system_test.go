package odeint

//import "fmt"
import "testing"

func TestNewSystem(t *testing.T) {
	//t.Log("Testing NewSystem()...")

	// Temporary variables
	vector := make([]complex64, 2)
	params := make([]complex64, 2)

	// State vector
	vector[0] = 1.0
	vector[1] = 2.0

	// Parameters vector
	params[0] = vector[0] - 1i
	params[1] = vector[1] + 1i

	// Temporary function to assign to the system
	f := func(x []complex64, parameters []complex64) []complex64 {

		dxdt := make([]complex64, len(x))

		dxdt[0] = parameters[0] * x[1]
		dxdt[1] = parameters[1] * x[0]

		return dxdt
	}

	// Here we use NewSystem
	sys := NewSystem(vector, params, f)

	// Check if the function f and the one of the system return the same
	res1 := f(vector, params)
	//res2 := sys.Function()(vector, params)
	res2 := sys.Evaluate(vector)

	// Check the results of the NewSystem created
	if sys.Parameters()[0] != (vector[0]-1i) && sys.Parameters()[1] != (vector[1]+1i) {
		t.Errorf("NewSystem() returned wrong parameters vector...")
	} else if res1[0] != res2[0] && res1[1] != res2[1] {
		t.Errorf("NewSystem() returned wrong function result...")
		t.Errorf("NewSystem() expected [%f, %f]", res1[0], res1[1])
		t.Errorf("NewSystem() obtained [%f, %f]", res2[0], res2[1])
	}
}

func TestSystemEvaluate(t *testing.T) {
	//t.Log("Testing System.Evaluate()...")

	// Temporary variables
	vector := make([]complex64, 2)
	params := make([]complex64, 2)

	// State vector
	vector[0] = 1.0
	vector[1] = 2.0

	// Parameters vector
	params[0] = vector[0] - 1i
	params[1] = vector[1] + 1i

	// Get a NewSystem. Previous test checks that
	// NewSystem works properly
	sys := NewSystem(vector, params, func(x []complex64, parameters []complex64) []complex64 {

		dxdt := make([]complex64, len(x))

		dxdt[0] = parameters[0] * x[1]
		dxdt[1] = parameters[1] * x[0]

		return dxdt
	})

	// Checks indirect function evaluation.
	res := sys.Evaluate(vector)

	if res[0] != params[0]*vector[1] && res[1] != params[0]*vector[1] {
		t.Errorf("System.Evaluate() failed...")
		t.Errorf("System.Evaluate() expected [%f, %f]", vector[0], vector[1])
		t.Errorf("System.Evaluate() got [%f, %f]", res[0], res[1])

	}
}
