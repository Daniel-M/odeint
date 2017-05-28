// Copyright 2017 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

//import "fmt"

// Euler implements the Euler stepper method. Euler is part of the Stepper
// interface.
type Euler struct {
	StepSize         Float
	StateVector      []Float
	ParametersVector []Float
	ODESystem        System
}

// NewEuler returns a new Euler stepper method.
func NewEuler(stepSize Float, stateVector []Float, parametersVector []Float, oDESystem System) (r *Euler) {
	if stepSize <= 0 {
		panic("NewEuler called with negative or null stepSize.")
	}
	if oDESystem == nil {
		panic("NewEuler called with nil ODESystem.")
	}
	r = &Euler{StepSize: stepSize, StateVector: stateVector, ParametersVector: parametersVector, ODESystem: oDESystem}
	return
}

// Set sets the state for the Euler stepper.
func (euler *Euler) Set(stepSize Float, stateVector []Float, parametersVector []Float, oDESystem System) error {
	if stepSize <= 0 {
		return &Error{"Can't use negative step size for Euler stepper."}
	} else {
		euler.StepSize = stepSize
		euler.StateVector = stateVector
		euler.ParametersVector = parametersVector
		euler.ODESystem = oDESystem
		return nil
	}
}

// SetState sets the state for the Euler stepper.
func (euler *Euler) SetState(state []Float) error {
	euler.StateVector = state
	return nil
}

// State returns the state of the Euler stepper.
func (euler *Euler) State() []Float {
	return euler.StateVector
}

// Step performs one step iteration call of the Euler stepper.
// It also updates the state of the Euler object.
func (euler *Euler) Step() ([]Float, error) {

	newstate := make([]Float, len(euler.StateVector))

	for i := 0; i < len(euler.StateVector); i++ {
		newstate[i] = euler.StateVector[i] + euler.StepSize*euler.ODESystem(euler.StateVector, euler.ParametersVector)[i]
	}

	euler.StateVector = newstate

	return newstate, nil
}
