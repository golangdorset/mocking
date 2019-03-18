package business

import (
	"fmt"
	"time"

	"mocking/magic"
)

// Serious businessing takes place within this function.
func Serious(spell magic.Spell, target magic.Target) error {
	things := []string{
		"very serious business:",
		"computing money...",
		"activating almonds...",
		"performing magic spells...",
	}
	for _, thing := range things {
		fmt.Println(thing)
		time.Sleep(250 * time.Millisecond)
	}
	if err := spell.Cast(target); err != nil {
		return err
	}
	fmt.Println("business complete.")
	return nil
}

// MoreSerious businessing takes place within this function.
func MoreSerious(spell magic.Caster, target magic.Target) error {
	things := []string{
		"very serious business:",
		"computing money...",
		"activating almonds...",
		"performing magic spells...",
	}
	for _, thing := range things {
		fmt.Println(thing)
		time.Sleep(250 * time.Millisecond)
	}
	if err := spell.Cast(target); err != nil {
		return err
	}
	fmt.Println("business complete.")
	return nil
}
