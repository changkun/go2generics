// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errors offers try/catch/final primitives for error handling
package errors

// Try tries a given function and see if it throw any error
func Try[T any](e func() (T, error)) catch[T] {
	v, err := e()
	return catch[T]{err: err, result: v}
}

type catch[T any] struct {
	err    error
	result T
}

// Catch catches the error that threw in the Try call.
func (c catch[T]) Catch(handler func(T, error) T) catch[T] {
	if c.err != nil {
		c.result = handler(c.result, c.err)
	}
	c.err = nil
	return c
}

// Final runs in any case if it is called.
func (c catch[T]) Final(handler func(T)) {
	handler(c.result)
}