package main

import (
	"time"

	"github.com/caarlos0/spin"
)

func main() {
	show("Default", spin.Default)

	show("Box1", spin.Box1)
	show("Box2", spin.Box2)
	show("Box3", spin.Box3)
	show("Box4", spin.Box4)
	show("Box5", spin.Box5)
	show("Box6", spin.Box6)
	show("Box7", spin.Box7)

	show("Spin1", spin.Spin1)
	show("Spin2", spin.Spin2)
	show("Spin3", spin.Spin3)
	show("Spin4", spin.Spin4)
	show("Spin5", spin.Spin5)
	show("Spin6", spin.Spin6)
	show("Spin7", spin.Spin7)
	show("Spin8", spin.Spin8)
	show("Spin9", spin.Spin9)
	show("Spin10", spin.Spin10)
}

func show(name, frames string) {
	s := spin.New("  \033[36m[" + name + "] computing\033[m %s")
	s.Set(frames)
	s.Start()
	time.Sleep(100 * 20 * time.Millisecond)
	s.Stop()
}
