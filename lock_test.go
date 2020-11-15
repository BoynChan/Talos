package talos

import (
	"testing"

	"github.com/go-redis/redis"
)

// Author:Boyn
// Date:2020/11/14

func TestLock(t *testing.T) {
	cli := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	locker := NewLock(cli, "NX_LOCK")
	err := locker.Lock()
	if err != nil {
		t.Fatal(err)
	}
	err = locker.UnLock()
	if err != nil {
		t.Fatal(err)
	}
}
