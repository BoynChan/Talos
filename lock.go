package talos

import (
	"github.com/creasty/defaults"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

// Author:Boyn
// Date:2020/11/14
const unLockLua = `
if redis.call("get",KEYS[1]) == ARGV[1] then
	return redis.call("del",KEYS[1])
else
	return 0
end
`

type Lock interface {
	Lock() error
	UnLock() error
}

type lockImpl struct {
	cli       *redis.Client
	LockValue string
	Key       string
	Option    LockOption
}

func (l *lockImpl) Lock() error {
	return l.cli.SetNX(l.Key, l.LockValue, l.Option.RetireTime).Err()
}

func (l *lockImpl) UnLock() error {
	return l.cli.Eval(unLockLua, []string{l.Key}, l.LockValue).Err()
}

func NewLock(cli *redis.Client, key string, option ...LockOption) Lock {
	value := uuid.New().String()
	if len(option) == 0 {
		opt := LockOption{}
		_ = defaults.Set(&opt)
		option = append(option, opt)
	}
	return &lockImpl{
		cli:       cli,
		LockValue: value,
		Key:       key,
		Option:    option[0],
	}
}
