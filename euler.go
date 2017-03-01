// Copyright 2017 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

//import "fmt"

// Euler implements the Euler stepper method. Euler is part of the Stepper
// interface.
type Euler struct {
	StepSize         Float
	StateVector      State
	ParametersVector State
	ODESystem        System
}

// Set sets the state for the Euler stepper.
func (euler *Euler) Set(step Float, state State, ode System) error {
	if step <= 0 {
		return &Error{"Can't use negative step size for Euler stepper."}
	} else {
		euler.StepSize = step
		euler.StateVector = state
		euler.ODESystem = ode
		return nil
	}
}

// SetState sets the state for the Euler stepper.
func (euler *Euler) SetState(state State) error {
	euler.StateVector = state
	return nil
}

// State returns the state of the Euler stepper.
func (euler *Euler) State() State {
	return euler.StateVector
}

// Step performs one step iteration call of the Euler stepper.
func (euler *Euler) Step() error {
	return nil
}
