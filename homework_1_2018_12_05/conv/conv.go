package conv

import (
	"errors"
	"fmt"
	"math"
	"math/bits"
	"strconv"
)

var maxUint uint = (1 << bits.UintSize) - 1
var maxInt = (1<<bits.UintSize)/2 - 1
var minInt = (1 << bits.UintSize) / -2

var errNegativeNotAllowed = errors.New("unable to cast negative value")

// ToString casts an interface to a string type.
func ToString(i interface{}) (string, error) {
	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(s), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case uint:
		return strconv.FormatInt(int64(s), 10), nil
	case uint64:
		return strconv.FormatInt(int64(s), 10), nil
	case uint32:
		return strconv.FormatInt(int64(s), 10), nil
	case uint16:
		return strconv.FormatInt(int64(s), 10), nil
	case uint8:
		return strconv.FormatInt(int64(s), 10), nil
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", i, i)
	}
}

// ToBool casts an interface to a bool type.
func ToBool(i interface{}) (bool, error) {
	switch b := i.(type) {
	case bool:
		return b, nil
	case nil:
		return false, nil
	case int:
		if i.(int) != 0 {
			return true, nil
		}
		return false, nil
	case string:
		return strconv.ParseBool(i.(string))
	default:
		return false, fmt.Errorf("unable to cast %#v of type %T to bool", i, i)
	}
}

// ToFloat64 casts an interface to a float32 type.
func ToFloat64(i interface{}) (float64, error) {
	switch s := i.(type) {
	case float64:
		return s, nil
	case float32:
		return float64(s), nil
	case int:
		return float64(s), nil
	case int64:
		return float64(s), nil
	case int32:
		return float64(s), nil
	case int16:
		return float64(s), nil
	case int8:
		return float64(s), nil
	case uint:
		return float64(s), nil
	case uint64:
		return float64(s), nil
	case uint32:
		return float64(s), nil
	case uint16:
		return float64(s), nil
	case uint8:
		return float64(s), nil
	case string:
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", i, i)
	}
}

