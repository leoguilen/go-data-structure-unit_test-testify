package main

import (
	"errors"
	"reflect"
)

const (
	InitialCapacity = 0
)

var (
	ErrEmptyStack = errors.New("err: stack not contains elements")
)

type Stack[T any] struct {
	Elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		Elements: make([]T, InitialCapacity),
	}
}

func NewStackWithCapacity[T any](c int) *Stack[T] {
	return &Stack[T]{
		Elements: make([]T, c),
	}
}

func NewStackWithElements[T any](e []T) *Stack[T] {
	return &Stack[T]{
		Elements: e,
	}
}

func (s *Stack[T]) Clear() {
	s.Elements = nil
}

func (s *Stack[T]) Contains(t T) bool {
	for _, el := range s.Elements {
		if reflect.ValueOf(el).Interface() == reflect.ValueOf(t).Interface() {
			return true
		}
	}

	return false
}

func (s *Stack[T]) Count() int {
	return len(s.Elements)
}

func (s *Stack[T]) Peek() (T, error) {
	if len(s.Elements) == 0 {
		return *new(T), ErrEmptyStack
	}

	return s.Elements[len(s.Elements)-1], nil
}

func (s *Stack[T]) Pop() error {
	if len(s.Elements) == 0 {
		return ErrEmptyStack
	}

	s.Elements = s.Elements[:len(s.Elements)-1]
	return nil
}

func (s *Stack[T]) Push(t T) error {
	s.Elements = append(s.Elements, t)
	return nil
}
