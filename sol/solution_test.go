package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	s := "anagram"
	t := "nagaram"
	for idx := 0; idx < b.N; idx++ {
		isAnagram(s, t)
	}
}
func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "s = \"anagram\", t = \"nagaram\"",
			args: args{s: "anagram", t: "nagaram"},
			want: true,
		},
		{
			name: "s = \"rat\", t = \"car\"",
			args: args{s: "rat", t: "car"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