// ToFloat32 casts an interface to a float32 type.
func ToFloat32(i interface{}) (result float32, err error) {
	switch s := i.(type) {
	case float64:
		if s > math.MaxFloat32 || s < -math.MaxFloat32 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return float32(s), nil
	case float32:
		return s, nil
	case int:
		return float32(s), nil
	case int64:
		return float32(s), nil
	case int32:
		return float32(s), nil
	case int16:
		return float32(s), nil
	case int8:
		return float32(s), nil
	case uint:
		return float32(s), nil
	case uint64:
		return float32(s), nil
	case uint32:
		return float32(s), nil
	case uint16:
		return float32(s), nil
	case uint8:
		return float32(s), nil
	case string:
		v, err := strconv.ParseFloat(s, 32)
		if err == nil {
			return float32(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to float32", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float32", i, i)
	}
}

// ToInt64 casts an interface to an int64 type.
func ToInt64(i interface{}) (result int64, err error) {
	switch s := i.(type) {
	case int:
		if s > math.MaxInt64 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int64(s), nil
	case int64:
		if s > math.MaxInt64 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return s, nil
	case int32:
		return int64(s), nil
	case int16:
		return int64(s), nil
	case int8:
		return int64(s), nil
	case uint:
		return int64(s), nil
	case uint64:
		if s > math.MaxInt64 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int64(s), nil
	case uint32:
		return int64(s), nil
	case uint16:
		return int64(s), nil
	case uint8:
		return int64(s), nil
	case float64:
		if s > math.MaxInt64 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int64(s), nil
	case float32:
		if s > math.MaxInt64 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int64(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	}
}

// ToInt32 casts an interface to an int32 type.
func ToInt32(i interface{}) (result int32, err error) {
	switch s := i.(type) {
	case int:
		if s > math.MaxInt32 || s < math.MinInt32 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int32(s), nil
	case int64:
		if s > math.MaxInt32 || s < math.MinInt32 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int32(s), nil
	case int32:
		if s > math.MaxInt32 || s < math.MinInt32 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return s, nil
	case int16:
		return int32(s), nil
	case int8:
		return int32(s), nil
	case uint:
		return int32(s), nil
	case uint64:
		if s > math.MaxInt32 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int32(s), nil
	case uint32:
		if s > math.MaxInt32 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int32(s), nil
	case uint16:
		return int32(s), nil
	case uint8:
		return int32(s), nil
	case float64:
		if s > math.MaxInt32 || s < math.MinInt32 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int32(s), nil
	case float32:
		if s > math.MaxInt32 || s < math.MinInt32 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int32(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int32(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
	}
}

// ToInt16 casts an interface to an int16 type.
func ToInt16(i interface{}) (result int16, err error) {
	switch s := i.(type) {
	case int:
		if s > math.MaxInt16 || s < math.MinInt16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int16(s), nil
	case int64:
		if s > math.MaxInt16 || s < math.MinInt16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int16(s), nil
	case int32:
		if s > math.MaxInt16 || s < math.MinInt16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int16(s), nil
	case int16:
		if s > math.MaxInt16 || s < math.MinInt16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return s, nil
	case int8:
		return int16(s), nil
	case uint:
		return int16(s), nil
	case uint64:
		if s > math.MaxInt16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int16(s), nil
	case uint32:
		if s > math.MaxInt16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int16(s), nil
	case uint16:
		if s > math.MaxInt16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int16(s), nil
	case uint8:
		return int16(s), nil
	case float64:
		if s > math.MaxInt16 || s < math.MinInt16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int16(s), nil
	case float32:
		if s > math.MaxInt16 || s < math.MinInt16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int16(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int16(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
	}
}

// ToInt8 casts an interface to an int8 type.
func ToInt8(i interface{}) (result int8, err error) {
	switch s := i.(type) {
	case int:
		if s > math.MaxInt8 || s < math.MinInt8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int8(s), nil
	case int64:
		if s > math.MaxInt8 || s < math.MinInt8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int8(s), nil
	case int32:
		if s > math.MaxInt8 || s < math.MinInt8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int8(s), nil
	case int16:
		if s > math.MaxInt8 || s < math.MinInt8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int8(s), nil
	case int8:
		return s, nil
	case uint:
		if s > math.MaxInt8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int8(s), nil
	case uint64:
		if s > math.MaxInt8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int8(s), nil
	case uint32:
		if s > math.MaxInt8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int8(s), nil
	case uint16:
		if s > math.MaxInt8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int8(s), nil
	case uint8:
		if s > math.MaxInt8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int8(s), nil
	case float64:
		if s > math.MaxInt8 || s < math.MinInt8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int8(s), nil
	case float32:
		if s > math.MaxInt8 || s < math.MinInt8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int8(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int8(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
	}
}

// ToInt casts an interface to an int type.
func ToInt(i interface{}) (result int, err error) {
	switch s := i.(type) {
	case int:
		if s > maxInt || s < minInt {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return s, nil
	case int64:
		if int(s) > maxInt || int(s) < minInt {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int(s), nil
	case int32:
		if int(s) > maxInt || int(s) < minInt {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int(s), nil
	case int16:
		return int(s), nil
	case int8:
		return int(s), nil
	case uint:
		return int(s), nil
	case uint64:
		if s > uint64(maxInt) {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int(s), nil
	case uint32:
		if s > uint32(maxInt) {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int(s), nil
	case uint16:
		return int(s), nil
	case uint8:
		return int(s), nil
	case float64:
		if s > float64(maxInt) || s < float64(minInt) {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int(s), nil
	case float32:
		if s > float32(maxInt) || s < float32(minInt) {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return int(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	}
}

// ToUint casts an interface to a uint type.
func ToUint(i interface{}) (result uint, err error) {
	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 0)
		if err == nil {
			return uint(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint: %s", i, err)
	case int:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if uint(s) > uint(maxUint) {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint(s), nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if uint(s) > uint(maxUint) {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case uint:
		if s > maxUint {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return s, nil
	case uint64:
		if uint(s) > uint(maxUint) {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint(s), nil
	case uint32:
		if uint(s) > uint(maxUint) {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint(s), nil
	case uint16:
		return uint(s), nil
	case uint8:
		return uint(s), nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if uint(s) > uint(maxUint) {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if uint(s) > uint(maxUint) {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
	}
}

// ToUint64 casts an interface to a uint64 type.
func ToUint64(i interface{}) (result uint64, err error) {
	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 64)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint64: %s", i, err)
	case int:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case uint:
		return uint64(s), nil
	case uint64:
		return s, nil
	case uint32:
		return uint64(s), nil
	case uint16:
		return uint64(s), nil
	case uint8:
		return uint64(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case float64:
		if s > math.MaxUint64 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
	}
}

// ToUint32 casts an interface to a uint32 type.
func ToUint32(i interface{}) (result uint32, err error) {
	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 32)
		if err == nil {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint32: %s", i, err)
	case int:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if s > math.MaxUint32 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint32(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case uint:
		return uint32(s), nil
	case uint64:
		return uint32(s), nil
	case uint32:
		return s, nil
	case uint16:
		return uint32(s), nil
	case uint8:
		return uint32(s), nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if s > math.MaxUint32 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint32(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if s > math.MaxUint32 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint32(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
	}
}

// ToUint16 casts an interface to a uint16 type.
func ToUint16(i interface{}) (result uint16, err error) {
	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 16)
		if err == nil {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint16: %s", i, err)
	case int:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if s > math.MaxUint16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint16(s), nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if s > math.MaxUint16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint16(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if s > math.MaxUint16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint16(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case uint:
		if s > math.MaxUint16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint16(s), nil
	case uint64:
		if s > math.MaxUint16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint16(s), nil
	case uint32:
		if s > math.MaxUint16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint16(s), nil
	case uint16:
		if s > math.MaxUint16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return s, nil
	case uint8:
		return uint16(s), nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if s > math.MaxUint16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint16(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if s > math.MaxUint16 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint16(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", i, i)
	}
}

// ToUint8 casts an interface to a uint type.
func ToUint8(i interface{}) (result uint8, err error) {
	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 8)
		if err == nil {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint8: %s", i, err)
	case int:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if s > math.MaxUint8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint8(s), nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if s > math.MaxUint8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint8(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if s > math.MaxUint8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint8(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if s > math.MaxUint8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint8(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(s), nil
	case uint:
		if s > math.MaxUint8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint8(s), nil
	case uint64:
		if s > math.MaxUint8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint8(s), nil
	case uint32:
		if s > math.MaxUint8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint8(s), nil
	case uint16:
		if s > math.MaxUint8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint8(s), nil
	case uint8:
		if s > math.MaxUint8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return s, nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if s > math.MaxUint8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint8(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		if s > math.MaxUint8 {
			return 0, fmt.Errorf("Can't convert to %T. %v overflows %T", result, s, result)
		}
		return uint8(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
	}
}
