package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	for idx := 0; idx < b.N; idx++ {
		groupAnagrams(strs)
	}
}
func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "strs = [\"eat\",\"tea\",\"tan\",\"ate\",\"nat\",\"bat\"]",
			args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name: "strs = [\"\"]",
			args: args{strs: []string{""}},
			want: [][]string{{""}},
		},
		{
			name: "strs = [\"a\"]",
			args: args{strs: []string{"a"}},
			want: [][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
