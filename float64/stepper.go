// Copyright 2017-2018 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

// Stepper defines the functions that any Ordinary Differential Equation
// Integrator Stepper should implement.
type Stepper interface {
	// Setter methods
	SetStep(step float64) error
	SetState(state []float64) error
	Set(stepSize float64, system System) error

	// Getter methods
	StepSize() float64
	State() []float64
	Step() ([]float64, error)
}
