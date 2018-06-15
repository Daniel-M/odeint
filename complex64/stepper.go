// Copyright 2017 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

// Stepper defines the functions that any Ordinary Differential Equation
// Integrator Stepper should implement.
type Stepper interface {
	// Setter methods
	SetStep(step complex64) error
	SetState(state []complex64) error
	Set(stepSize complex64, system System) error

	// Getter methods
	StepSize() complex64
	State() []complex64
	Step() ([]complex64, error)
}
