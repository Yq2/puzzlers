package q21

import (
	"sync"
	"reflect"
	"errors"
	"fmt"
)

type StrToStrMap struct {
	m sync.Map
}

func (imap *StrToStrMap) Delete(key string) {
	imap.m.Delete(key)
}

func (imap *StrToStrMap) Load(key string) (value string, ok bool) {
	v, ok := imap.m.Load(key)
	if v != nil {
		value = v.(string)
	}
	return
}

func (imap *StrToStrMap) LoadOrStore(key string, value string) (actual string, loaded bool) {
	a, loaded := imap.m.LoadOrStore(key, value)
	actual = a.(string)
	return
}

func (imap *StrToStrMap) Range(f func(key string, value string) bool) {
	fx := func(key , value interface{}) bool {
		return f(key.(string), value.(string))
	}
	imap.m.Range(fx)
}

func (imap *StrToStrMap) Store(key string, value string) {
	imap.m.Store(key, value)
}

type ConcurrentMap struct {
	keyType reflect.Type
	valueType reflect.Type
	m sync.Map
}

func NewConcurrentMap(keyType, valueType reflect.Type) (*ConcurrentMap, error) {
	if keyType == nil {
		return nil, errors.New("keyType is null")
	}
	if !keyType.Comparable() {
		return nil, errors.New("keyType is not Comparable")
	}
	if valueType == nil {
		return nil, errors.New("valueType is null")
	}
	imap := &ConcurrentMap{
		keyType:keyType,
		valueType:valueType,
	}
	return imap, nil
}

func (imap *ConcurrentMap) Delete(key reflect.Type)  {
	if reflect.TypeOf(key) != imap.keyType {
		return
	}
	imap.m.Delete(key)
}

func (imap *ConcurrentMap) LoadOrStore(key, value reflect.Type, isPanic bool) (actual interface{}, loaded bool, errMsg error) {
	if reflect.TypeOf(key) != imap.keyType {
		if isPanic {
			panic(fmt.Errorf("wrong key type: %v", reflect.TypeOf(key)))
		} else {
			return nil, false, fmt.Errorf("wrong key type: %v", reflect.TypeOf(key))
		}
	}
	if reflect.TypeOf(value) != imap.valueType {
		if isPanic {
			panic(fmt.Errorf("wrong value type: %v", reflect.TypeOf(value)))
		} else {
			return nil, false, fmt.Errorf("wrong value type: %v", reflect.TypeOf(key))
		}
	}
	actual, loaded = imap.m.LoadOrStore(key, value)
	return
}

func (imap *ConcurrentMap) Range(f func(key, value interface{}) bool) {
	imap.m.Range(f)
}

func (imap *ConcurrentMap) Store(key, value interface{}, isPanic bool) error {
	if reflect.TypeOf(key) != imap.keyType {
		if isPanic {
			panic(fmt.Errorf("wrong key type: %v", reflect.TypeOf(key)))
		} else {
			return fmt.Errorf("wrong key type: %v", reflect.TypeOf(key))
		}
	}
	if reflect.TypeOf(value) != imap.valueType {
		if isPanic {
			panic(fmt.Errorf("wrong value type: %v", reflect.TypeOf(value)))
		} else {
			return fmt.Errorf("wrong value type: %v", reflect.TypeOf(key))
		}
	}
	imap.m.Store(key, value)
	return nil
}