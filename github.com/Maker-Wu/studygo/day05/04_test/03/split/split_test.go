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

	tests := map[string]test{
		"simple": {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep": {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep": {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(test.input, test.sep)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("expected:%v, got:%v\n", test.want, got)
			}
		})
	}
}
