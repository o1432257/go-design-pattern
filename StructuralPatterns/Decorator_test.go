package StructuralPatterns

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDecorator(t *testing.T) {
	testCases := []struct {
		name    string
		clothes Clothes
		show    string
	}{
		{
			name:    "TShirt",
			clothes: new(TShirt),
			show:    tShirt,
		},
		{
			name:    "Pants",
			clothes: new(Pants),
			show:    pants,
		},
	}

	var clothes Clothes
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			clothes = tc.clothes

			require.Equal(t, clothes.Show(), tc.show)

			clothes = NewWhiteDecorator(clothes)

			require.Equal(t, clothes.Show(), white+tc.show)

			clothes = NewYellowDecorator(clothes)

			require.Equal(t, clothes.Show(), yellow+white+tc.show)
		})
	}
}