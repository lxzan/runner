#### Runner
> golang concurrent controller

- Package:
```
github.com/lxzan/runner
```

- Example:
```go
package main

import (
	"context"
	"fmt"
	"github.com/lxzan/runner"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	r := runner.New(ctx, 10, 10*time.Millisecond)
	r.OnClose = func(data []runner.TaskInterface) {
		println(fmt.Sprintf("rest tasks: %d", len(data)))
	}

	for i := 0; i < 100; i++ {
		tmp := i
		r.Add(&runner.Task{Work: func() error {
			println(tmp)
			time.Sleep(500 * time.Millisecond)
			return nil
		}})
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cancel()
	<-r.Wait
}
```