package BehavioralPatterns

//簡單工廠

const (
	catNoise   = "meow"
	dogNoise   = "bark"
	snakeNoise = "hiss"
)

type Animal interface {
	makeNoise() string
}

type Cat struct {
	Animal
}

func (cat *Cat) makeNoise() string {
	return catNoise
}

type Dog struct {
	Animal
}

func (dog *Dog) makeNoise() string {
	return dogNoise
}

type Snake struct {
	Animal
}

func (snake *Snake) makeNoise() string {
	return snakeNoise
}

type Factory struct{}

func (fac *Factory) CreateAnimal(kind string) Animal {
	var animal Animal

	if kind == "cat" {
		animal = new(Cat)
	} else if kind == "dog" {
		animal = new(Dog)
	} else if kind == "snake" {
		animal = new(Snake)
	}

	return animal
}
