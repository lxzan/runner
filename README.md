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
	"fmt"
	"github.com/lxzan/runner"
	"time"
)

func main() {
	var t0 = time.Now().UnixNano()
	r := runner.New(10, func(doc interface{}) {
		var t1 = time.Now().UnixNano()
		fmt.Printf("idx=%d, time=%dms\n", doc.(int), (t1-t0)/1000000)
	})
	r.Start()

	for i := 0; i < 100; i++ {
		r.Push(i)
	}

	go func() {
		time.Sleep(3 * time.Second)
		r.Stop()
	}()

	select {}
}

```