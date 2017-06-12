package odeint

//import "fmt"
import "testing"

func TestNewMidpoint(t *testing.T) {
	//t.Log("Testing NewMidpoint()...")

	var step float64 = 0.01

	// Temporary variables
	vector := make([]float64, 2)
	params := make([]float64, 2)

	// State vector
	vector[0] = 2.0
	vector[1] = 3.0

	// Parameters vector
	params[0] = vector[1]
	params[1] = vector[0]

	// Get a NewSystem. Previous test checks that
	// NewSystem works properly
	sys := NewSystem(vector, params, func(x []float64, parameters []float64) []float64 {

		dxdt := make([]float64, len(x))

		dxdt[0] = parameters[0] * x[1]
		dxdt[1] = parameters[1] * x[0]

		return dxdt
	})

	mdpu := NewMidpoint(step, *sys)
	if mdpu.State()[0] != vector[0] && mdpu.State()[1] != vector[1] {
		t.Errorf("NewMidpoint() has wrong state...")
		t.Errorf("NewMidpoint() expected state [%f, %f]", vector[0], vector[1])
		t.Errorf("NewMidpoint() got state [%f, %f]", mdpu.State()[0], mdpu.State()[1])
	} else if step != mdpu.StepSize() {
		t.Errorf("NewMidpoint() has wrong step...")
		t.Errorf("NewMidpoint() expected step %f", step)
		t.Errorf("NewMidpoint() got step %f", mdpu.StepSize())
	}
}

func TestMidpointSet(t *testing.T) {
	//t.Log("Testing Midpoint.Set()...")
	var step float64 = 0.01

	// Temporary variables
	vector := make([]float64, 2)
	params := make([]float64, 2)

	// State vector
	vector[0] = 2.0
	vector[1] = 3.0

	// Parameters vector
	params[0] = vector[1]
	params[1] = vector[0]

	// Get a NewSystem. Previous test checks that
	// NewSystem works properly
	sys := NewSystem(vector, params, func(x []float64, parameters []float64) []float64 {

		dxdt := make([]float64, len(x))

		dxdt[0] = parameters[0] * x[1]
		dxdt[1] = parameters[1] * x[0]

		return dxdt
	})

	var mdpu Midpoint

	if mdpu.Set(step, *sys) != nil {
		//t.Fail()
		t.Errorf("Midpoint.Set() not returned nil. Return value ")
	}

}

func TestMidpointStep(t *testing.T) {
	//t.Log("Testing Midpoint.Set()...")

	// A step for the simulation
	var step float64 = 0.01

	// Temporary variables
	vector := make([]float64, 2)
	params := make([]float64, 2)
	expected := make([]float64, 2)

	// State vector
	vector[0] = 2.0
	vector[1] = 3.0

	// Parameters vector
	params[0] = vector[1]
	params[1] = vector[0]

	// This is what we expect in the first step of the simulation

	stateK1 := make([]float64, len(vector))

	BufferK1 := make([]float64, len(vector))

	// First Step
	stateK1[0] = params[0] * vector[1]
	stateK1[1] = params[1] * vector[0]

	stateK1[0] = vector[0] + 0.5*step*stateK1[0]
	stateK1[1] = vector[1] + 0.5*step*stateK1[1]

	// Second Step
	BufferK1[0] = params[0] * stateK1[1]
	BufferK1[1] = params[1] * stateK1[0]

	// Result
	expected[0] = vector[0] + step*BufferK1[0]
	expected[1] = vector[1] + step*BufferK1[1]

	// Get a NewSystem. Previous test checks that
	// NewSystem works properly
	sys := NewSystem(vector, params, func(x []float64, parameters []float64) []float64 {

		dxdt := make([]float64, len(x))

		dxdt[0] = parameters[0] * x[1]
		dxdt[1] = parameters[1] * x[0]

		return dxdt
	})

	var mdpu Midpoint

	if mdpu.Set(step, *sys) != nil {
		//t.Fail()
		t.Errorf("Midpoint.Set() did not returned nil...")
	}

	//t.Log("Now testing Midpoint.Step()...")

	//fmt.Println(vector, stateK1, expected)
	state, err := mdpu.Step()
	//fmt.Println(state)

	if err != nil {
		t.Errorf("Midpoint.Step() did not returned nil...")
	} else if state[0] != expected[0] && state[1] != expected[1] {
		t.Errorf("Midpoint.Step() expected %f, got %f", expected[0], state[0])
		t.Errorf("Midpoint.Step() expected %f, got %f", expected[1], state[1])
	}
}
