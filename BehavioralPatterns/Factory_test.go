package BehavioralPatterns

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFactory_CreateAnimal(t *testing.T) {
	testCases := []struct {
		name    string
		animal  string
		factory AbstractFactory
		noise   string
	}{
		{
			name:    "createCat",
			factory: new(CatFactory),
			noise:   catNoise,
		},
		{
			name:    "createDog",
			factory: new(DogFactory),
			noise:   dogNoise,
		},
		{
			name:    "createSnake",
			factory: new(SnakeFactory),
			noise:   snakeNoise,
		},
		{
			name:    "createBlackCat",
			factory: new(BlackCatFactory),
			noise:   catNoise,
		},
	}

	var factory AbstractFactory
	var animal Animal

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			factory = tc.factory
			animal = factory.CreateAnimal()
			require.Equal(t, animal.makeNoise(), tc.noise)
		})
	}
}
