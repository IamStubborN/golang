package main

import (
	"fmt"
	"github.com/IamStubborN/golang/homework_1_2018_12_05/conv"
	"reflect"
)

func main() {
	//To Float
	f64, err := conv.ToFloat64(12)
	fmt.Println(reflect.TypeOf(f64), err)
	f32, err := conv.ToFloat32(15.4)
	fmt.Println(reflect.TypeOf(f32), err)

	//To Int
	i, err := conv.ToInt(32)
	fmt.Println(reflect.TypeOf(i), err)
	i8, err := conv.ToInt8(32)
	fmt.Println(reflect.TypeOf(i8), err)
	i16, err := conv.ToInt16(32)
	fmt.Println(reflect.TypeOf(i16), err)
	i32, err := conv.ToInt32(32)
	fmt.Println(reflect.TypeOf(i32), err)
	i64, err := conv.ToInt64(32)
	fmt.Println(reflect.TypeOf(i64), err)

	//To Uint
	u, err := conv.ToUint(32)
	fmt.Println(reflect.TypeOf(u), err)
	u8, err := conv.ToUint8(-6)
	fmt.Println(reflect.TypeOf(u8), err)
	u16, err := conv.ToUint16(32)
	fmt.Println(reflect.TypeOf(u16), err)
	u32, err := conv.ToUint32(32)
	fmt.Println(reflect.TypeOf(u32), err)
	u64, err := conv.ToUint64(32)
	fmt.Println(reflect.TypeOf(u64), err)

	//To Bool
	b, err := conv.ToBool(32)
	fmt.Println(reflect.TypeOf(b), err)

	//To String
	s, err := conv.ToString(32)
	fmt.Println(reflect.TypeOf(s), err)
}
