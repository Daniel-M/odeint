// Copyright 2017 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

// Float wraps around float64 or any float representation we want. This is
// introduced to provide a mean to change the floating point precision
// by changing this line
type Float float64

// State implements a wrapper for the state type. State types can be arrays or
// slices of float or float64. This allows the user to change the state type
// by modifying the State struct methods, without the need of changing the
// whole package methods.
type State struct {
	//StateVector []float64
	StateVector []Float
}

// State returns the state array
//func (state *State) State() []float64 {
func (state *State) State() []Float {
	return state.StateVector
}

// State returns the state array
//func (state *State) SetState(s []float64) {
func (state *State) SetState(s []Float) {
	state.StateVector = s
}
