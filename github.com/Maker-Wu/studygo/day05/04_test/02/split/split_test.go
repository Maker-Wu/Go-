package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep string
		want []string
	}

	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}

	for _, test := range tests {
		got := Split(test.input, test.sep)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("expected:%v, got:%v\n", test.want, got)
		}
	}
}
