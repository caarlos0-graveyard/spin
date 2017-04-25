package spin_test

import (
	"testing"
	"time"

	"github.com/caarlos0/spin"
)

var spins = map[string]string{
	"Box1":    spin.Box1,
	"Box2":    spin.Box2,
	"Box3":    spin.Box3,
	"Box4":    spin.Box4,
	"Box5":    spin.Box5,
	"Box6":    spin.Box6,
	"Box7":    spin.Box7,
	"Spin1":   spin.Spin1,
	"Spin2":   spin.Spin2,
	"Spin3":   spin.Spin3,
	"Spin4":   spin.Spin4,
	"Spin5":   spin.Spin5,
	"Spin6":   spin.Spin6,
	"Spin7":   spin.Spin7,
	"Spin8":   spin.Spin8,
	"Spin9":   spin.Spin9,
	"Spin10":  spin.Spin10,
	"Default": spin.Default,
}

func TestSpin(t *testing.T) {
	for name, spinner := range spins {
		t.Run(name, func(t *testing.T) {
			show(name, spinner)
		})
	}
}

func show(name, frames string) {
	s := spin.New("  \033[36m[" + name + "] Testing\033[m %s")
	s.Set(frames)
	s.Start()
	defer s.Stop()
	time.Sleep(2 * time.Second)
}
