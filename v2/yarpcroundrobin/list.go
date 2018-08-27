// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package yarpcroundrobin

import (
	"time"

	yarpc "go.uber.org/yarpc/v2"
	"go.uber.org/yarpc/v2/yarpcpeerlist"
)

type listConfig struct {
	capacity int
	shuffle  bool
	seed     int64
}

var defaultListConfig = listConfig{
	capacity: 10,
	shuffle:  true,
	seed:     time.Now().UnixNano(),
}

// ListOption customizes the behavior of a roundrobin list.
type ListOption func(*listConfig)

// Capacity specifies the default capacity of the underlying
// data structures for this list.
//
// Defaults to 10.
func Capacity(capacity int) ListOption {
	return func(c *listConfig) {
		c.capacity = capacity
	}
}

// New creates a new round robin peer list.
func New(dialer yarpc.Dialer, opts ...ListOption) *List {
	cfg := defaultListConfig
	for _, o := range opts {
		o(&cfg)
	}

	plOpts := []yarpcpeerlist.ListOption{
		yarpcpeerlist.Capacity(cfg.capacity),
		yarpcpeerlist.Seed(cfg.seed),
	}
	if !cfg.shuffle {
		plOpts = append(plOpts, yarpcpeerlist.NoShuffle())
	}

	return &List{
		List: yarpcpeerlist.New(
			"roundrobin",
			dialer,
			newPeerRing(),
			plOpts...,
		),
	}
}

// List is a PeerList that rotates which peers are to be selected in a cycle.
type List struct {
	*yarpcpeerlist.List
}