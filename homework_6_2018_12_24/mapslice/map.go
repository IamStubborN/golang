package main

import "fmt"

// CacheMap example
type CacheMap map[string]string

// CreateCacheMap create instance of CacheMap
func CreateCacheMap() CacheMap {
	return make(CacheMap)
}

// AddToMap add to map
func (c CacheMap) AddToMap(key, value string) error {
	if _, ok := c[key]; ok {
		return fmt.Errorf("Can't add to map, key %q is exist", key)
	}
	c[key] = value
	return nil
}

// CheckByKey is key exist
func (c CacheMap) CheckByKey(key string) bool {
	_, ok := c[key]
	return ok
}

// GetByKey get value by key
func (c CacheMap) GetByKey(key string) (string, error) {
	if value, ok := c[key]; ok {
		return value, nil
	}
	return "", fmt.Errorf("Can't find key, key %q isn't exist", key)
}

// UpdateByKey update value by key
func (c CacheMap) UpdateByKey(key, value string) {
	c[key] = value
}

// DeleteByKey by key
func (c CacheMap) DeleteByKey(key string) {
	delete(c, key)
}
