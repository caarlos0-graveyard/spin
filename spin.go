// Package spin provides a very simple spinner for cli applications.
package spin

import (
	"fmt"
	"sync/atomic"
	"time"
)

// ClearLine go to the beginning of the line and clear it
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
	Spin11  = `⌜⌝⌟⌞`
	Spin12  = `┤┘┴└├┌┬┐`
	Spin13  = `⇑⇗⇒⇘⇓⇙⇐⇖`
	Spin14  = `☰☱☳☷☶☴`
	Spin15  = `䷀䷪䷡䷊䷒䷗䷁䷖䷓䷋䷠䷫`
	Default = Box1
)

// Spinner main type
type Spinner struct {
	frames []rune
	pos    int
	active uint64
	text   string
	tpf    time.Duration
}

// Option describes an option to override a default
// when creating a new Spinner.
type Option func(s *Spinner)

// New creates a Spinner object with the provided
// text. By default, the Default spinner frames are
// used, and new frames are rendered every 100 milliseconds.
// Options can be provided to override these default
// settings.
func New(text string, opts ...Option) *Spinner {
	s := &Spinner{
		text:   ClearLine + text,
		frames: []rune(Default),
		tpf:    100 * time.Millisecond,
	}
	for _, o := range opts {
		o(s)
	}
	return s
}

// WithFrames sets the frames string.
func WithFrames(frames string) Option {
	return func(s *Spinner) {
		s.Set(frames)
	}
}

// WithTimePerFrame sets how long each frame shall
// be shown.
func WithTimePerFrame(d time.Duration) Option {
	return func(s *Spinner) {
		s.tpf = d
	}
}

// Set frames to the given string which must not use spaces.
func (s *Spinner) Set(frames string) {
	s.frames = []rune(frames)
}

// Start shows the spinner.
func (s *Spinner) Start() *Spinner {
	if atomic.LoadUint64(&s.active) > 0 {
		return s
	}
	atomic.StoreUint64(&s.active, 1)
	go func() {
		for atomic.LoadUint64(&s.active) > 0 {
			fmt.Printf(s.text, s.next())
			time.Sleep(s.tpf)
		}
	}()
	return s
}

// Stop hides the spinner.
func (s *Spinner) Stop() bool {
	if x := atomic.SwapUint64(&s.active, 0); x > 0 {
		fmt.Printf(ClearLine)
		return true
	}
	return false
}

func (s *Spinner) next() string {
	r := s.frames[s.pos%len(s.frames)]
	s.pos++
	return string(r)
}
