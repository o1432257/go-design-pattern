package BehavioralPatterns

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAbstractFactory_CreateFruit(t *testing.T) {
	testCases := []struct {
		name    string
		factory NewAbstractFactory
		apple   string
		banana  string
		pear    string
	}{
		{
			name:    "TaiwanFactory",
			factory: new(TaiwanFactory),
			apple:   taiwanApple,
			banana:  taiwanBanana,
			pear:    taiwanPear,
		},
		{
			name:    "JapanFactory",
			factory: new(JapanFactory),
			apple:   japanApple,
			banana:  japanBanana,
			pear:    japanPear,
		},
		// 新增美國
		{
			name:    "AmericaFactory",
			factory: new(AmericaFactory),
			apple:   americaApple,
			banana:  americaBanana,
			pear:    americaPear,
		},
	}

	var factory NewAbstractFactory

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			factory = tc.factory

			apple := factory.CreateApple()
			banana := factory.CreateBanana()
			pear := factory.CreatePear()

			require.Equal(t, apple.ShowApple(), tc.apple)
			require.Equal(t, banana.ShowBanana(), tc.banana)
			require.Equal(t, pear.ShowPear(), tc.pear)
		})
	}
}
