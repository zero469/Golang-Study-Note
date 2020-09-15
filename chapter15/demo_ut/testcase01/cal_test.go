package main
import(
	_"fmt"
	"testing"
)

func TestAddUpper(t *testing.T)  {
	res := AddUpper(10)
	if res != 55{
		t.Fatalf("addUpper(10)执行错误，期望值：%v，实际制：%v", 55, res)
	}
	t.Logf("addUpper(10)执行正确")
}