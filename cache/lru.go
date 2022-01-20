/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : lru.go
*   coder: zemanzeng
*   date : 2022-01-20 19:39:49
*   desc : lru
*
================================================================*/

package cache

import (
	"container/list"
)

type LRUCache struct {
	bucket int
	list   *list.List
	keys   map[interface{}]*list.Element
}

type listElementValue struct {
	Key   interface{}
	Value interface{}
}

func NewLRUCache(bucket int) *LRUCache {
	return &LRUCache{
		bucket: bucket,
		list:   list.New(),
		keys:   make(map[interface{}]*list.Element),
	}
}

func (l *LRUCache) Set(key interface{}, value interface{}) {
	element, exist := l.keys[key]
	if exist {
		if v, ok := element.Value.(*listElementValue); ok {
			v.Value = value
			l.list.MoveToFront(element)
			return
		}
		l.deleteElement(key, element) // warn
	}

	if l.list.Len() >= l.bucket {
		l.removeOldest()
	}

	eleValue := &listElementValue{Key: key, Value: value}
	element = l.list.PushFront(eleValue)
	l.keys[key] = element
}

func (l *LRUCache) Get(key interface{}) (interface{}, bool) {
	eleValue, exist := l.keys[key]
	if !exist {
		return nil, false
	}
	if v, ok := eleValue.Value.(*listElementValue); ok {
		l.list.MoveToFront(eleValue)
		return v.Value, true
	}

	l.deleteElement(key, eleValue) // warn
	return nil, false
}

func (l *LRUCache) Keys() []interface{} {
	keys := make([]interface{}, 0)

	frontEle := l.list.Front()
	for {
		if frontEle == nil {
			break
		}
		if v, ok := frontEle.Value.(*listElementValue); ok {
			keys = append(keys, v.Key)
			frontEle = frontEle.Next()
			continue
		}

		l.list.Remove(frontEle) // warn

	}
	return keys

}

func (l *LRUCache) removeOldest() bool {
	if element := l.list.Back(); element != nil {
		if v, ok := element.Value.(*listElementValue); ok {
			l.deleteElement(v.Key, element)
			return true
		}

		l.list.Remove(element) // warn
		return true
	}
	return false
}

func (l *LRUCache) deleteElement(key interface{}, element *list.Element) {
	delete(l.keys, key)
	l.list.Remove(element)

}
