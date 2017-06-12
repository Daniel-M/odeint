package odeint

//import "fmt"
import "testing"

func TestNewSTEPPER_METHOD(t *testing.T) {
	//t.Log("Testing NewSTEPPER_METHOD()...")

	var step Float = 0.01

	// Temporary variables
	vector := make([]Float, 2)
	params := make([]Float, 2)

	// State vector
	vector[0] = 2.0
	vector[1] = 3.0

	// Parameters vector
	params[0] = vector[0]
	params[1] = vector[1]

	// Get a NewSystem. Previous test checks that
	// NewSystem works properly
	sys := NewSystem(vector, params, func(x []Float, parameters []Float) []Float {

		dxdt := make([]Float, len(x))

		dxdt[0] = parameters[0] * x[1]
		dxdt[1] = parameters[1] * x[0]

		return dxdt
	})

	stepper_method := NewSTEPPER_METHOD(step, *sys)
	if stepper_method.State()[0] != vector[0] && stepper_method.State()[1] != vector[1] {
		t.Errorf("NewSTEPPER_METHOD() has wrong state...")
		t.Errorf("NewSTEPPER_METHOD() expected state [%f, %f]", vector[0], vector[1])
		t.Errorf("NewSTEPPER_METHOD() got state [%f, %f]", stepper_method.State()[0], stepper_method.State()[1])
	} else if step != stepper_method.StepSize() {
		t.Errorf("NewSTEPPER_METHOD() has wrong step...")
		t.Errorf("NewSTEPPER_METHOD() expected step %f", step)
		t.Errorf("NewSTEPPER_METHOD() got step %f", stepper_method.StepSize())
	}
}

func TestSTEPPER_METHODSet(t *testing.T) {
	//t.Log("Testing STEPPER_METHOD.Set()...")
	var step Float = 0.01

	// Temporary variables
	vector := make([]Float, 2)
	params := make([]Float, 2)

	// State vector
	vector[0] = 2.0
	vector[1] = 3.0

	// Parameters vector
	params[0] = vector[0]
	params[1] = vector[1]

	// Get a NewSystem. Previous test checks that
	// NewSystem works properly
	sys := NewSystem(vector, params, func(x []Float, parameters []Float) []Float {

		dxdt := make([]Float, len(x))

		dxdt[0] = parameters[0] * x[1]
		dxdt[1] = parameters[1] * x[0]

		return dxdt
	})

	var stepper_method STEPPER_METHOD

	if stepper_method.Set(step, *sys) != nil {
		//t.Fail()
		t.Errorf("STEPPER_METHOD.Set() not returned nil. Return value ")
	}
}

func TestSTEPPER_METHODStep(t *testing.T) {
	//t.Log("Testing STEPPER_METHOD.Set()...")

	// A step for the simulation
	var step Float = 0.01

	// Temporary variables
	vector := make([]Float, 2)
	params := make([]Float, 2)
	expected := make([]Float, 2)

	// State vector
	vector[0] = 2.0
	vector[1] = 3.0

	// Parameters vector
	params[0] = vector[1]
	params[1] = vector[0]

	// This is what we expect in the first step of the simulation

	stateK1 := make([]Float, len(vector))
	stateK2 := make([]Float, len(vector))
	stateK3 := make([]Float, len(vector))
	stateK4 := make([]Float, len(vector))

	BufferK1 := make([]Float, len(vector))
	BufferK2 := make([]Float, len(vector))
	BufferK3 := make([]Float, len(vector))

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
	sys := NewSystem(vector, params, func(x []Float, parameters []Float) []Float {

		dxdt := make([]Float, len(x))

		dxdt[0] = parameters[0] * x[1]
		dxdt[1] = parameters[1] * x[0]

		return dxdt
	})

	var stepper_method STEPPER_METHOD

	if stepper_method.Set(step, *sys) != nil {
		//t.Fail()
		t.Errorf("STEPPER_METHOD.Set() did not returned nil...")
	}

	//t.Log("Now testing STEPPER_METHOD.Step()...")

	//fmt.Println(vector, stateK1, stateK2, stateK3, stateK4, expected)

	state, err := stepper_method.Step()
	//fmt.Println(state)

	if err != nil {
		t.Errorf("STEPPER_METHOD.Step() did not returned nil...")
	} else if state[0] != expected[0] && state[1] != expected[1] {
		t.Errorf("STEPPER_METHOD.Step() expected %f, got %f", expected[0], state[0])
		t.Errorf("STEPPER_METHOD.Step() expected %f, got %f", expected[1], state[1])
	}
}
