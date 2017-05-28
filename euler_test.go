package odeint

//import "fmt"
import "testing"

func odesys(x []Float, parameters []Float) []Float {

	dxdt := make([]Float, len(x))

	dxdt[0] = parameters[0] * x[1]
	dxdt[1] = parameters[1] * x[0]

	return dxdt
}

func TestEulerSet(t *testing.T) {
	t.Log("Testing Euler.Set()...")

	var vector []Float
	var params []Float
	var step Float

	vector = make([]Float, 2)
	params = make([]Float, 2)
	step = 0.01

	var eu Euler

	if eu.Set(step, vector, params, odesys) != nil {
		//t.Fail()
		t.Errorf("Euler.Set() returned nil...")
	}

	//eu := NewEuler(step, vector, params, odesys)

}

func TestEulerStep(t *testing.T) {
	t.Log("Testing Euler.Set()...")

	var step Float

	vector := make([]Float, 2)
	expected := make([]Float, 2)
	params := make([]Float, 2)

	// Seed initial values to call odesys
	vector[0] = 1.0
	vector[1] = 1.0

	// Parameters
	params[0] = 1.0
	params[1] = 0.0

	// A step for the simulation
	step = 0.01

	// This is what expect in the first simulation
	expected[0] = vector[0] + 0.01*params[0]*vector[1]
	expected[1] = vector[1] + 0.01*params[1]*vector[0]

	var eu Euler

	if eu.Set(step, vector, params, odesys) != nil {
		//t.Fail()
		t.Errorf("Euler.Set() returned nil...")
	}

	t.Log("Now testing Euler.Step()...")

	state, err := eu.Step()

	if err != nil {
		t.Errorf("Euler.Step() returned nil...")
	} else if state[0] != expected[0] && state[1] != expected[1] {
		t.Errorf("Euler.Step() expected %f, got %f", expected[0], state[0])
		t.Errorf("Euler.Step() expected %f, got %f", expected[1], state[1])
	}
}
