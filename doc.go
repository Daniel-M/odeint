// Copyright 2017 Daniel Mejía Raigosa. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
   Package odeint implements Ordinary Differential Equations integrators.
   this documentation is under development, expect updates.

   Float type wrapper

   Float provides a wrapper for floating point precision numbers. the default
   is
       type Float float64
   but can be changed to provide support for other numeric types like
       float32
       complex64
       complex128
   to give support to any valid go numeric type from the standard library:
   https://golang.org/pkg/go/types/#pkg-variables

   License

   This code is licensed under MIT license that can be found in the LICENSE
   file as,

       MIT License

       Copyright (c) 2017 Daniel Mejía Raigosa

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

       Copyright 2017 Daniel Mejía Raigosa. All rights reserved.
       Use of this source code is governed by a MIT
       license that can be found in the LICENSE file.
*/
package odeint
