package q3

import (
	"testing"
	"time"
)

func BenchmarkGetPrimes(b *testing.B) {
	// 你可以注释或者还原下面这四行代码中的第一行和第四行，
	// 并观察测试结果的不同。
	//停止计时器
	b.StopTimer()
	time.Sleep(time.Millisecond * 500) // 模拟某个耗时但与被测程序关系不大的操作。
	max := 10000
	//开始计时器
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		GetPrimes(max)
	}
}
