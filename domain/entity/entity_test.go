package entity

import (
	"testing"
)

func TestClientDiscount_Validate(t *testing.T) {
	t.Parallel()

	type args struct {
		name   string
		number string
	}

	type testCase struct {
		name    string
		args    args
		wantErr bool
	}

	testCases := []testCase{
		{
			name: "All fields are valid",
			args: args{
				name:   "Test",
				number: "+375291112233",
			},
			wantErr: false,
		},
		{
			name: "Name < 3",
			args: args{
				name:   "OK",
				number: "+375291112233",
			},
			wantErr: true,
		},
		{
			name: "Name > 50",
			args: args{
				name:   "111111111111111111111111111111111111111111111111111",
				number: "+375291112233",
			},
			wantErr: true,
		},
		{
			name: "Incorrect number",
			args: args{
				name:   "OK",
				number: "+375991112233",
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cd := ClientDiscount{
				ClientName:   tc.args.name,
				ClientNumber: tc.args.number,
			}

			err := cd.Validate()
			if err != nil && !tc.wantErr {
				t.Errorf("Got: %s\nWant: nil", err)
			}
			if err == nil && tc.wantErr {
				t.Errorf("Got: nil\nWant: err")
			}
		})
	}
}
