package opt

import "testing"

func Test_mergeOption(t *testing.T) {
	type args struct {
		options []Option
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeOption(tt.args.options...); got != tt.want {
				t.Errorf("mergeOption() = %v, want %v", got, tt.want)
			}
		})
	}
}
