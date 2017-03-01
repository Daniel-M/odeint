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
//type System func(params State, state State) State
type System func(state State) State
