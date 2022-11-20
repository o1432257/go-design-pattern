package BehavioralPatterns

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimpleFactory_CreateAnimal(t *testing.T) {
	testCases := []struct {
		name   string
		animal string
		noise  string
	}{
		{
			name:   "createCat",
			animal: "cat",
			noise:  catNoise,
		},
		{
			name:   "createDog",
			animal: "dog",
			noise:  dogNoise,
		},
		{
			name:   "createSnake",
			animal: "snake",
			noise:  snakeNoise,
		},
	}

	factory := new(Factory)
	var animal Animal

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			animal = factory.CreateAnimal(tc.animal)
			require.Equal(t, animal.makeNoise(), tc.noise)
		})
	}
}
