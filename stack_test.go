package main

import (
	"math/rand"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

/*
INITIALIZING STACK TESTS
*/
func TestNewStack_GivenWithoutParameter_ThenReturnsEmptyStack(t *testing.T) {
	// Arrange
	expectedCapacity := 0

	// Act
	sut := NewStack[int]()

	// Assert
	assert.Len(t, sut.Elements, expectedCapacity)
}

func TestNewStack_GivenWithCapacityParameter_ThenReturnsEmptyStackWithCapacity(t *testing.T) {
	// Arrange
	capacity := rand.Intn(10)

	// Act
	sut := NewStackWithCapacity[int](capacity)

	// Assert
	assert.Len(t, sut.Elements, capacity)
}

func TestNewStack_GivenWithElementsParameter_ThenReturnsStackWithElements(t *testing.T) {
	// Arrange
	elements, _ := faker.RandomInt(0, 10)

	// Act
	sut := NewStackWithElements(elements)

	// Assert
	assert.ElementsMatch(t, sut.Elements, elements)
}

/*
INITIALIZING STACK TESTS
*/
func TestClear_GivenStackContainElements_ThenCleanElementsAndReturnTrue(t *testing.T) {
	// Arrange
	elements, _ := faker.RandomInt(0, 10)
	sut := NewStackWithElements(elements)

	// Act
	sut.Clear()

	// Assert
	assert.Len(t, sut.Elements, 0)
}

func TestContains_GivenElementNotFoundInStackElements_ThenReturnFalse(t *testing.T) {
	// Arrange
	elements, _ := faker.RandomInt(11, 20)
	itemToFind := rand.Intn(10)
	sut := NewStackWithElements(elements)

	// Act
	result := sut.Contains(itemToFind)

	// Assert
	assert.False(t, result)
}

func TestContains_GivenElementFoundInStackElements_ThenReturnTrue(t *testing.T) {
	// Arrange
	elements, _ := faker.RandomInt(0, 10)
	itemToFind := elements[0]
	sut := NewStackWithElements(elements)

	// Act
	result := sut.Contains(itemToFind)

	// Assert
	assert.True(t, result)
}

func TestCount_GivenStackWithElements_ThenReturnsTheCount(t *testing.T) {
	// Arrange
	elements, _ := faker.RandomInt(0, 10)
	sut := NewStackWithElements(elements)

	// Act
	result := sut.Count()

	// Assert
	assert.Equal(t, len(elements), result)
}

func TestPeek_GivenStackIsEmpty_ThenReturnsErrorWithEmptyStackMessage(t *testing.T) {
	// Arrange
	expectedErrorMsg := "err: stack not contains elements"
	sut := NewStack[int]()

	// Act
	result, err := sut.Peek()

	// Assert
	assert.Zero(t, result)
	assert.EqualError(t, err, expectedErrorMsg)
}

func TestPeek_GivenStackWithElements_ThenReturnsTheTopElementOfTheStackAndKeepYourPosition(t *testing.T) {
	// Arrange
	elements, _ := faker.RandomInt(0, 10)
	expectedTopElement := elements[len(elements)-1]
	sut := NewStackWithElements(elements)

	// Act
	result, err := sut.Peek()

	// Assert
	assert.Equal(t, expectedTopElement, result)
	assert.NoError(t, err)
}

func TestPop_GivenStackIsEmpty_ThenReturnsErrorWithEmptyStackMessage(t *testing.T) {
	// Arrange
	expectedErrorMsg := "err: stack not contains elements"
	sut := NewStack[int]()

	// Act
	err := sut.Pop()

	// Assert
	assert.EqualError(t, err, expectedErrorMsg)
}

func TestPop_GivenStackWithElements_ThenRemoveTheTopElementOfTheStack(t *testing.T) {
	// Arrange
	elements, _ := faker.RandomInt(0, 10)
	expectedCount := len(elements) - 1
	expectedTopElement := elements[len(elements)-2]
	sut := NewStackWithElements(elements)

	// Act
	err := sut.Pop()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, sut.Elements, expectedCount)
	assert.Equal(t, expectedTopElement, sut.Elements[len(sut.Elements)-1])
}

func TestPush_GivenNewElementAndEmptyStack_ThenInsertTheElementToTheFirstOfTheStack(t *testing.T) {
	// Arrange
	newElement := rand.Intn(10)
	sut := NewStack[int]()

	// Act
	err := sut.Push(newElement)

	// Assert
	assert.NoError(t, err)
	assert.Len(t, sut.Elements, 1)
	assert.Equal(t, newElement, sut.Elements[0])
}

func TestPush_GivenNewElementAndStackWithElements_ThenInsertTheElementToTheTopOfTheStack(t *testing.T) {
	// Arrange
	elements, _ := faker.RandomInt()
	newElement := rand.Intn(10)
	expectedLength := len(elements) + 1
	sut := NewStackWithElements(elements)

	// Act
	err := sut.Push(newElement)

	// Assert
	assert.NoError(t, err)
	assert.Len(t, sut.Elements, expectedLength)
	assert.Equal(t, newElement, sut.Elements[len(sut.Elements)-1])
}
