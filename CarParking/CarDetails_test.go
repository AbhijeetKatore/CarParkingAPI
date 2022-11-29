package CarParking

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddCarDetails(t *testing.T) {
	body := strings.NewReader(`{
		"CarNumber" :"MH05AX4158",
		"CarModel"  :"Hyundai i10"
	}`)
	recorder := httptest.NewRecorder()
	httpreq := httptest.NewRequest("POST", "/", body)

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
				req:    httpreq,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddCarDetails(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestDeleteCarDetails(t *testing.T) {

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
				req:    httptest.NewRequest("DELETE", "/car", strings.NewReader(`{"carnumber":"MH05AX4158"}`)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteCarDetails(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestUpdateCarDetails(t *testing.T) {
	body := strings.NewReader(`{
	"CarNumber" :"MH05AX4158",
	"CarModel"  :"Hyundai i20"
}`)
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
				req:    httptest.NewRequest("PUT", "/car/md", body),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateCarDetails(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}
