package zoo

import (
	animal "zoo/animal"
)

type Cage struct {
	Name string
}

type RabbitCage struct {
	Cage
	Animal *animal.Rabbit
}

type LizardCage struct {
	Cage
	Animal *animal.Lizard
}

type ChickenCage struct {
	Cage
	Animal *animal.Chicken
}

func (rc *RabbitCage) PutRabbit(a *animal.Rabbit) {
	rc.Animal = a
}

func (lc LizardCage) PutLizard(a *animal.Lizard) {
	lc.Animal = a
}

func (cc ChickenCage) PutChicken(a *animal.Chicken) {
	cc.Animal = a
}

func NewRabbitCage(name string) RabbitCage {
	return RabbitCage{Cage{Name: name}, nil}
}

func NewLizardCage(name string) LizardCage {
	return LizardCage{Cage{Name: name}, nil}
}

func NewChickenCage(name string) ChickenCage {
	return ChickenCage{Cage{Name: name}, nil}
}
