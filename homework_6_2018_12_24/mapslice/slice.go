package main

import (
	"fmt"
)

// CacheSlice type
type CacheSlice []string

// CreateCacheSlice create instance of CacheMap
func CreateCacheSlice() CacheSlice {
	return make(CacheSlice, 0, 0)
}

// AddToSlice add to slice
func (c *CacheSlice) AddToSlice(value ...string) {
	*c = append(*c, value...)
}

// Check -  check element exist
func (c CacheSlice) Check(idx int) bool {
	return idx >= 0 && idx < len(c)
}

// GetElemByIndex - return element by index
func (c CacheSlice) GetElemByIndex(idx int) (val string, err error) {
	if idx >= len(c) {
		err = fmt.Errorf("Error, index %d out of range", idx)
		return
	}
	val = c[idx]
	return
}

// Update - update by index if exists
func (c CacheSlice) Update(idx int, value string) (ok bool, err error) {
	if !c.Check(idx) {
		err = fmt.Errorf("Error, index %d doesn't exist", idx)
		return
	}
	ok = true
	c[idx] = value
	return
}

// Delete - delete by index if exists
func (c *CacheSlice) Delete(idx int) (ok bool, err error) {
	if !c.Check(idx) {
		err = fmt.Errorf("Error, index %d doesn't exist", idx)
		return
	}
	*c = append((*c)[:idx], (*c)[idx+1:]...)

	return
}
