package talos

import (
	"testing"
	"time"

	"github.com/creasty/defaults"
	"github.com/stretchr/testify/require"
)

// Author:Boyn
// Date:2020/11/15

func TestBaseOptionSetDefault(t *testing.T) {
	opt := BaseOption{}
	err := defaults.Set(&opt)
	require.Nil(t, err)
	require.Equal(t, 1*time.Minute, opt.RetireTime)
}
