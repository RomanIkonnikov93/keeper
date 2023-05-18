package cardvalid

import "testing"

func TestCheckCard(t *testing.T) {

	tests := []struct {
		name string
		data string
		want bool
	}{
		{
			name: "valid",
			data: "4111111111111111,12,2005,123",
			want: true,
		},
		{
			name: "an extra symbol",
			data: "4111111111111111,12,2005,1231",
			want: false,
		},
		{
			name: "string",
			data: "some string",
			want: false,
		},
		{
			name: "empty",
			data: "",
			want: false,
		},
		{
			name: "missing comma",
			data: "4111111111111111,12,2005123",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := CheckCard(tt.data)
			if got != tt.want {
				t.Errorf("CheckCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}
