package stacks_test

import (
	"reflect"
	"testing"

	"github.com/twharmon/stacks"
)

func assertEqual(t *testing.T, want, got interface{}) {
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestPop(t *testing.T) {
	s := stacks.New[string]()
	s.Push("foo")
	s.Push("bar")
	assertEqual(t, "bar", s.Pop())
	assertEqual(t, "foo", s.Pop())
	assertEqual(t, "", s.Pop())
}

func TestPeek(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		s := stacks.New[string]()
		s.Push("foo")
		s.Push("bar")
		assertEqual(t, "bar", s.Peek())
		assertEqual(t, "bar", s.Peek())
		assertEqual(t, 2, s.Len())
	})
	t.Run("empty", func(t *testing.T) {
		s := stacks.New[string]()
		assertEqual(t, "", s.Peek())
	})
}

func TestLen(t *testing.T) {
	s := stacks.New[string]()
	s.Push("foo")
	s.Push("bar")
	assertEqual(t, 2, s.Len())
	s.Pop()
	assertEqual(t, 1, s.Len())
	s.Pop()
	assertEqual(t, 0, s.Len())
}

func TestOverwrite(t *testing.T) {
	s := stacks.New[string]()
	s.Push("foo")
	s.Push("bar")
	s.Push("baz")
	s.Pop()
	s.Pop()
	s.Pop()
	s.Push("foo")
	s.Push("bar")
	s.Push("baz")
	assertEqual(t, "baz", s.Pop())
}

func TestSlice(t *testing.T) {
	s := stacks.New[string]()
	s.Push("foo")
	s.Push("bar")
	s.Push("baz")
	assertEqual(t, []string{"foo", "bar", "baz"}, s.Slice())
}

func TestClear(t *testing.T) {
	s := stacks.New[string]()
	s.Push("foo")
	s.Push("bar")
	s.Push("baz")
	s.Clear()
	assertEqual(t, 0, s.Len())
}

func TestString(t *testing.T) {
	s := stacks.New[string]()
	s.Push("foo")
	s.Push("bar")
	s.Push("baz")
	s.Pop()
	assertEqual(t, "[foo bar]", s.String())
}
