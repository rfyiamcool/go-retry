# retry

retry func, ensure call finish. support retry count、timeout、context.

- base delay
- backoff
- times
- context
- recovery

## Usage

example

```
package main

import (
	"errors"
	"log"

	"github.com/rfyiamcool/go-retry"
)

func main() {
	r := retry.New()
	var running = false
	err := r.Ensure(func() error {
		log.Println("enter")
		if !running {
			log.Println("111")
			running = true
			return retry.Retriable(errors.New("diy"))
		}

		log.Println("222")
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

```

more usage from retry_test.go

```
func TestBase(t *testing.T) {
	r := New(WithRecovery(), WithBaseDelay(100*time.Millisecond))
	err := r.EnsureRetryTimes(10, func() error {
		fmt.Println(time.Now())
		return Retriable(errors.New("haha"))
	})
	assert.ErrorContains(t, err, "haha")
}

func TestCtx(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	r := New(WithCtx(ctx))
	err := r.Ensure(func() error {
		t.Log(time.Now())
		return RetriableMesg("haha")
	})
	assert.Equal(t, err, ctx.Err())
}

func TestBackoff(t *testing.T) {
	bo := &Backoff{
		MinDelay: time.Duration(10 * time.Millisecond),
		MaxDelay: time.Duration(1 * time.Second),
		Factor:   2,
	}
	r := New(WithRecovery(), WithBaseDelay(100*time.Millisecond), WithBackoff(bo))
	err := r.EnsureRetryTimes(20, func() error {
		fmt.Println(time.Now())
		return Retriable(errors.New("haha"))
	})
	assert.ErrorContains(t, err, "haha")
}

func TestPanic(t *testing.T) {
	r := New(WithRecovery())
	err := r.Ensure(func() error {
		if 1 > 0 {
			panic("haha")
		}
		return nil
	})
	assert.ErrorContains(t, err, "haha")
}
```
