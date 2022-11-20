package StructuralPatterns

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProxy(t *testing.T) {
	p1 := Products{
		Kind: "nike運動鞋",
		Fact: false,
	}

	p2 := Products{
		Kind: "手機",
		Fact: true,
	}

	testCases := []struct {
		name     string
		shopping Shopping
	}{
		{
			name:     "AmericaShopping",
			shopping: new(AmericaShopping),
		},
		{
			name:     "JapanShopping",
			shopping: new(JapanShopping),
		},
	}

	var shopping Shopping
	var overseasProxy Shopping

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 自己購物 無法分辨假貨
			shopping = tc.shopping

			product, err := shopping.Buy(&p1)
			require.NoError(t, err)
			require.Equal(t, product, p1.Kind)

			product, err = shopping.Buy(&p2)
			require.NoError(t, err)
			require.Equal(t, product, p2.Kind)

			// 使用代理模式 遇到假貨不買
			overseasProxy = NewProxy(shopping)

			product, err = overseasProxy.Buy(&p1)
			require.Error(t, err, fakeError)
			require.Equal(t, product, "")

			product, err = overseasProxy.Buy(&p2)
			require.NoError(t, err)
			require.Equal(t, product, p2.Kind)
		})
	}

}
