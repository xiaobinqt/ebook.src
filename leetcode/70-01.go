package main

import (
	"fmt"
	"time"
)

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	f1, f2, f3 := 1, 2, 3
	for i := 3; i <= n+1; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}

	return f3
}

func main() {
	sprintf := fmt.Sprintf("%s IDCMgr(id=%s) withdraw #v%% m1c,action=withdraw dissolve",
		time.Now().Format("2006-01-02 15:04:05"), "ssdad", 7.89)
	fmt.Println(sprintf)
}
