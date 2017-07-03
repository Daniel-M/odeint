package main

import "fmt"

import "github.com/Daniel-M/odeint/float64"

// Renaming the type just for short code writting
//type float64 odeint.float64

func odesys(x []float64, parameters []float64) []float64 {

	dxdt := make([]float64, len(x))

	dxdt[0] = parameters[0] + x[0]*x[0]*x[1] - (parameters[1]+1)*x[0]
	dxdt[1] = (parameters[1] - x[0]*x[1]) * x[0]

	return dxdt
}

func main() {

	state := make([]float64, 2)
	params := make([]float64, 2)

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
		fmt.Println(float64(i)*integrator.StepSize(), state)

		state, err = integrator.Step()
		if err != nil {
			panic(err)
		}
	}
}
