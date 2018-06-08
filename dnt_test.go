package donottrack

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRequest(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		{
			"Dnt OptOut",
			args{makeRequest(map[string][]string{"Dnt": []string{"1"}})},
			OptOut,
		},
		{
			"Dnt OptIn",
			args{makeRequest(map[string][]string{"Dnt": []string{"0"}})},
			OptIn,
		},
		{
			"Dnt NotSet",
			args{makeRequest(map[string][]string{"abc": []string{"0"}})},
			NotSet,
		},
		{
			"Dnt NotSet wrong parameter",
			args{makeRequest(map[string][]string{"Dnt": []string{"null"}})},
			NotSet,
		},
		{
			"DNT OptOut",
			args{makeRequest(map[string][]string{"DNT": []string{"1"}})},
			OptOut,
		},
		{
			"DNT OptIn",
			args{makeRequest(map[string][]string{"DNT": []string{"0"}})},
			OptIn,
		},
		{
			"DNT NotSet",
			args{makeRequest(map[string][]string{"aNT": []string{"0"}})},
			NotSet,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Request(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeRequest(header http.Header) *http.Request {
	r := httptest.NewRequest("GET", "http://example.com/foo", nil)
	r.Header = header
	return r
}
