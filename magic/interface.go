package magic

// Caster defines an interface for casting spells
type Caster interface {
	Cast(target Target) error
}
