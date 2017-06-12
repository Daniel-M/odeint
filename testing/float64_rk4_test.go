package odeint

//import "fmt"
import "testing"

func TestNewRk4(t *testing.T) {
	//t.Log("Testing NewRk4()...")

	var step float64 = 0.01

	// Temporary variables
	vector := make([]float64, 2)
	params := make([]float64, 2)

	// State vector
	vector[0] = 2.0
	vector[1] = 3.0

	// Parameters vector
	params[0] = vector[0]
	params[1] = vector[1]

	// Get a NewSystem. Previous test checks that
	// NewSystem works properly
	sys := NewSystem(vector, params, func(x []float64, parameters []float64) []float64 {

		dxdt := make([]float64, len(x))

		dxdt[0] = parameters[0] * x[1]
		dxdt[1] = parameters[1] * x[0]

		return dxdt
	})

	rk4u := NewRk4(step, *sys)
	if rk4u.State()[0] != vector[0] && rk4u.State()[1] != vector[1] {
		t.Errorf("NewRk4() has wrong state...")
		t.Errorf("NewRk4() expected state [%f, %f]", vector[0], vector[1])
		t.Errorf("NewRk4() got state [%f, %f]", rk4u.State()[0], rk4u.State()[1])
	} else if step != rk4u.StepSize() {
		t.Errorf("NewRk4() has wrong step...")
		t.Errorf("NewRk4() expected step %f", step)
		t.Errorf("NewRk4() got step %f", rk4u.StepSize())
	}
}

func TestRk4Set(t *testing.T) {
	//t.Log("Testing Rk4.Set()...")
	var step float64 = 0.01

	// Temporary variables
	vector := make([]float64, 2)
	params := make([]float64, 2)

	// State vector
	vector[0] = 2.0
	vector[1] = 3.0

	// Parameters vector
	params[0] = vector[0]
	params[1] = vector[1]

	// Get a NewSystem. Previous test checks that
	// NewSystem works properly
	sys := NewSystem(vector, params, func(x []float64, parameters []float64) []float64 {

		dxdt := make([]float64, len(x))

		dxdt[0] = parameters[0] * x[1]
		dxdt[1] = parameters[1] * x[0]

		return dxdt
	})

	var rk4u Rk4

	if rk4u.Set(step, *sys) != nil {
		//t.Fail()
		t.Errorf("Rk4.Set() not returned nil. Return value ")
	}
}

func TestRk4Step(t *testing.T) {
	//t.Log("Testing Rk4.Set()...")

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
	stateK2 := make([]float64, len(vector))
	stateK3 := make([]float64, len(vector))
	stateK4 := make([]float64, len(vector))

	BufferK1 := make([]float64, len(vector))
	BufferK2 := make([]float64, len(vector))
	BufferK3 := make([]float64, len(vector))

	// First Step
	stateK1[0] = params[0] * vector[1]
	stateK1[1] = params[1] * vector[0]

	BufferK1[0] = vector[0] + 0.5*step*stateK1[0]
	BufferK1[1] = vector[1] + 0.5*step*stateK1[1]

	// Second Step
	stateK2[0] = params[0] * BufferK1[1]
	stateK2[1] = params[1] * BufferK1[0]

	BufferK2[0] = vector[0] + 0.5*step*stateK2[0]
	BufferK2[1] = vector[1] + 0.5*step*stateK2[1]

	// Third Step
	stateK3[0] = params[0] * BufferK2[1]
	stateK3[1] = params[1] * BufferK2[0]

	// Fourth Step
	BufferK3[0] = vector[0] + step*stateK3[0]
	BufferK3[1] = vector[1] + step*stateK3[1]

	stateK4[0] = params[0] * BufferK3[1]
	stateK4[1] = params[1] * BufferK3[0]

	// Result
	expected[0] = vector[0] + (step/6.0)*(stateK1[0]+2*stateK2[0]+2*stateK3[0]+stateK4[0])
	expected[1] = vector[1] + (step/6.0)*(stateK1[1]+2*stateK2[1]+2*stateK3[1]+stateK4[1])

	// Get a NewSystem. Previous test checks that
	// NewSystem works properly
	sys := NewSystem(vector, params, func(x []float64, parameters []float64) []float64 {

		dxdt := make([]float64, len(x))

		dxdt[0] = parameters[0] * x[1]
		dxdt[1] = parameters[1] * x[0]

		return dxdt
	})

	var rk4u Rk4

	if rk4u.Set(step, *sys) != nil {
		//t.Fail()
		t.Errorf("Rk4.Set() did not returned nil...")
	}

	//t.Log("Now testing Rk4.Step()...")

	//fmt.Println(vector, stateK1, stateK2, stateK3, stateK4, expected)

	state, err := rk4u.Step()
	//fmt.Println(state)

	if err != nil {
		t.Errorf("Rk4.Step() did not returned nil...")
	} else if state[0] != expected[0] && state[1] != expected[1] {
		t.Errorf("Rk4.Step() expected %f, got %f", expected[0], state[0])
		t.Errorf("Rk4.Step() expected %f, got %f", expected[1], state[1])
	}
}
