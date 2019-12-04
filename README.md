#### Runner
> golang concurrent controller

- Package:
```go
import "github.com/lxzan/runner"
```

- Example:
```go
ctx, cancel := context.WithCancel(context.Background())
r := runner.New(ctx, 10, 10*time.Millisecond)
for i := 0; i < 100; i++ {
	tmp := i
	r.Add(&runner.Task{Work: func() error {
		println(tmp)
		time.Sleep(500 * time.Millisecond)
		return nil
	}})
}

r.OnClose = func() {
	println("hello")
}

quit := make(chan os.Signal)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit
cancel()
```