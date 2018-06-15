// Copyright 2017 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

// Stepper defines the functions that any Ordinary Differential Equation
// Integrator Stepper should implement.
type Stepper interface {
	// Setter methods
	SetStep(step float32) error
	SetState(state []float32) error
	Set(stepSize float32, system System) error

	// Getter methods
	StepSize() float32
	State() []float32
	Step() ([]float32, error)
}
