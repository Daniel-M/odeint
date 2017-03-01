package main

import "fmt"

import "github.com/Daniel-M/odeint"

//func odesys(params odeint.State, state odeint.State) odeint.State {
func odesys(state odeint.State) odeint.State {
	//var x []odeint.Float
	var vector odeint.State
	//x = make([]odeint.Float, 2)

	x := state.StateVector

	//fmt.Println(x)

	x[0] = -x[1]
	x[1] = 0 * x[0]

	vector.SetState(x)

	//fmt.Println(x, vector)

	return vector
}

func main() {

	var vector odeint.State
	vector.SetState([]odeint.Float{1, 2})

	fmt.Println("State", vector)

	var integrator odeint.Euler
	err := integrator.Set(0.02, vector, odesys)
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
	fmt.Println("Integrator state", integrator.State())

	integrator.Step()

	vector = integrator.State()
	fmt.Println("After Iteration", integrator.State())

}
