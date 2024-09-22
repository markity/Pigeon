package distributelock

import (
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestLock(t *testing.T) {
	cli := NewDisLockClient(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	}), "tttttt/")
	err := cli.Lock("123123")
	if err != nil {
		t.Error(err)
	}
	t.Logf("outer lock success")
	go func() {
		now := time.Now()
		err := cli.Lock("123123")
		if err != nil {
			t.Logf(err.Error())
		}
		after := time.Now()
		t.Logf("lock success, %v", after.Sub(now))
		if after.Sub(now) < time.Second*5 {
			t.Error("fail")
		} else {
			t.Logf("ok")
		}
		cli.UnLock("123123")
	}()
	time.Sleep(time.Second * 5)
	cli.UnLock("123123")
	time.Sleep(time.Second)
}
