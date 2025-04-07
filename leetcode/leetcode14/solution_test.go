package leetcode14

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Example 1",
			args: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "Example 2",
			args: []string{"dog", "racecar", "car"},
			want: "",
		},
	}

	for _, tcase := range cases {
		t.Run(tcase.name, func(t *testing.T) {
			got := longestCommonPrefix(tcase.args)

			if got != tcase.want {
				t.Errorf("Not equal, got: %s, want: %s", got, tcase.want)
			}
		})
	}
}
