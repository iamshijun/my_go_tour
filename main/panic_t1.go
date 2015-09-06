package main

import (
	"fmt"
	"math"
)

func ConvertIntoInt16(x int) int16 {
	if math.MinInt16 <= x && x <= math.MaxInt16 {
		return int16(x)
	}

	panic(fmt.Sprintf("%d is out of int16 range", x)) //手动触发pani
}

func Int16FromInt(x int) (i int16, err error) {
	defer func() { //延迟执行匿名函数,并使用recover捕捉了panic,将panic转换成了error
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	i = ConvertIntoInt16(x)
	//发生panic时候 延迟函数被执行,执行依然是在下面的return语句之后?
	return i, nil
}

func main() {
	if _, e := Int16FromInt(655567); e != nil {
		fmt.Printf("%v", e)
	} else {
		fmt.Printf("no errors\n")
	}

}
