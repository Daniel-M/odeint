// Copyright 2017-2018 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
  Package odeint implements Ordinary Differential Equations integrators.
  for initial value problems to be solved by explicit methods.
  The package features methods for the standard library types,

    float32
    float64
    complex64
    complex128


  Methods for float64

  The import to use float64 is

    import "github.com/Daniel-M/odeint/float64"

  The integrator methods implemented so far are,

  * Euler

  * Mid point

  * Runge-Kutta 4


  The package is easily extensible to provide other methods, you can follow
  the template files as reference,

    templates/stepper_method.go.t
    templates/stepper_method_test.go.t

  Example - Simple harmonic oscillator

  The integrator can be used to integrate the ODE for the harmonic oscillator.

    x'' + p*x' + k*x = 0


  which can be decomposed as the system,

    x' = u
    u' = -p*u - k*x

  If we want to solve the system using float64, we must import the adequate
  subpackage,

    import "github.com/Daniel-M/odeint/float64"

  So we begin by defining the system of coupled differential equations

    func odesys(x []float64, parameters []float64) []float64 {

      dxdt := make([]float64, len(x))

      dxdt[0] = x[1]
      dxdt[1] = -parameters[0]*x[0] - parameters[1]*x[1]

      return dxdt
    }

  declare the state and parameters variables,
    state := make([]float64, 2)
    params := make([]float64, 2)

  Putting the inital conditions
    state[0] = 0.2
    state[1] = 0.8

  And the parameters
    params[0] = 1.2 * 1.2
    params[1] = 0.2

  We create an instance of the system,
    system := odeint.NewSystem(state, params, odesys)

  And an instance of the integrator with Midpoint method,
    var integrator odeint.Midpoint

  Set the system to the integrator before integrating the system

    err := integrator.Set(0.1, *system)
    if err != nil {
      panic(err)
    }

  And finally we integrate within a loop

    for i := 0; i < int(30.0/integrator.StepSize()); i++ {
      fmt.Println(float64(i)*integrator.StepSize(), state)

      state, err = integrator.Step()
      if err != nil {
        panic(err)
      }
    }

  The code above will print the data columns to the standard output.
  To write to a file you could create a file with
    os.Create
  and write to it with
    fmt.Fprintf(w,...)
  where w implements the interface
    io.Writter
  There are more examples at the examples path


  FAQ



  A subpackage for each numeric type?

  You might be thinking, why does this guy have a subpackage for each numeric type?
  Well, though it makes the package harder to maintain, having type specific
  integrators is a priority for me. I could have used interface-based integrators
  but it would be at the expense of the extensibility of the integrators to more
  custom numerical types, a feature which I find relevant too.

  License

  This code is licensed under MIT license that can be found in the LICENSE
  file as,

  		MIT License

  		Copyright (c) 2017-2018 Daniel Mejía Raigosa

  		Permission is hereby granted, free of charge, to any person obtaining a copy
  		of this software and associated documentation files (the "Software"), to deal
  		in the Software without restriction, including without limitation the rights
  		to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
  		copies of the Software, and to permit persons to whom the Software is
  		furnished to do so, subject to the following conditions:

  		The above copyright notice and this permission notice shall be included in all
  		copies or substantial portions of the Software.

  		THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
  		IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
  		FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
  		AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
  		LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
  		OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
  		SOFTWARE.

  And at the begining of the source code files as the fragment,

  		Copyright 2017-2018 Daniel Mejía Raigosa. All rights reserved.
  		Use of this source code is governed by a MIT
  		license that can be found in the LICENSE file.
*/
package odeint
