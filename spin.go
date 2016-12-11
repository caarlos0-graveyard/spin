package spin

import (
	"fmt"
	"time"

	tj "github.com/tj/go-spin"
)

const CLEAR_LINE = "\r\033[K"

// Spinner main type
type Spinner struct {
	spin   *tj.Spinner
	text   string
	active bool
}

// New Spinner with args
func New(frames, text string) *Spinner {
	spin := tj.New()
	spin.Set(frames)
	return &Spinner{
		spin: spin,
		text: CLEAR_LINE + text,
	}
}

// Work shows the spinner, execute the given task and then hide the spinner
func (s *Spinner) Work(task func() error) error {
	s.active = true
	go func() {
		for s.active {
			fmt.Printf(s.text, s.spin.Next())
			time.Sleep(100 * time.Millisecond)
		}
	}()
	err := task()
	s.active = false
	fmt.Printf(CLEAR_LINE)
	return err
}
