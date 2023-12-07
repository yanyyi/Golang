package lru

import (
	"reflect"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1=1234 failed")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

func TestRemoveOldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := "value1", "value2", "value3"
	maxBytes := int64(len(k1) + len(k2) + len(v1) + len(v2))
	lru := New(maxBytes, nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))
	if _, ok := lru.Get(k1); ok || lru.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}

func TestOnEvicted(t *testing.T) {
	expiredKeys, expiredValues := []string{}, []string{}
	lru := New(10, func(k string, v Value) {
		expiredKeys = append(expiredKeys, k)
		if val, ok := v.(String); ok {
			expiredValues = append(expiredValues, string(val))
		}
	})
	lru.Add("key1", String("123456"))
	lru.Add("k2", String("v2"))
	lru.Add("k3", String("v3"))
	lru.Add("k4", String("v4"))
	expectedKeys := []string{"key1", "k2"}
	expectedValues := []string{"123456", "v2"}
	if !reflect.DeepEqual(expiredKeys, expectedKeys) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", expectedKeys)
	}
	if !reflect.DeepEqual(expiredValues, expectedValues) {
		t.Fatalf("Call OnEvicted failed, expect values equals to %s", expectedValues)
	}
}
