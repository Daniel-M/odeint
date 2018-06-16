// Copyright 2017-2018 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

// System wraps the function that represents the right-hand side of the
// ordinary differential equations system, its parameters t, and state x
// Consider the system,
//     x'(t) = f(x,t)  (1)
// System wraps around f(x,t) storing also t and x.
// stateVector is the present state of the system, i.e. the values stored
// at the components of x.
// parametersVector are the parameters of to the system (i.e. the t in f(x,t))
// function describes the func(state []float64, parameters []float64) []float64
// that represents the right hand side of system (1)
type System struct {
	stateVector      []float64
	parametersVector []float64
	function         func(state []float64, parameters []float64) []float64
}

// NewSystem returns a reference to a new System object with the properties given
func NewSystem(state []float64, parameters []float64, system func(state []float64, parameters []float64) []float64) (s *System) {
	if system != nil {
		s = &System{stateVector: state, parametersVector: parameters, function: system}
	} else {
		panic("NewSystem called with nil system")
	}
	return
}

// Evaluate returns the result of evaluating f(x,t) with x = state.
// if the size of state is zero, it returns f(x,t) using the internal state x
func (s *System) Evaluate(state []float64) []float64 {
	if len(state) > 0 {
		return s.function(state, s.parametersVector)
	} else {
		return s.function(s.stateVector, s.parametersVector)
	}
}

// State returns the internal state vector of the System
func (s *System) State() []float64 {
	return s.stateVector
}

// Parameters returns the internal parameters vector of the System
func (s *System) Parameters() []float64 {
	return s.parametersVector
}

// State returns the internal function of the System
func (s *System) Function() func(state []float64, parameters []float64) []float64 {
	return s.function
}
