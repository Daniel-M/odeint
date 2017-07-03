package main

import "fmt"

import "github.com/Daniel-M/odeint/float32"

func odesys(x []float32, parameters []float32) []float32 {

	dxdt := make([]float32, len(x))

	dxdt[0] = x[1]
	dxdt[1] = -parameters[0]*x[0] - parameters[1]*x[1]

	return dxdt
}

func main() {

	state := make([]float32, 2)
	params := make([]float32, 2)

	// State vector
	state[0] = 0.2
	state[1] = 0.8

	// Parameters vector
	params[0] = 1.2 * 1.2
	params[1] = 0.2

	system := odeint.NewSystem(state, params, odesys)

	fmt.Println("#Parameters", state)
	fmt.Println("#State", state)

	var integrator odeint.Midpoint
	err := integrator.Set(0.1, *system)
	if err != nil {
		panic(err)
	}

	for i := 0; i < int(30.0/integrator.StepSize()); i++ {
		fmt.Println(float32(i)*integrator.StepSize(), state)

		state, err = integrator.Step()
		if err != nil {
			panic(err)
		}
	}
}
