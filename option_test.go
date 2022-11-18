package csvton

import (
	"reflect"
	"testing"
)

func TestParseOption(t *testing.T) {
	type testOption struct {
		HasName    bool `csv:"has_name"`
		HasAddress bool `csv:"has_address"`
	}

	type testOption2 struct {
		HasName     bool `csv:"has_name"`
		HasAddress  bool `csv:"has_address"`
		DummyField  bool
		DummyField2 bool `csv:""`
		DummyField3 string
		DummyField4 string `csv:""`
	}

	type testOption3 struct {
		HasName    bool   `csv:"has_name"`
		HasAddress string `csv:"has_address"`
	}

	type args struct {
		opt any
	}
	tests := []struct {
		name    string
		args    args
		want    *Option
		wantErr bool
	}{
		{
			name: "[OK]has_name:false,has_address:true",
			args: args{
				opt: testOption{
					HasName:    false,
					HasAddress: true,
				},
			},
			want: &Option{
				value: map[string]bool{
					"has_name":    false,
					"has_address": true,
				},
			},
			wantErr: false,
		},
		{
			name: "[OK]has_name:true,has_address:false",
			args: args{
				opt: testOption{
					HasName:    true,
					HasAddress: false,
				},
			},
			want: &Option{
				value: map[string]bool{
					"has_name":    true,
					"has_address": false,
				},
			},
			wantErr: false,
		},
		{
			name: "[OK]ignore filed",
			args: args{
				opt: testOption2{
					HasName:     true,
					HasAddress:  true,
					DummyField:  true,
					DummyField2: true,
					DummyField3: "Field3",
					DummyField4: "Field4",
				},
			},
			want: &Option{
				value: map[string]bool{
					"has_name":    true,
					"has_address": true,
				},
			},
			wantErr: false,
		},
		{
			name: "[NG]unexpected field",
			args: args{
				opt: testOption3{
					HasName:     true,
					HasAddress:  "tokyo",
				},
			},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseOption(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseOption() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseOption() = %v, want %v", got, tt.want)
			}
		})
	}
}
