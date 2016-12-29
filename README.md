# spin-ext

A very simple spinner for cli golang apps.

Example:

```go
package main

import (
	"fmt"
	"time"

	"github.com/caarlos0/spin"
)

func main() {
	s := spin.New("%s Working...")
	s.Start()
	time.Sleep(100 * 20 * time.Millisecond)
	s.Stop()
}
```

[![asciicast](https://asciinema.org/a/97581.png)](https://asciinema.org/a/97581)
