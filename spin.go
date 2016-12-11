package spin

import (
	"fmt"
	"time"

	tj "github.com/tj/go-spin"
)

type Spin struct {
	spin   *tj.Spinner
	text   string
	active bool
}

func New(frames, text string) *Spin {
	spin := tj.New()
	spin.Set(frames)
	return &Spin{
		spin: spin,
		text: text,
	}
}

func (s *Spin) Work(task func() error) error {
	go func() {
		for s.active {
			fmt.Printf("\r %s %s", s.text, s.spin.Next())
			time.Sleep(100 * time.Millisecond)
		}
	}()
	err := task()
	s.active = false
	fmt.Printf("\r")
	return err
}
