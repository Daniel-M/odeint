// Copyright 2017 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

// System wraps the function that represents the right-hand side of the
// ordinary differential equations system.
// Consider the system
//     x'(t) = f(x,t)
// System wraps around f(x,t).
// params are the parameters that can be passed to the system (i.e. the t in
// f(x,t)), while state is the state where f(x,t) should be evaluated
// (i.e. the x in f(x,t). The System always return a State
//type System func(state []Float, parameters []Float) []Float

type System struct {
	parametersVector []Float
	function         func(state []Float, parameters []Float) []Float
}

func NewSystem(parameters []Float, system func(state []Float, parameters []Float) []Float) (s *System) {
	if system != nil {
		s = &System{parametersVector: parameters, function: system}
	} else {
		panic("NewSystem called with nil system")
	}
	return
}

func (s *System) Evaluate(state []Float) []Float {
	return s.function(state, s.parametersVector)
}

func (s *System) Parameters() []Float {
	return s.parametersVector
}

func (s *System) Function() func(state []Float, parameters []Float) []Float {
	return s.function
}

//type System func(state State) State
