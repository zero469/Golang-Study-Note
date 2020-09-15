package main
import(
	"testing"
)

func TestSub(t *testing.T)  {
	res := Sub(5 , 10)
	if res != 0{
		t.Fatalf("Sub(5,10)错误，期望值为: %v，实际值为: %v\n", 0, res)
	}
	t.Logf("Sub(5,10)测试正确")
}