// Copyright 2017-2018 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

// System wraps the function that represents the right-hand side of the
// ordinary differential equations system, its parameters t, and state x
// Consider the system,
//     x'(t) = f(t,x)  (1)
// System wraps around f(t,x) storing also t and x.
// stateVector is the present state of the system, i.e. the values stored
// at the components of x.
// parametersVector are the parameters of to the system (i.e. the t in f(t,x))
// function describes the func(state []float32, parameters []float32) []float32
// that represents the right hand side of system (1)
type System struct {
	stateVector      []float32
	parametersVector []float32
	function         func(state []float32, parameters []float32) []float32
}

// NewSystem returns a reference to a new System object with the properties given
func NewSystem(state []float32, parameters []float32, system func(state []float32, parameters []float32) []float32) (s *System) {
	if system != nil {
		s = &System{stateVector: state, parametersVector: parameters, function: system}
	} else {
		panic("NewSystem called with nil system")
	}
	return
}

// Evaluate returns the result of evaluating f(t,x) with x = state.
// if the size of state is zero, it returns f(t,x) using the internal state x
func (s *System) Evaluate(state []float32) []float32 {
	if len(state) > 0 {
		return s.function(state, s.parametersVector)
	} else {
		return s.function(s.stateVector, s.parametersVector)
	}
}

// State returns the internal state vector of the System
func (s *System) State() []float32 {
	return s.stateVector
}

// Parameters returns the internal parameters vector of the System
func (s *System) Parameters() []float32 {
	return s.parametersVector
}

// State returns the internal function of the System
func (s *System) Function() func(state []float32, parameters []float32) []float32 {
	return s.function
}
