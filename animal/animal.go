package zoo

import "fmt"

type Animal struct {
	name   string
	weight string
	age    int
	animalClass
}

type animalClass string

const (
	reptile animalClass = "reptile"
	mammal  animalClass = "mammal"
	bird    animalClass = "bird"
)

type Rabbit struct {
	Animal
}

type Lizard struct {
	Animal
}

type Chicken struct {
	Animal
}

func (a Animal) String() string {
	return fmt.Sprintf("Name: %s\nWeight: %s\n Age: %d\nClass: %s", a.name, a.weight, a.age, a.animalClass)
}

func NewRabbit(name string, weight string, age int) Rabbit {
	return Rabbit{
		Animal{
			name:        name,
			weight:      weight,
			age:         age,
			animalClass: mammal,
		},
	}
}

func NewLizard(name string, weight string, age int) Lizard {
	return Lizard{
		Animal{
			name:        name,
			weight:      weight,
			age:         age,
			animalClass: reptile,
		},
	}
}

func NewChicken(name string, weight string, age int) Chicken {
	return Chicken{
		Animal{
			name:        name,
			weight:      weight,
			age:         age,
			animalClass: bird,
		},
	}
}
