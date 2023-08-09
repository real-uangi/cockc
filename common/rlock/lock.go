// Package rlock
// @author uangi
// @date 2023/8/9 16:13
package rlock

import (
	"github.com/real-uangi/cockc/common/rdb"
	"github.com/real-uangi/cockc/common/snowflake"
	"time"
)

const MaxTtl = 86400

type RLock struct {
	key    string
	parse  string
	locked bool
}

func New(key string) *RLock {
	return &RLock{
		key:    key,
		parse:  snowflake.NextId().String(),
		locked: false,
	}
}

func (r *RLock) TryLock(ttl int) bool {
	r.locked = rdb.TryLock(r.key, r.parse, ttl)
	return r.locked
}

func (r *RLock) Unlock() {
	if !r.locked {
		return
	}
	rdb.Unlock(r.key, r.parse)
}

func (r *RLock) Lock() {
	if !rdb.TryLock(r.key, r.parse, MaxTtl) {
		time.Sleep(100 * time.Millisecond)
		r.Lock()
	}
	r.locked = true
}
