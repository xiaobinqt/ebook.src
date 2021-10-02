package add

import "testing"

func TestAdd(t *testing.T) {
	r := Add(1, 2)
	if r != 2 { // 故意写错
		t.Errorf("错了....")
	}
}

// go test -v -bench=BenchmarkAdd$ -run=none
func BenchmarkAdd(b *testing.B) {
	//b.StopTimer()
	//doSomeThing()
	//b.StartTimer()

	for i := 0; i < b.N; i++ {
		Add(1, 3)
	}
}
