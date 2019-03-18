package business

import (
	"log"
	"testing"

	"mocking/magic"
)

type MockSpell struct{ magic.Spell }

func (m *MockSpell) Cast(target magic.Target) error {
	log.Println("using mock spell!")
	return nil
}

func TestMoreSerious(t *testing.T) {
	type args struct {
		spell  magic.Caster
		target magic.Target
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "spell tester",
			args: args{
				spell: &MockSpell{
					Spell: magic.Spell{
						Name: "test spell",
						Type: magic.Fire,
					},
				},
				target: magic.Target{
					Name: "Piers Morgan",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MoreSerious(tt.args.spell, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("Serious() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSerious(t *testing.T) {
	type args struct {
		spell  magic.Spell
		target magic.Target
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test spell",
			args: args{
				spell: magic.Spell{
					Name: "Test Spell",
					Type: magic.Dark,
				},
				target: magic.Target{
					Name: "Piers Morgan",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Serious(tt.args.spell, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("Serious() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
