// Copyright 2017-2018 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

//import "fmt"

// STEPPER_METHOD implements the stepper_method stepper method. stepper_method is part of the Stepper
// interface.
type STEPPER_METHOD struct {
	stepSize Float //Float is used as the template name for the numeric type
	system   System
}

// NewSTEPPER_METHOD returns a reference to a new stepper_method stepper method.
func NewSTEPPER_METHOD(stepSize Float, system System) (r *STEPPER_METHOD) {
	if stepSize <= 0 {
		panic("NewSTEPPER_METHOD called with negative or null stepSize.")
	}
	if system.Function() == nil {
		panic("NewSTEPPER_METHOD called with nil system.Function.")
	}
	r = &STEPPER_METHOD{stepSize: stepSize, system: system}
	return
}

// These methods are required by the interface Stepper

// Setter methods

// SetStep sets the state for the STEPPER_METHOD stepper.
func (stepper_method *STEPPER_METHOD) SetStep(step Float) error {
	stepper_method.stepSize = step
	return nil
}

// SetState sets the state for the STEPPER_METHOD stepper.
func (stepper_method *STEPPER_METHOD) SetState(state []Float) error {
	stepper_method.system.stateVector = state
	return nil
}

// Set sets the step size and System for the STEPPER_METHOD stepper.
func (stepper_method *STEPPER_METHOD) Set(stepSize Float, system System) error {
	if stepSize <= 0 {
		return &Error{"Can't use negative or null step size for STEPPER_METHOD stepper."}
	} else {
		stepper_method.stepSize = stepSize
		stepper_method.system = system
		return nil
	}
}

// Getter methods

// StepSize returns the step size of the method
func (stepper_method *STEPPER_METHOD) StepSize() Float {
	return stepper_method.stepSize
}

// State returns the state of the STEPPER_METHOD stepper.
func (stepper_method *STEPPER_METHOD) State() []Float {
	return stepper_method.system.stateVector
}

// Step performs one step iteration call of the STEPPER_METHOD stepper.
// It also updates the state of the STEPPER_METHOD object.
func (stepper_method *STEPPER_METHOD) Step() ([]Float, error) {

	newstate := stepper_method.system.stateVector
	stateK1 := make([]Float, len(stepper_method.system.stateVector))

	BufferK1 := make([]Float, len(stepper_method.system.stateVector))

	// Here you perform one step iteration of stepper_method
	// Result
	for i := 0; i < len(stepper_method.system.stateVector); i++ {
		// Fill new state. Runge Kutta for reference
		newstate[i] = newstate[i] + (stepper_method.stepSize/6.0)*(stateK1[i]+2*stateK2[i]+2*stateK3[i]+stateK4[i])
	}

	stepper_method.system.stateVector = newstate

	return newstate, nil
}
