package monster

import(
	"testing"
)

func TestStore(t *testing.T)  {
	var mons *monster = &monster{
		Name : "张三",
		Age : 20,
		Skill : "玩",
	}
	ret := mons.Store()
	if ret != true{
		t.Fatalf("monster.Store()运行错误, 期望值: %v, 实际值: %v\n", true, ret)
	}
	t.Logf("monster.Store()运行正确")
}

func TestReStore(t *testing.T)  {
	var mons = &monster{}
	ret := mons.ResStore()
	if !ret{
		t.Fatalf("monster.ReStore()运行错误，期望值：%v，实际值：%v\n", true, ret)
	}
	t.Logf("monster.ReStore()运行正确")
}