// Copyright 2017-2018 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package odeint

//import "errors"
import "fmt"

// Error implements an error custom type.
type Error struct {
	errorString string
}

// Error implementation of the Error() interface for the package "error".
func (e *Error) Error() string {
	return fmt.Sprintf("ERROR: %s", e.errorString)
}

//func New(text string) error {
//return &Error{text}
//}
