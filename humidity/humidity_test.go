package humidity

import (
	"math"
	"testing"
)

func TestRelativeToAbsolute(t *testing.T) {
	type args struct {
		relH       float64
		targetTemp float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "90% 20°",
			args: args{
				relH:       90,
				targetTemp: 20,
			},
			want: 15.52,
		},
		{
			name: "80% 8°",
			args: args{
				relH:       80,
				targetTemp: 8,
			},
			want: 6.61,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RelativeToAbsolute(tt.args.relH, tt.args.targetTemp); math.Round(got*100)/100 != tt.want {
				t.Errorf("RelativeToAbsolute() = %v, want %v", math.Round(got*100)/100, tt.want)
			}
		})
	}
}
func TestRelative100ToAbsolute(t *testing.T) {
	type args struct {
		targetTemp float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "2",
			args: args{
				targetTemp: 2,
			},
			want: 5.56,
		},
		{
			name: "0",
			args: args{
				targetTemp: 0,
			},
			want: 4.85,
		},
		{
			name: "35",
			args: args{
				targetTemp: 35,
			},
			want: 39.47,
		},

		{
			name: "20",
			args: args{
				targetTemp: 20,
			},
			want: 17.24,
		},
		{
			name: "14",
			args: args{
				targetTemp: 14,
			},
			want: 12.04,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RelativeToAbsolute(100.0, tt.args.targetTemp); math.Round(got*100)/100 != tt.want {
				t.Errorf("RelativeToAbsolute(100,) = %v, want %v", math.Round(got*100)/100, tt.want)
			}
		})
	}
}

func TestAbsoluteToRelative(t *testing.T) {
	type args struct {
		ah    float64
		tempC float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "90%, 20",
			args: args{
				ah:    15.56,
				tempC: 20,
			},
			want: 90,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsoluteToRelative(tt.args.ah, tt.args.tempC); math.Round(got) != tt.want {
				t.Errorf("AbsoluteToRelative() = %v, want %v", math.Round(got), tt.want)
			}
		})
	}
}

func Test_dewPoint(t *testing.T) {
	type args struct {
		T  float64
		RH float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "25° 50%",
			args: args{
				T:  25.0,
				RH: 50,
			},
			want: 13.9,
		},
		{
			name: "2° 95%",
			args: args{
				T:  2.0,
				RH: 95,
			},
			want: 1.3,
		},
		{
			name: "40° 20%",
			args: args{
				T:  40.0,
				RH: 20,
			},
			want: 12.8,
		},
		{
			name: "40° 60%",
			args: args{
				T:  40.0,
				RH: 60.0,
			},
			want: 30.8,
		},
		{
			name: "40° 100%",
			args: args{
				T:  40.0,
				RH: 100.0,
			},
			want: 40,
		},
		{
			name: "0° 100%",
			args: args{
				T:  0.0,
				RH: 100.0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DewPointTemperature(tt.args.RH, tt.args.T); math.Round(got*10)/10 != tt.want {
				t.Errorf("dewPoint() = %v, want %v", math.Round(got*10)/10, tt.want)
			}
		})
	}
}
