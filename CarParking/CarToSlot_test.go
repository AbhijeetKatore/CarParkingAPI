package CarParking

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestAddNewCarToSlot(t *testing.T) {
	// body:=strings.NewReader(`{

	// }`)
	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddNewCarToSlot(tt.args.writer, tt.args.req)
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetUserFromID(tt.args.user_id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserFromID() got = %v, want %v", got, tt.want)
			}
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetCarFromID(tt.args.carnumber)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCarFromID() got = %v, want %v", got, tt.want)
			}
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
		})
	}
}
