package zoo

import (
	"fmt"
	animals "zoo/animal"
	cage "zoo/cage"
)

type Keeper struct {
	Name string
}

func (k *Keeper) CatchRabbit(a *animals.Rabbit, c *cage.RabbitCage) {
	if c.Animal != nil {
		fmt.Println("The cage is not empty!")
		return
	}
	c.PutRabbit(a)
}

func (k *Keeper) CatchLizard(a *animals.Lizard, c *cage.LizardCage) {
	if c.Animal != nil {
		fmt.Println("The cage is not empty!")
		return
	}
	c.PutLizard(a)
}

func (k *Keeper) CatchChicken(a *animals.Chicken, c *cage.ChickenCage) {
	if c.Animal != nil {
		fmt.Println("The cage is not empty!")
		return
	}
	c.PutChicken(a)
}

func NewZooKeeper(name string) Keeper {
	return Keeper{
		Name: name,
	}
}
