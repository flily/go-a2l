package syntax

import (
	"fmt"

	"github.com/flily/go-a2l/strutil"
)

type DataType int

type valueLimit struct {
	Lower string
	Upper string
}

const (
	InvalidDataType DataType = iota
	UBYTE
	SBYTE
	UWORD
	SWORD
	ULONG
	SLONG
	A_UINT64
	A_INT64
	FLOAT32_IEEE
	FLOAT64_IEEE
)

var (
	nameMapDataType = map[string]DataType{
		"UBYTE":        UBYTE,
		"SBYTE":        SBYTE,
		"UWORD":        UWORD,
		"SWORD":        SWORD,
		"ULONG":        ULONG,
		"SLONG":        SLONG,
		"A_UINT64":     A_UINT64,
		"A_INT64":      A_INT64,
		"FLOAT32_IEEE": FLOAT32_IEEE,
		"FLOAT64_IEEE": FLOAT64_IEEE,
	}

	stringMapDataType = map[DataType]string{}

	defaultValueLimit = map[DataType]valueLimit{
		UBYTE:        {"0", "255"},
		SBYTE:        {"-128", "127"},
		UWORD:        {"0", "65535"},
		SWORD:        {"-32768", "32767"},
		ULONG:        {"0", "4294967295"},
		SLONG:        {"-2147483648", "2147483647"},
		A_UINT64:     {"0", "18446744073709551615"},
		A_INT64:      {"-9223372036854775808", "9223372036854775807"},
		FLOAT32_IEEE: {"-3.403E+38", "3.403E+38"},
		FLOAT64_IEEE: {"-1.798E+308", "1.798E+308"},
	}
)

func (t DataType) String() string {
	if s, ok := stringMapDataType[t]; ok {
		return s
	}

	return "InvalidDataType"
}

func (t DataType) Pick(valueInt int64, valueFloat float64) string {
	switch t {
	case UBYTE, SBYTE, UWORD, SWORD, ULONG, SLONG, A_UINT64, A_INT64:
		return strutil.Number(valueInt)

	case FLOAT32_IEEE, FLOAT64_IEEE:
		return strutil.Number(valueFloat)
	}

	return ""
}

func GetDataTypeByName(name string) DataType {
	if dt, ok := nameMapDataType[name]; ok {
		return dt
	}

	return InvalidDataType
}

type Value struct {
	dataType    DataType
	valueInt    int64
	valueUint   uint64
	valueFloat  float32
	valueDouble float64
	LowerLimit  string
	UpperLimit  string
}

func (v *Value) setDefaultLimit() {
	if limit, ok := defaultValueLimit[v.dataType]; ok {
		v.LowerLimit = limit.Lower
		v.UpperLimit = limit.Upper
	}
}

func NewUByteValue(value uint8) *Value {
	v := &Value{
		dataType:  UBYTE,
		valueUint: uint64(value),
	}

	v.setDefaultLimit()
	return v
}

func NewSByteValue(value int8) *Value {
	v := &Value{
		dataType: SBYTE,
		valueInt: int64(value),
	}

	v.setDefaultLimit()
	return v
}

func NewUWordValue(value uint16) *Value {
	v := &Value{
		dataType:  UWORD,
		valueUint: uint64(value),
	}

	v.setDefaultLimit()
	return v
}

func NewSWordValue(value int16) *Value {
	v := &Value{
		dataType: SWORD,
		valueInt: int64(value),
	}

	v.setDefaultLimit()
	return v
}

func NewULongValue(value uint32) *Value {
	v := &Value{
		dataType:  ULONG,
		valueUint: uint64(value),
	}

	v.setDefaultLimit()
	return v
}

func NewSLongValue(value int32) *Value {
	v := &Value{
		dataType: SWORD,
		valueInt: int64(value),
	}

	v.setDefaultLimit()
	return v
}

func NewUInt64Value(value uint64) *Value {
	v := &Value{
		dataType:  A_UINT64,
		valueUint: value,
	}

	v.setDefaultLimit()
	return v
}

func NewInt64Value(value int64) *Value {
	v := &Value{
		dataType: A_INT64,
		valueInt: value,
	}

	v.setDefaultLimit()
	return v
}

func NewFloat32Value(value float32) *Value {
	v := &Value{
		dataType:   FLOAT32_IEEE,
		valueFloat: value,
	}

	v.setDefaultLimit()
	return v
}

func NewFloat64Value(value float64) *Value {
	v := &Value{
		dataType:    FLOAT64_IEEE,
		valueDouble: value,
	}

	v.setDefaultLimit()
	return v
}

func NewValue(dataType DataType) *Value {
	v := &Value{
		dataType: dataType,
	}

	v.setDefaultLimit()
	return v
}

func (v *Value) DataType() DataType {
	return v.dataType
}

func (v *Value) UByte() uint8 {
	return uint8(v.valueUint)
}

func (v *Value) SByte() int8 {
	return int8(v.valueInt)
}

func (v *Value) UWord() uint16 {
	return uint16(v.valueUint)
}

func (v *Value) SWord() int16 {
	return int16(v.valueInt)
}

func (v *Value) ULong() uint32 {
	return uint32(v.valueUint)
}

func (v *Value) SLong() int32 {
	return int32(v.valueInt)
}

func (v *Value) UInt64() uint64 {
	return v.valueUint
}

func (v *Value) Int64() int64 {
	return v.valueInt
}

func (v *Value) Float32() float32 {
	return v.valueFloat
}

func (v *Value) Float64() float64 {
	return v.valueDouble
}

func (v *Value) ValueString() string {
	result := "unknown value"

	switch v.dataType {
	case UBYTE, UWORD, ULONG, A_UINT64:
		result = strutil.Number(v.valueUint)

	case SBYTE, SWORD, SLONG, A_INT64:
		result = strutil.Number(v.valueInt)

	case FLOAT32_IEEE:
		result = strutil.Number(v.valueFloat)

	case FLOAT64_IEEE:
		result = strutil.Number(v.valueDouble)

	default:
		err := fmt.Errorf("invalid data type: %v", v.dataType)
		panic(err)
	}

	return result
}
