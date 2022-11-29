package CarParking

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddUser(t *testing.T) {
	recorder := httptest.NewRecorder()
	body := strings.NewReader(`{
		"FName": "TestData",
		"LName": "TestData",
		"Age": 111
	}`)
	req := httptest.NewRequest("POST", "/user", body)
	type args struct {
		writer http.ResponseWriter
		res    *http.Request
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
				res:    req,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddUser(tt.args.writer, tt.args.res)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	req := httptest.NewRequest("GET", "/user", nil)
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
			name: "Test without Query",
			args: args{
				writer: recorder,
				req:    req,
			},
		},
		{
			name: "Test with Query",
			args: args{
				writer: recorder,
				req:    httptest.NewRequest("GET", "/user?page=1", nil),
			},
		},
		{
			name: "Test with overflown Query value",
			args: args{
				writer: recorder,
				req:    httptest.NewRequest("GET", "/user?page=9", nil),
			},
		},
		{
			name: "Test with Wrong Query value",
			args: args{
				writer: recorder,
				req:    httptest.NewRequest("GET", "/user?page=abcd", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetUser(tt.args.writer, tt.args.req)
			res := recorder.Result()
			// defer res.Body.Close()
			// op, _ := ioutil.ReadAll(res.Body)
			// t.Log(string(op))
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
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
			name: "Delete case with Available field",
			args: args{
				writer: recorder,
				req:    httptest.NewRequest(http.MethodDelete, "/user", strings.NewReader(`{"userid":55}`)),
			},
		},
		{
			name: "Delete case with Not Available field",
			args: args{
				writer: recorder,
				req:    httptest.NewRequest(http.MethodDelete, "/user", strings.NewReader(`{"userid":59}`)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteUser(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}
