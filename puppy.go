package puppy

import (
	"github.com/madhur/dog"
)

func Bark() string {
	return "I am barking"
}

func Barks() string {
	return "Woof! woof! woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
