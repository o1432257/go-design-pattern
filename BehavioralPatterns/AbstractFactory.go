package BehavioralPatterns

// 抽象工廠

const (
	taiwanApple   = "台灣蘋果"
	taiwanBanana  = "台灣香蕉"
	taiwanPear    = "台灣梨"
	japanApple    = "日本蘋果"
	japanBanana   = "日本香蕉"
	japanPear     = "日本梨"
	americaApple  = "美國蘋果"
	americaBanana = "美國香蕉"
	americaPear   = "美國梨"
)

// 抽象層

type AbstractApple interface {
	ShowApple() string
}

type AbstractBanana interface {
	ShowBanana() string
}

type AbstractPear interface {
	ShowPear() string
}

type AbstractPeach interface {
}

type NewAbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

// 實現層

type TaiwanApple struct{}

func (ta *TaiwanApple) ShowApple() string {
	return taiwanApple
}

type TaiwanBanana struct{}

func (tb *TaiwanBanana) ShowBanana() string {
	return taiwanBanana
}

type TaiwanPear struct{}

func (tp *TaiwanPear) ShowPear() string {
	return taiwanPear
}

type TaiwanFactory struct{}

func (tf *TaiwanFactory) CreateApple() AbstractApple {
	var apple AbstractApple

	apple = new(TaiwanApple)

	return apple
}

func (tf *TaiwanFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana

	banana = new(TaiwanBanana)

	return banana
}

func (tf *TaiwanFactory) CreatePear() AbstractPear {
	var pear AbstractPear

	pear = new(TaiwanPear)

	return pear
}

type JapanApple struct{}

func (ta *JapanApple) ShowApple() string {
	return japanApple
}

type JapanBanana struct{}

func (tb *JapanBanana) ShowBanana() string {
	return japanBanana
}

type JapanPear struct{}

func (tp *JapanPear) ShowPear() string {
	return japanPear
}

type JapanFactory struct{}

func (tf *JapanFactory) CreateApple() AbstractApple {
	var apple AbstractApple

	apple = new(JapanApple)

	return apple
}

func (tf *JapanFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana

	banana = new(JapanBanana)

	return banana
}

func (tf *JapanFactory) CreatePear() AbstractPear {
	var pear AbstractPear

	pear = new(JapanPear)

	return pear
}

// 新增美國

type AmericaApple struct{}

func (ta *AmericaApple) ShowApple() string {
	return americaApple
}

type AmericaBanana struct{}

func (tb *AmericaBanana) ShowBanana() string {
	return americaBanana
}

type AmericaPear struct{}

func (tp *AmericaPear) ShowPear() string {
	return americaPear
}

type AmericaFactory struct{}

func (tf *AmericaFactory) CreateApple() AbstractApple {
	var apple AbstractApple

	apple = new(AmericaApple)

	return apple
}

func (tf *AmericaFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana

	banana = new(AmericaBanana)

	return banana
}

func (tf *AmericaFactory) CreatePear() AbstractPear {
	var pear AbstractPear

	pear = new(AmericaPear)

	return pear
}
