package distributelock

import (
	"errors"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestLock(t *testing.T) {
	cli := NewDisLockClient(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	}), "tttttt/")
	l, err := cli.Lock("123123", time.Second*15)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("outer lock success")
	go func() {
		now := time.Now()
		l, err := cli.Lock("123123", time.Second*15)
		if err != nil {
			t.Error(err)
			return
		}
		after := time.Now()
		t.Logf("lock success, %v", after.Sub(now))
		if after.Sub(now) < time.Second*5 {
			t.Error("fail")
		} else {
			t.Logf("ok")
		}
		err = l.UnLock()
		if err != nil {
			t.Error(err)
		}
	}()
	time.Sleep(time.Second * 5)
	err = l.UnLock()
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Second)
}

func TestLock2(t *testing.T) {
	cli := NewDisLockClient(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	}), "tttttt/")
	l, err := cli.Lock("123123", time.Second*5)
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(time.Second * 7)
	err = l.UnLock()
	if err == nil {
		t.Error(errors.New("should have error"))
	}
}
