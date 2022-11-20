package BehavioralPatterns

// 工廠

type AbstractFactory interface {
	CreateAnimal() Animal
}

type CatFactory struct {
	AbstractFactory
}

func (fac *CatFactory) CreateAnimal() Animal {
	var animal Animal

	//生產一個具體的貓
	animal = new(Cat)

	return animal
}

type DogFactory struct {
	AbstractFactory
}

func (fac *DogFactory) CreateAnimal() Animal {
	var animal Animal

	//生產一個具體的貓
	animal = new(Dog)

	return animal
}

type SnakeFactory struct {
	AbstractFactory
}

func (fac *SnakeFactory) CreateAnimal() Animal {
	var animal Animal

	//生產一個具體的貓
	animal = new(Snake)

	return animal
}

// 新增需求 黑貓

type BlackCatFactory struct {
	AbstractFactory
}

type BlackCat struct {
	Animal
}

func (blackCat *BlackCat) makeNoise() string {
	return catNoise
}

func (fac *BlackCatFactory) CreateAnimal() Animal {
	var animal Animal

	animal = new(BlackCat)

	return animal
}
