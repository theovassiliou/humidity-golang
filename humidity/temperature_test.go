package humidity

import (
	"testing"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	type args struct {
		celsius float64
	}
	tests := []struct {
		name           string
		args           args
		wantFahrenheit float64
	}{
		{
			name: "0",
			args: args{
				celsius: 0,
			},
			wantFahrenheit: 32,
		},
		{
			name: "40",
			args: args{
				celsius: 40,
			},
			wantFahrenheit: 104,
		},
		{
			name: "-17.78",
			args: args{
				celsius: -17.74,
			},
			wantFahrenheit: 0.06800000000000139,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFahrenheit := CelsiusToFahrenheit(tt.args.celsius); gotFahrenheit != tt.wantFahrenheit {
				t.Errorf("CelsiusToFahrenheit() = %v, want %v", gotFahrenheit, tt.wantFahrenheit)
			}
		})
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	type args struct {
		fahrenheit float64
	}
	tests := []struct {
		name        string
		args        args
		wantCelsius float64
	}{
		{
			name: "104",
			args: args{
				fahrenheit: 104,
			},
			wantCelsius: 40,
		},
		{
			name: "32",
			args: args{
				fahrenheit: 32,
			},
			wantCelsius: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCelsius := FahrenheitToCelsius(tt.args.fahrenheit); gotCelsius != tt.wantCelsius {
				t.Errorf("FahrenheitToCelsius() = %v, want %v", gotCelsius, tt.wantCelsius)
			}
		})
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type args struct {
		celsius float64
	}
	tests := []struct {
		name       string
		args       args
		wantKelvin float64
	}{
		{
			name: "0",
			args: args{
				celsius: 0,
			},
			wantKelvin: 273.15,
		},
		{
			name: "100",
			args: args{
				celsius: 100,
			},
			wantKelvin: 373.15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotKelvin := CelsiusToKelvin(tt.args.celsius); gotKelvin != tt.wantKelvin {
				t.Errorf("CelsiusToKelvin() = %v, want %v", gotKelvin, tt.wantKelvin)
			}
		})
	}
}

func TestKelvinToCelsiu(t *testing.T) {
	type args struct {
		kelvin float64
	}
	tests := []struct {
		name        string
		args        args
		wantCelsius float64
	}{
		{
			name: "0",
			args: args{
				kelvin: 0,
			},
			wantCelsius: -273.15,
		},
		{
			name: "100",
			args: args{
				kelvin: 273.15,
			},
			wantCelsius: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCelsius := KelvinToCelsiu(tt.args.kelvin); gotCelsius != tt.wantCelsius {
				t.Errorf("KelvinToCelsiu() = %v, want %v", gotCelsius, tt.wantCelsius)
			}
		})
	}
}
