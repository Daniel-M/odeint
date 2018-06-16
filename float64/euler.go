// Copyright 2017-2018 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

//import "fmt"

// Euler implements the Euler stepper method. Euler is part of the Stepper
// interface.
type Euler struct {
	stepSize float64
	system   System
}

// NewEuler returns a reference to a new Euler stepper method.
func NewEuler(stepSize float64, system System) (r *Euler) {
	if stepSize <= 0 {
		panic("NewEuler called with negative or null stepSize.")
	}
	if system.Function() == nil {
		panic("NewEuler called with nil system.Function.")
	}
	r = &Euler{stepSize: stepSize, system: system}
	return
}

// These methods are required by the interface Stepper

// Setter methods

// SetStep sets the state for the Euler stepper.
func (euler *Euler) SetStep(step float64) error {
	euler.stepSize = step
	return nil
}

// SetState sets the state for the Euler stepper.
func (euler *Euler) SetState(state []float64) error {
	euler.system.stateVector = state
	return nil
}

// Set sets the step size and System for the Euler stepper.
func (euler *Euler) Set(stepSize float64, system System) error {
	if stepSize <= 0 {
		return &Error{"Can't use negative or null step size for Euler stepper."}
	} else {
		euler.stepSize = stepSize
		euler.system = system
		return nil
	}
}

// Getter methods

// StepSize returns the step size of the method
func (euler *Euler) StepSize() float64 {
	return euler.stepSize
}

// State returns the state of the Euler stepper.
func (euler *Euler) State() []float64 {
	return euler.system.stateVector
}

// Step performs one step iteration call of the Euler stepper.
// It also updates the state of the Euler object.
//
// The Euler method for the system,
// 		y = f(t, y)
// Consists of building the sequence of numbers t_n, y_n,
// following the recurrence,
// 		t_n+1 = t_n + dt
// 		y_n+1 = y_n + dt*f(t_n, y_n) // First step
func (euler *Euler) Step() ([]float64, error) {

	newstate := make([]float64, len(euler.system.stateVector))

	for i := 0; i < len(euler.system.stateVector); i++ {
		newstate[i] = euler.system.stateVector[i] + euler.stepSize*euler.system.Evaluate(euler.system.stateVector)[i]
	}

	euler.system.stateVector = newstate

	return newstate, nil
}
