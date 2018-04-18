// Copyright 2017 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

//import "fmt"

// Midpoint implements the Midpoint stepper method. Midpoint is part of the Stepper
// interface.
type Midpoint struct {
	stepSize complex128
	system   System
}

// NewMidpoint returns a reference to a new Midpoint stepper method.
func NewMidpoint(stepSize complex128, system System) (r *Midpoint) {
	//if stepSize <= 0 {
	if real(stepSize) <= 0 && imag(stepSize) <= 0.0 {
		panic("NewMidpoint called with negative or null stepSize.")
	}
	if system.Function() == nil {
		panic("NewMidpoint called with nil system.Function.")
	}
	r = &Midpoint{stepSize: stepSize, system: system}
	return
}

// These methods are required by the interface Stepper

// Setter methods

// SetStep sets the state for the Midpoint stepper.
func (midpoint *Midpoint) SetStep(step complex128) error {
	midpoint.stepSize = step
	return nil
}

// SetState sets the state for the Midpoint stepper.
func (midpoint *Midpoint) SetState(state []complex128) error {
	midpoint.system.stateVector = state
	return nil
}

// Set sets the step size and System for the Midpoint stepper.
func (midpoint *Midpoint) Set(stepSize complex128, system System) error {
	//if stepSize <= 0 {
	if real(stepSize) <= 0 && imag(stepSize) <= 0.0 {
		return &Error{"Can't use negative or null step size for Midpoint stepper."}
	} else {
		midpoint.stepSize = stepSize
		midpoint.system = system
		return nil
	}
}

// Getter methods

// StepSize returns the step size of the method
func (midpoint *Midpoint) StepSize() complex128 {
	return midpoint.stepSize
}

// State returns the state of the Midpoint stepper.
func (midpoint *Midpoint) State() []complex128 {
	return midpoint.system.stateVector
}

// Step performs one step iteration call of the Midpoint stepper.
// It also updates the state of the Midpoint object.
// The mid-point method for the system,
// 		y = f(t,y)
// Consists of building the sequence,
// 		t_n+1 = t_n + dt
// 		ya_n = y_n + 0.5*dt*f(y_n,t) // First step
//		y_n+1 = y_n + dt*f(ya_n,t)   // Second step
func (midpoint *Midpoint) Step() ([]complex128, error) {

	stateK1 := make([]complex128, len(midpoint.system.stateVector))

	// First Step
	stateK1 = midpoint.system.Evaluate(midpoint.system.stateVector)

	for i := 0; i < len(midpoint.system.stateVector); i++ {
		stateK1[i] = midpoint.system.stateVector[i] + 0.5*midpoint.stepSize*stateK1[i]
	}

	// Second Step
	stateK1 = midpoint.system.Evaluate(stateK1)

	for i := 0; i < len(midpoint.system.stateVector); i++ {
		midpoint.system.stateVector[i] = midpoint.system.stateVector[i] + midpoint.stepSize*stateK1[i]
	}

	return midpoint.system.stateVector, nil
}
