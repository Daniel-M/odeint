package odeint

//import "fmt"
import "testing"

func TestNewEuler(t *testing.T) {
	//t.Log("Testing NewEuler()...")

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

	eu := NewEuler(step, *sys)
	if eu.State()[0] != vector[0] && eu.State()[1] != vector[1] {
		t.Errorf("NewEuler() has wrong state...")
		t.Errorf("NewEuler() expected state [%f, %f]", vector[0], vector[1])
		t.Errorf("NewEuler() got state [%f, %f]", eu.State()[0], eu.State()[1])
	} else if step != eu.StepSize() {
		t.Errorf("NewEuler() has wrong step...")
		t.Errorf("NewEuler() expected step %f", step)
		t.Errorf("NewEuler() got step %f", eu.StepSize())
	}

}

func TestEulerSet(t *testing.T) {
	//t.Log("Testing Euler.Set()...")
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

	var eu Euler

	if eu.Set(step, *sys) != nil {
		//t.Fail()
		t.Errorf("Euler.Set() not returned nil. Return value ")
	}

	//eu := NewEuler(step, vector, params, odesys)

}

func TestEulerStep(t *testing.T) {
	//t.Log("Testing Euler.Set()...")

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
	params[0] = vector[0]
	params[1] = vector[1]

	// This is what we expect in the first step of the simulation
	expected[0] = vector[0] + 0.01*params[0]*vector[1]
	expected[1] = vector[1] + 0.01*params[1]*vector[0]

	// Get a NewSystem. Previous test checks that
	// NewSystem works properly
	sys := NewSystem(vector, params, func(x []float64, parameters []float64) []float64 {

		dxdt := make([]float64, len(x))

		dxdt[0] = parameters[0] * x[1]
		dxdt[1] = parameters[1] * x[0]

		return dxdt
	})

	var eu Euler

	//if eu.Set(step, vector, params, odesys) != nil {
	if eu.Set(step, *sys) != nil {
		//t.Fail()
		t.Errorf("Euler.Set() did not returned nil...")
	}

	//t.Log("Now testing Euler.Step()...")

	state, err := eu.Step()

	if err != nil {
		t.Errorf("Euler.Step() did not returned nil...")
	} else if state[0] != expected[0] && state[1] != expected[1] {
		t.Errorf("Euler.Step() expected %f, got %f", expected[0], state[0])
		t.Errorf("Euler.Step() expected %f, got %f", expected[1], state[1])
	}
}
