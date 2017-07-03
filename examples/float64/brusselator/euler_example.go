package main

import "fmt"

import "github.com/Daniel-M/odeint"

// Renaming the type just for short code writting
//type odeint.Float odeint.odeint.Float

func odesys(x []odeint.Float, parameters []odeint.Float) []odeint.Float {

	dxdt := make([]odeint.Float, len(x))

	dxdt[0] = parameters[0] + x[0]*x[0]*x[1] - (parameters[1]+1)*x[0]
	dxdt[1] = (parameters[1] - x[0]*x[1]) * x[0]

	return dxdt
}

func main() {

	state := make([]odeint.Float, 2)
	params := make([]odeint.Float, 2)

	// State vector
	state[0] = 1.0
	state[1] = 1.0

	// Parameters vector
	params[0] = 1.0
	params[1] = 3.0

	system := odeint.NewSystem(state, params, odesys)

	fmt.Println("#Parameters", state)
	fmt.Println("#State", state)

	var integrator odeint.Euler
	err := integrator.Set(0.1, *system)
	if err != nil {
		panic(err)
	}

	for i := 0; i < int(30.0/integrator.StepSize()); i++ {
		fmt.Println(odeint.Float(i)*integrator.StepSize(), state)

		state, err = integrator.Step()
		if err != nil {
			panic(err)
		}
	}
}
