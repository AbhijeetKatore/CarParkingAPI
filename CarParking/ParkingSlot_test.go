package CarParking

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddParkingSlot(t *testing.T) {
	body := strings.NewReader(`{
		"FloorNumber" : 21,
		"UniqueSlotNumber" : 27,
		"Occupancy" :  false
	}`)
	httpreq := httptest.NewRequest("POST", "/slot", body)
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
				req:    httpreq,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddParkingSlot(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestDeleteParkingSlots(t *testing.T) {
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
				req:    httptest.NewRequest("DELETE", "/slot", strings.NewReader(`{"uniqueslotid":50}`)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteParkingSlots(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestUpdateParkingSlot(t *testing.T) {
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
				req: httptest.NewRequest("PUT", "/slot/2", strings.NewReader(`{
					"FloorNumber" : 1111,
					"UniqueSlotNumber" : 1111,
					"Occupancy" :  false
				}`)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateParkingSlot(tt.args.writer, tt.args.req)
			res := recorder.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected Status OK but got %v", res.StatusCode)
			}
		})
	}
}

func TestGetFreeParkingSlots(t *testing.T) {
	httpreq := httptest.NewRequest("GET", "/slot", nil)
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
			name: "Test without query",
			args: args{
				writer: recorder,
				req:    httpreq,
			},
		},
		{
			name: "Test with query",
			args: args{
				writer: recorder,
				req:    httptest.NewRequest("GET", "/slot?page=1", nil),
			},
		},
		{
			name: "Test with overflown Query value",
			args: args{
				writer: recorder,
				req:    httptest.NewRequest("GET", "/slot?page=9", nil),
			},
		},
		{
			name: "Test with wrong Query value",
			args: args{
				writer: recorder,
				req:    httptest.NewRequest("GET", "/slot?page=abcd", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetFreeParkingSlots(tt.args.writer, tt.args.req)
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
