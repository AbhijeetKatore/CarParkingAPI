package CarParking

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddNewCarToSlot(t *testing.T) {

	recorder := httptest.NewRecorder()
	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    httptest.NewRequest("PUT", "/cartoslot/2/MH05AX4158/3", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddNewCarToSlot(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestGetUserFromID(t *testing.T) {
	type args struct {
		user_id int
	}
	tests := []struct {
		name  string
		args  args
		want  UserDetails
		want1 bool
	}{
		// TODO: Add test cases.
		{
			name: "Test for False case",
			args: args{
				user_id: 0,
			},
			want1: false,
		},
		{
			name: "Test for True case",
			args: args{
				user_id: 1,
			},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := GetUserFromID(tt.args.user_id)

			if got1 != tt.want1 {
				t.Errorf("GetUserFromID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetCarFromID(t *testing.T) {
	type args struct {
		carnumber string
	}
	tests := []struct {
		name  string
		args  args
		want  CarDetails
		want1 bool
	}{
		// TODO: Add test cases.
		{
			name: "Test for False case",
			args: args{
				carnumber: "",
			},
			want1: false,
		},
		{
			name: "Test for True case",
			args: args{
				carnumber: "MH27V4099",
			},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := GetCarFromID(tt.args.carnumber)

			if got1 != tt.want1 {
				t.Errorf("GetCarFromID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetSlotFromID(t *testing.T) {
	type args struct {
		uniqueslotid int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				uniqueslotid: 4,
			},
			want: true,
		},
		{
			name: "Test1",
			args: args{
				uniqueslotid: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSlotFromID(tt.args.uniqueslotid); got != tt.want {
				t.Errorf("GetSlotFromID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteCarFromSlot(t *testing.T) {
	recorder := httptest.NewRecorder()

	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				writer: recorder,
				req:    httptest.NewRequest("DELETE", "/cartoslot", strings.NewReader(`{"uniqueslotid":49}`)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteCarFromSlot(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}
