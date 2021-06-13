package test

import (
	"reflect"
	"testing"
)

// 单元测试
func TestSplit(t *testing.T) {
	//单个测试
	got := Split("i love you", "love")
	want := []string{"i ", " you"}

	//slice不能直接比较
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v got:%v", want, got)
	}

	//子测试组
	type test struct {
		input string
		sep string
		want []string
	}
	tests := map[string]test{
		"first" : {input: "cxk", sep: "x", want: []string{"c", "k"}},
		"second" : {input: "蔡徐坤", sep: "徐", want: []string{"蔡", "坤"}},
		"third" : {input:"what is your name?", sep:" ", want: []string{"what", "is", "your", "name?"}},
	}
	for name, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want){
			t.Errorf("name: %s failed, want: %v, got: %v\n", name, tc.want, got)
		}
	}

}

//性能基准测试
func BenchmarkSplit(b *testing.B) {
	// b.N不是固定的
	for i := 0; i < b.N; i++ {
		Split("test who is cxk", "t")
	}
}