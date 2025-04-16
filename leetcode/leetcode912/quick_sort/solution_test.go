package quick_sort

import "testing"

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "test1",
			input: []int{5, 2, 3, 1},
			want:  []int{1, 2, 3, 5},
		},
		{
			name:  "test2",
			input: []int{5, 1, 1, 2, 0, 0},
			want:  []int{0, 0, 1, 1, 2, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortArray(tt.input)

			for i := range len(tt.want) {
				if got[i] != tt.want[i] {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			}
		})
	}
}
