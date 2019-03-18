package magic

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//go:generate stringer -type SpellType
type SpellType int

const (
	Dark SpellType = iota
	Light
	Fire
	Water
	Air
	Earth
)

// Spell to construct
type Spell struct {
	Name string    `json:"name"`
	Type SpellType `json:"type"`
}

// Target a poor sod
type Target struct{ Name string `json:"name"` }

var errSpellFailed = errors.New("the spell failed or was deflected by the target")

// Cast the spell
func (s *Spell) Cast(target Target) error {
	rand.Seed(time.Now().UnixNano())

	// this is problematic... our tests could fail!
	if rand.Intn(3) > 1 {
		fmt.Printf("zap! %s took the brunt of your spell!\n", target.Name)
		return nil
	}
	return errSpellFailed
}
