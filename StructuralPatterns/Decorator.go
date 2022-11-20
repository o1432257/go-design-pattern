package StructuralPatterns

const (
	tShirt = "T 恤"
	pants  = "長褲"
	white  = "白色的"
	yellow = "黃色的"
)

type Clothes interface {
	Show() string
}

type Decorator struct {
	clothes Clothes
}

func (d *Decorator) Show() {}

type TShirt struct{}

func (ts TShirt) Show() string {
	return tShirt
}

type Pants struct{}

func (pt Pants) Show() string {
	return pants
}

type WhiteDecorator struct {
	Decorator
}

func NewWhiteDecorator(clothes Clothes) Clothes {
	return &WhiteDecorator{Decorator{clothes}}
}

func (w WhiteDecorator) Show() string {
	return white + w.clothes.Show()
}

type YellowDecorator struct {
	Decorator
}

func NewYellowDecorator(clothes Clothes) Clothes {
	return &YellowDecorator{Decorator{clothes}}
}

func (y YellowDecorator) Show() string {
	return yellow + y.clothes.Show()
}
