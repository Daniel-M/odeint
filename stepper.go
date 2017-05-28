// Copyright 2017 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

// Stepper defines the functions that any Ordinary Differential Equation
// Integrator Stepper should implement.
type Stepper interface {
	// Setter methods
	SetStep(step Float) error
	SetState(state []Float) error
	Set(stepSize Float, system System) error

	// Getter methods
	StepSize() Float
	State() []Float
	Step() ([]Float, error)
}
