package main

import (
	"log"

	"mocking/business"
	"mocking/magic"
)

func main() {
	spell := magic.Spell{
		Name: "Spontaneous Combustion",
		Type: magic.Fire,
	}
	target := magic.Target{Name: "Piers Morgan"}

	if err := business.Serious(spell, target); err != nil {
		log.Panicln(err)
	}
}
