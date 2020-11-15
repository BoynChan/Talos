package talos

import "time"

// Author:Boyn
// Date:2020/11/15

type BaseOption struct {
	RetireTime time.Duration `default:"1m"`
}

type LockOption struct {
	BaseOption
}
