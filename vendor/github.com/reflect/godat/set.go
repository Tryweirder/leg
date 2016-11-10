// Portions of this file are derived from GoDS, a data structure library for
// Go.
//
// Copyright (c) 2015, Emir Pasic. All rights reserved.
//
// https://github.com/emirpasic/gods/blob/213367f1ca932600ce530ae11c8a8cc444e3a6da/sets/sets.go

package godat

type SetIterationFunc func(element interface{}) error

// A well-defined collection of distinct objects.
type Set interface {
	Container

	// Returns true if the set contains all of the elements specified, and
	// false otherwise.
	Contains(elements ...interface{}) bool

	// Adds all of the elements specified to the set. If an element already
	// exists in the set, it will not be duplicated.
	Add(elements ...interface{})

	// Removes all of the elements specified from the set.
	Remove(elements ...interface{})

	// Iterates each element in the set and executes the given callback
	// function. If the callback function returns an error, this function will
	// return the same error and immediately stop iteration.
	//
	// To stop iteration without returning an error, return ErrStopIteration.
	ForEach(fn SetIterationFunc) error

	// Iterates each element in the set and executes the given callback
	// function, which must be of a type similar to SetIterationFunc, except
	// that the element parameter may be any type assignable by every element
	// in the set.
	//
	// If the requirements for the fn parameter are not met, this function will
	// panic.
	ForEachInto(fn interface{}) error
}
