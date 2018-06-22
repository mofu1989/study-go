package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	i = 1000000
	fmt.Println(strconv.FormatInt(int64(i),10))
	fmt.Println("Hello! World!")
}

func sprint(x interface{}) string{
	type stringer interface {
		String() string
	}
	switch x:= x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int,int8,int16,int32,int64:
		return strconv.FormatInt(x.(int64),10)
	case uint,uint8,uint16,uint32,uint64:
		return strconv.FormatUint(x.(uint64),10)
	case bool:
		return strconv.FormatBool(x)
	case float32:
		return strconv.FormatFloat(float64(x),'f',-1,32)
	case float64:
		return strconv.FormatFloat(x,'f',-1,32)

	default:
		return ""
	}
}

