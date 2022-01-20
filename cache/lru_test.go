/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : lru_test.go
*   coder: zemanzeng
*   date : 2022-01-20 20:00:54
*   desc :
*
================================================================*/

package cache

import "testing"

func TestAvailable(t *testing.T) {

	l := NewLRUCache(3)
	l.Set("key_test1", "value_test1")
	l.Set("key_test2", "value_test2")
	l.Set("key_test3", "value_test3")
	l.Set("key_test4", "value_test4")

	// key_test4 key_test3 key_test2
	if _, exist := l.Get("key_test1"); exist {
		t.Errorf("key_test1 expect not exist")
	}

	l.Set("key_test3", "value_test3_1")
	if value, exist := l.Get("key_test2"); !exist || value != "value_test2" {
		t.Errorf("key_test2 expect exist:%v and value:%v is value_test2", exist, value)
	}
	if value, exist := l.Get("key_test3"); !exist || value != "value_test3_1" {
		t.Errorf("key_test3 expect exist:%v and value:%v is value_test3_1", exist, value)
	}

	// key_test3(value_test3_1) key_test2 key_test4
	t.Logf("list:%+v", l.Keys())

	list := l.Keys()
	if len(list) != 3 ||
		list[0].(string) != "key_test3" || list[1].(string) != "key_test2" || list[2].(string) != "key_test4" {
		t.Errorf("list:%+v sort should be [key_test3, key_test2, key_test4]", list)
	}

}
