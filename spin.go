package spin

import (
	"fmt"
	"time"
)

// ClearLine go to the beggining of the line and clear it
const ClearLine = "\r\033[K"

// Spinner types.
var (
	Box1    = `⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏`
	Box2    = `⠋⠙⠚⠞⠖⠦⠴⠲⠳⠓`
	Box3    = `⠄⠆⠇⠋⠙⠸⠰⠠⠰⠸⠙⠋⠇⠆`
	Box4    = `⠋⠙⠚⠒⠂⠂⠒⠲⠴⠦⠖⠒⠐⠐⠒⠓⠋`
	Box5    = `⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠴⠲⠒⠂⠂⠒⠚⠙⠉⠁`
	Box6    = `⠈⠉⠋⠓⠒⠐⠐⠒⠖⠦⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈`
	Box7    = `⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈`
	Spin1   = `|/-\`
	Spin2   = `◴◷◶◵`
	Spin3   = `◰◳◲◱`
	Spin4   = `◐◓◑◒`
	Spin5   = `▉▊▋▌▍▎▏▎▍▌▋▊▉`
	Spin6   = `▌▄▐▀`
	Spin7   = `╫╪`
	Spin8   = `■□▪▫`
	Spin9   = `←↑→↓`
	Spin10  = `⦾⦿`
	Default = Box1
)

// Spinner main type
type Spinner struct {
	frames []rune
	length int
	pos    int
	active bool
	text   string
}

// New Spinner with args
func New(text string) *Spinner {
	s := &Spinner{
		text: ClearLine + text,
	}
	s.Set(Default)
	return s
}

// Set frames to the given string which must not use spaces.
func (s *Spinner) Set(frames string) {
	s.frames = []rune(frames)
	s.length = len(s.frames)
}

// Start shows the spinner
func (s *Spinner) Start() {
	s.active = true
	go func() {
		for s.active {
			fmt.Printf(s.text, s.next())
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

// Stop hides the spinner
func (s *Spinner) Stop() {
	s.active = false
	fmt.Printf(ClearLine)
}

func (s *Spinner) next() string {
	r := s.frames[s.pos%s.length]
	s.pos++
	return string(r)
}
