package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	var s1 Set[string] = NewSet[string]()
	s1.add("Apple")
	s1.add("Banana")
	s1.add("Apple")
	s1.add("Mango")
	s1.add("Orange")

	assert.Equal(t, true, s1.contains("Orange"))
	assert.Equal(t, true, s1.contains("Mango"))
	assert.Equal(t, true, s1.contains("Apple"))
	assert.Equal(t, true, s1.contains("Banana"))
}

func TestRemove(t *testing.T) {
	var s = NewSet[string]()
	s.add("Apple")
	s.add("Banana")
	s.add("Apple")
	s.add("Mango")
	s.add("Orange")
	assert.Equal(t, true, s.contains("Apple"))
	s.remove("Apple")
	assert.Equal(t, false, s.contains("Apple"))
}
