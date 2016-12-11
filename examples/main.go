package main

import (
	"fmt"
	"time"

	spin "github.com/caarlos0/spin-ext"
)

func main() {
	spin := spin.New(`⦾⦿`, "%s Working...")
	spin.Work(func() error {
		time.Sleep(1000 * time.Millisecond)
		return nil
	})
	fmt.Println("Done!")
}
