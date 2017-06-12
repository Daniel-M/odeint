package main

import "fmt"
import "math"

import "github.com/Daniel-M/odeint/complex64"

func odesys(x []complex64, parameters []complex64) []complex64 {

	dxdt := make([]complex64, len(x))

	dxdt[0] = -1i * x[0] * x[0] * x[1]
	dxdt[1] = 1i * x[1]

	return dxdt
}

func main() {

	state := make([]complex64, 2)
	params := make([]complex64, 2)

	// State vector
	state[0] = 1 / 2.1
	state[1] = 1.0

	// Parameters vector
	params[0] = 2.0
	params[1] = 0.0

	system := odeint.NewSystem(state, params, odesys)

	fmt.Println("#Parameters", state)
	fmt.Println("#State", state)

	var integrator odeint.Rk4
	err := integrator.Set(2*math.Pi/500, *system)
	if err != nil {
		panic(err)
	}

	for i := 0; i <= int(2*math.Pi/real(integrator.StepSize()))+1; i++ {
		fmt.Println(float32(i)*real(integrator.StepSize()), state)

		state, err = integrator.Step()
		if err != nil {
			panic(err)
		}
	}
}
