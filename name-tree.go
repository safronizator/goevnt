package goevnt

import (
	"strings"
)

type nameTree struct {
	subtree map[string]*nameTree
	val     interface{}
}

func (t *nameTree) Set(key string, val interface{}) {
	t.set(strings.Split(key, "."), val)
}

func (t *nameTree) Get(key string) *nameTree {
	return t.get(strings.Split(key, "."))
}

func (t *nameTree) Val() interface{} {
	return t.val
}

func (t *nameTree) set(parts []string, val interface{}) {
	if len(parts) == 0 {
		t.val = val
		return
	}
	name := parts[0]
	if t.subtree == nil {
		t.subtree = make(map[string]*nameTree, 1)
	}
	if _, isset := t.subtree[name]; !isset {
		t.subtree[name] = &nameTree{}
	}
	t.subtree[name].set(parts[1:], val)
}

func (t *nameTree) get(parts []string) *nameTree {
	if len(parts) == 0 {
		return t
	}
	name := parts[0]
	if _, isset := t.subtree[name]; !isset {
		return nil
	}
	return t.subtree[name].get(parts[1:])
}
