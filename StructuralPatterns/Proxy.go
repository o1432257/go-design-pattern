package StructuralPatterns

import "errors"

var fakeError = errors.New("假貨")

type Products struct {
	Kind string //商品的種類
	Fact bool   //商品的真偽
}

type Shopping interface {
	Buy(products *Products) (string, error)
}

type AmericaShopping struct{}

func (ks AmericaShopping) Buy(products *Products) (string, error) {
	return products.Kind, nil
}

type JapanShopping struct{}

func (ks JapanShopping) Buy(products *Products) (string, error) {
	return products.Kind, nil
}

// OverseasProxy 海外代理
type OverseasProxy struct {
	shopping Shopping
}

func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{shopping}
}

func (op *OverseasProxy) Buy(products *Products) (string, error) {
	//驗貨為真 購買
	if op.distinguish(products) == true {
		// 調用原被代理的具體任務
		return op.shopping.Buy(products)
	}

	return "", fakeError
}

// distinguish 驗貨
func (op *OverseasProxy) distinguish(products *Products) bool {
	return products.Fact
}
