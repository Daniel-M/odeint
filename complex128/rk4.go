// Copyright 2017 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

//import "fmt"

// Rk4 implements the rk4 stepper method. rk4 is part of the Stepper
// interface.
type Rk4 struct {
	stepSize complex128
	system   System
}

// NewRk4 returns a reference to a new rk4 stepper method.
func NewRk4(stepSize complex128, system System) (r *Rk4) {
	//if stepSize <= 0 {
	if real(stepSize) <= 0 && imag(stepSize) <= 0.0 {
		panic("NewRk4 called with negative or null stepSize.")
	}
	if system.Function() == nil {
		panic("NewRk4 called with nil system.Function.")
	}
	r = &Rk4{stepSize: stepSize, system: system}
	return
}

// These methods are required by the interface Stepper

// Setter methods

// SetStep sets the state for the Rk4 stepper.
func (rk4 *Rk4) SetStep(step complex128) error {
	rk4.stepSize = step
	return nil
}

// SetState sets the state for the Rk4 stepper.
func (rk4 *Rk4) SetState(state []complex128) error {
	rk4.system.stateVector = state
	return nil
}

// Set sets the step size and System for the Rk4 stepper.
func (rk4 *Rk4) Set(stepSize complex128, system System) error {
	//if stepSize <= 0 {
	if real(stepSize) <= 0 && imag(stepSize) <= 0.0 {
		return &Error{"Can't use negative or null step size for Rk4 stepper."}
	} else {
		rk4.stepSize = stepSize
		rk4.system = system
		return nil
	}
}

// Getter methods

// StepSize returns the step size of the method
func (rk4 *Rk4) StepSize() complex128 {
	return rk4.stepSize
}

// State returns the state of the Rk4 stepper.
func (rk4 *Rk4) State() []complex128 {
	return rk4.system.stateVector
}

// Step performs one step iteration call of the Rk4 stepper.
// It also updates the state of the Rk4 object.
func (rk4 *Rk4) Step() ([]complex128, error) {

	newstate := rk4.system.stateVector
	stateK1 := make([]complex128, len(rk4.system.stateVector))
	stateK2 := make([]complex128, len(rk4.system.stateVector))
	stateK3 := make([]complex128, len(rk4.system.stateVector))
	stateK4 := make([]complex128, len(rk4.system.stateVector))

	BufferK1 := make([]complex128, len(rk4.system.stateVector))
	BufferK2 := make([]complex128, len(rk4.system.stateVector))
	BufferK3 := make([]complex128, len(rk4.system.stateVector))

	// First Step
	stateK1 = rk4.system.Evaluate(newstate)

	// Second Step
	for i := 0; i < len(rk4.system.stateVector); i++ {
		BufferK1[i] = newstate[i] + 0.5*rk4.stepSize*stateK1[i]
	}

	stateK2 = rk4.system.Evaluate(BufferK1)

	// Third Step
	for i := 0; i < len(rk4.system.stateVector); i++ {
		BufferK2[i] = newstate[i] + 0.5*rk4.stepSize*stateK2[i]
	}

	stateK3 = rk4.system.Evaluate(BufferK2)

	// Fourth Step
	for i := 0; i < len(rk4.system.stateVector); i++ {
		BufferK3[i] = newstate[i] + rk4.stepSize*stateK3[i]
	}

	stateK4 = rk4.system.Evaluate(BufferK3)

	// Result
	for i := 0; i < len(rk4.system.stateVector); i++ {
		newstate[i] = newstate[i] + (rk4.stepSize/6.0)*(stateK1[i]+2*stateK2[i]+2*stateK3[i]+stateK4[i])
	}

	rk4.system.stateVector = newstate

	return newstate, nil
}
