package main

import (
	"testing"
)

func TestCanDehumidfyInterior1(t *testing.T) {
	type args struct {
		rHin    float64
		tempIn  float64
		rHout   float64
		tempOut float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Warm and Humid inside. warm but dry outside",
			args: args{
				rHin:    90,
				tempIn:  20,
				rHout:   50,
				tempOut: 20,
			},
			want: true,
		},
		{
			name: "Warm and Humid inside. cold outside",
			args: args{
				rHin:    90,
				tempIn:  20,
				rHout:   50,
				tempOut: 5,
			},
			want: true,
		},
		{
			name: "Warm and Humid inside. cold outside",
			args: args{
				rHin:    90,
				tempIn:  20,
				rHout:   100,
				tempOut: 10,
			},
			want: true,
		},
		{
			name: "Warm and little bit humid inside. warm and humid outside",
			args: args{
				rHin:    50,
				tempIn:  20,
				rHout:   100,
				tempOut: 18,
			},
			want: false,
		},
		{
			name: "Inside nearly outside",
			args: args{
				rHin:    70,
				tempIn:  20,
				rHout:   60,
				tempOut: 20,
			},
			want: false,
		},
		{
			name: "Inside nearly outside, low diff",
			args: args{
				rHin:    70,
				tempIn:  20,
				rHout:   60,
				tempOut: 20,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanDehumidfyInterior(tt.args.rHin, tt.args.tempIn, tt.args.rHout, tt.args.tempOut); got != tt.want {
				t.Errorf("CanDehumidfyInterior() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCanDehumidfyInterior2(t *testing.T) {
	type args struct {
		rHin    float64
		tempIn  float64
		rHout   float64
		tempOut float64
		minDiff float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Inside nearly outside, low diff",
			args: args{
				rHin:    70,
				tempIn:  20,
				rHout:   60,
				tempOut: 20,
				minDiff: 0.05,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanDehumidfyInterior(tt.args.rHin, tt.args.tempIn, tt.args.rHout, tt.args.tempOut, tt.args.minDiff); got != tt.want {
				t.Errorf("CanDehumidfyInterior() = %v, want %v", got, tt.want)
			}
		})
	}

}
