package rds

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestGetGroupId(t *testing.T) {
	rdsAddrPort := "127.0.0.1:6379"
	rdsCli := redis.NewClient(&redis.Options{
		Addr: rdsAddrPort,
	})
	if err := rdsCli.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	act := NewRdsAction(rdsCli, "im_relation/")
	err := rdsCli.FlushDB(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	g1, err := act.GenerateGroupId()
	if err != nil {
		panic(err)
	}
	g2, err := act.GenerateGroupId()
	if err != nil {
		panic(err)
	}
	g3, err := act.GenerateGroupId()
	if err != nil {
		panic(err)
	}
	t.Logf("group id start %v", GroupIDStart)
	if g1 != fmt.Sprint(GroupIDStart) {
		t.Fatalf("failed %v != %v", g1, GroupIDStart)
	}
	if g2 != fmt.Sprint(GroupIDStart+1) {
		t.Fatalf("failed %v != %v", g2, GroupIDStart+1)
	}
	if g3 != fmt.Sprint(GroupIDStart+2) {
		t.Fatalf("failed %v != %v", g3, GroupIDStart+2)
	}

}
