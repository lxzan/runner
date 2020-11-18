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
	"time"
)

func main() {
	r := runner.New(context.Background(), 10)
	r.OnClose = func(data []interface{}) {
		println(fmt.Sprintf("rest tasks: %d", len(data)))
	}

	var t0 = time.Now().UnixNano()
	r.OnTask = func(opt interface{}) {
		var t1 = time.Now().UnixNano()
		fmt.Printf("idx=%d, time=%dms\n", opt.(int), (t1-t0)/1000000)
		time.Sleep(200 * time.Millisecond)
	}

	for i := 0; i < 100; i++ {
		r.Add(i)
	}

	select {}
}
```