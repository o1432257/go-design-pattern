package BehavioralPatterns

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	prev *singleton
	next *singleton
)

func TestSingleton(t *testing.T) {
	s1 := GetInstance()
	s2 := GetInstance()

	require.Equal(t, s1, s2)
}

// TestSingletonWithHighConcurrency 高併發
func TestSingletonWithHighConcurrency(t *testing.T) {
	n := 100
	singleton := make(chan *singleton, n)

	for i := 0; i < n; i++ {
		go func() {
			tmp := GetInstance()
			singleton <- tmp
		}()
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			next = <-singleton
			continue
		}

		prev = next
		next = <-singleton

		require.Equal(t, prev, next)
	}
}
