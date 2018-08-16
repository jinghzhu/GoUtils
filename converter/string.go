package converter

import "strconv"

// Int64ToStr turns an int64 into a string.
func Int64ToStr(value int64) string {
	return strconv.FormatInt(value, 10)
}

// Int32ToStr turns an int32 into a string.
func Int32ToStr(value int32) string {
	return strconv.Itoa(int(value))
}

// Int16ToStr turns an int16 into a string.
func Int16ToStr(value int16) string {
	return strconv.FormatInt(int64(value), 10)
}

// Int8ToStr turns an int8 into a string.
func Int8ToStr(value int8) string {
	return strconv.FormatInt(int64(value), 10)
}

// Uint8ToStr turns an uint8 into a string.
func Uint8ToStr(value uint8) string {
	return strconv.FormatUint(uint64(value), 10)
}

// Uint16ToStr turns an uint16 into a string.
func Uint16ToStr(value uint16) string {
	return strconv.FormatUint(uint64(value), 10)
}

// Uint32ToStr turns an uint32 into a string.
func Uint32ToStr(value uint32) string {
	return strconv.FormatUint(uint64(value), 10)
}

// Uint64ToStr turns an uint64 into a string.
func Uint64ToStr(value uint64) string {
	return strconv.FormatUint(value, 10)
}

// Float32ToStr turns a float32 into a string.
func Float32ToStr(value float32) string {
	return strconv.FormatFloat(float64(value), 'f', -1, 32)
}

// Float64ToStr turns a float64 into a string.
func Float64ToStr(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}
