// Copyright 2017 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

//import "fmt"

// Euler implements the Euler stepper method. Euler is part of the Stepper
// interface.
type Euler struct {
	stepSize    Float
	stateVector []Float
	ODESystem   System
	//ParametersVector []Float
	//ODESystem        System
}

// NewEuler returns a new Euler stepper method.
func NewEuler(stepSize Float, stateVector []Float, oDESystem System) (r *Euler) {
	//func NewEuler(stepSize Float, stateVector []Float, parametersVector []Float, oDESystem System) (r *Euler) {
	if stepSize <= 0 {
		panic("NewEuler called with negative or null stepSize.")
	}
	if oDESystem.Function == nil {
		panic("NewEuler called with nil ODESystem.")
	}
	//r = &Euler{stepSize: stepSize, stateVector: stateVector, ParametersVector: parametersVector, ODESystem: oDESystem}
	r = &Euler{stepSize: stepSize, stateVector: stateVector, ODESystem: oDESystem}
	return
}

// Set sets the state for the Euler stepper.
func (euler *Euler) Set(stepSize Float, stateVector []Float, oDESystem System) error {
	//func (euler *Euler) Set(stepSize Float, stateVector []Float, parametersVector []Float, oDESystem System) error {
	if stepSize <= 0 {
		return &Error{"Can't use negative step size for Euler stepper."}
	} else {
		euler.stepSize = stepSize
		euler.stateVector = stateVector
		//euler.ParametersVector = parametersVector
		euler.ODESystem = oDESystem
		return nil
	}
}

func (euler *Euler) StepSize() Float {
	return euler.stepSize
}

// SetState sets the state for the Euler stepper.
func (euler *Euler) SetState(state []Float) error {
	euler.stateVector = state
	return nil
}

// SetStep sets the state for the Euler stepper.
func (euler *Euler) SetStep(step Float) error {
	euler.stepSize = step
	return nil
}

// State returns the state of the Euler stepper.
func (euler *Euler) State() []Float {
	return euler.stateVector
}

// Step performs one step iteration call of the Euler stepper.
// It also updates the state of the Euler object.
func (euler *Euler) Step() ([]Float, error) {

	newstate := make([]Float, len(euler.stateVector))

	for i := 0; i < len(euler.stateVector); i++ {
		//newstate[i] = euler.stateVector[i] + euler.stepSize*euler.ODESystem(euler.stateVector, euler.ParametersVector)[i]
		newstate[i] = euler.stateVector[i] + euler.stepSize*euler.ODESystem.Evaluate(euler.stateVector)[i]
	}

	euler.stateVector = newstate

	return newstate, nil
}
