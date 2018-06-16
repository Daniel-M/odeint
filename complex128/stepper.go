// Copyright 2017-2018 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

// Stepper defines the functions that any Ordinary Differential Equation
// Integrator Stepper should implement.
type Stepper interface {
	// Setter methods
	SetStep(step complex128) error
	SetState(state []complex128) error
	Set(stepSize complex128, system System) error

	// Getter methods
	StepSize() complex128
	State() []complex128
	Step() ([]complex128, error)
}
