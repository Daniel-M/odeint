# odeint [![Build Status](https://travis-ci.org/Daniel-M/odeint.svg?branch=master)](https://travis-ci.org/Daniel-M/odeint) 

## Description
Go package with ordinary differential equations solvers for initial value problems to be solved by explicit methods.
The package features methods for `float32` `float64` `complex64` and `complex128` types.
The methods implemented so far are,

+ Euler  
+ Mid point  
+ Runge-Kutta 4  

The package is easily extensible to provide other methods.   

## The docs,

The API is pretty consistent, and you'll get it easily,

+ [Getting started](https://godoc.org/github.com/Daniel-M/odeint)   

Type specific docs,

+ [GoDoc documentation for float32](https://godoc.org/github.com/Daniel-M/odeint/float32)   
+ [GoDoc documentation for float64](https://godoc.org/github.com/Daniel-M/odeint/float64)   
+ [GoDoc documentation for complex64](https://godoc.org/github.com/Daniel-M/odeint/complex64)   
+ [GoDoc documentation for complex128](https://godoc.org/github.com/Daniel-M/odeint/complex128)   

## Usage

Import the type you intend to use. For the recommended `float64` just type   

```go
import "github.com/Daniel-M/odeint/float64"   
```
and you are ready to use methods. The stepper used (euler, mid point or Runge-Kutta-4) can be changed easily.   

To get a feeling of the usage check the files under `examples/brusselator`. You will see that to change the method used 
you just need to change the integrator used (line 38 of `examples/brusselator/rk4` source code).   

## FAQ

### A subpackage for each numeric type?

You might be thinking, why does this guy have a subpackage for each numeric type?
Well, though it makes the package harder to maintain, having type specific
integrators is a priority for me. I could have used interface-based integrators
but it would be at the expense of the extensibility of the integrators to more
custom numerical types, a feature which I find relevant too.

## TEST

You can run all tests with   

```sh
go test -v ./...   
```

or individual type oriented tests as   

```sh
go test -v ./float64 # For float64 type   
```


## To Do
+ Support higher order methods.   
+ Implement adaptative step methods.   

