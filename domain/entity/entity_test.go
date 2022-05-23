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
				name:   "Test",
				number: "+375991112233",
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cd := Client{
				Name:  tc.args.name,
				Phone: tc.args.number,
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

func TestUpdateClientDiscount_Validate(t *testing.T) {
	t.Parallel()

	type args struct {
		name   *string
		number *string
		sale   *int8
	}

	var (
		correctName     = "Test"
		correctNumber   = "+375291112233"
		correctSale     = int8(10)
		shortName       = "OK"
		longName        = "111111111111111111111111111111111111111111111111111"
		incorrectNumber = "+375991112233"
		lowSale         = int8(-1)
		highSale        = int8(51)
	)

	type testCase struct {
		name    string
		args    args
		wantErr bool
	}

	testCases := []testCase{
		{
			name: "All fields are valid",
			args: args{
				name:   &correctName,
				number: &correctNumber,
				sale:   &correctSale,
			},
			wantErr: false,
		},
		{
			name: "Name < 3",
			args: args{
				name:   &shortName,
				number: &correctNumber,
				sale:   &correctSale,
			},
			wantErr: true,
		},
		{
			name: "Name > 50",
			args: args{
				name:   &longName,
				number: &correctNumber,
				sale:   &correctSale,
			},
			wantErr: true,
		},
		{
			name: "Incorrect number",
			args: args{
				name:   &correctName,
				number: &incorrectNumber,
				sale:   &correctSale,
			},
			wantErr: true,
		},
		{
			name: "Sale < 0",
			args: args{
				name:   &correctName,
				number: &correctNumber,
				sale:   &lowSale,
			},
			wantErr: true,
		},
		{
			name: "Sale > 50",
			args: args{
				name:   &shortName,
				number: &correctNumber,
				sale:   &highSale,
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cd := UpdateClient{
				Name:  tc.args.name,
				Phone: tc.args.number,
				Sale:  tc.args.sale,
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
