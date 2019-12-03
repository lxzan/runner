#### Runner
> golang concurrent controller

- Package:
```go
import "github.com/lxzan/runner"
```

- Example:
```go
r := runner.New(10, 10*time.Millisecond)
for i := 0; i < 100; i++ {
  tmp := i
  task := &Task{Do: func() error {
    println(tmp)
    time.Sleep(500 * time.Millisecond)
    return nil
  }}
  r.Add(task)
}
time.Sleep(5 * time.Second)
```