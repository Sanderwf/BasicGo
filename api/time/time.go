package time

import (
	"fmt"
	"time"
)

func TimeDemo()  {
	fmt.Println("Unix时间戳： ", time.Now().Unix())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println("\nstring -> time")
	// 第一个参数是格式，第二个参数是需要转换的时间字符串
	parseTime, _ := time.Parse("2006-01-02 15:04:05", "2021-08-21 15:04:05")
	fmt.Println(parseTime)

	fmt.Println("\nunix -> time")
	time := time.Unix(time.Now().Unix(), 0)
	fmt.Println(time)
}
