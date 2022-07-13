package stacks_test

import (
	"testing"

	"github.com/twharmon/stacks"
)

func BenchmarkPush(b *testing.B) {
	b.Run("push", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := stacks.New[int]()
			for j := 0; j < 1000; j++ {
				s.Push(j)
			}
		}
	})
	b.Run("push pop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := stacks.New[int]()
			for j := 0; j < 1000; j++ {
				s.Push(j)
				if j%100 == 0 {
					for k := 0; k < 100; k++ {
						s.Pop()
					}
				}
			}
		}
	})
}
