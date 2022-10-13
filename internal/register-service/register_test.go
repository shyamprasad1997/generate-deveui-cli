package registerservice

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_workerPool_RegisterToLorawan(t *testing.T) {
	type args struct {
		devEui string
	}
	tests := []struct {
		name    string
		wp      *workerPool
		args    args
		wantErr bool
	}{
		{
			name: "success case",
			wp: &workerPool{
				client: mockHttpClient{
					DoResponse: &http.Response{
						Status:     "200",
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewBufferString("OK")),
					},
					DoError: nil,
				},
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "error case - error in http Do()",
			wp: &workerPool{
				client: mockHttpClient{
					DoResponse: &http.Response{
						Status:     "500",
						StatusCode: 500,
						Body:       ioutil.NopCloser(bytes.NewBufferString("NOT OK")),
					},
					DoError: errors.New("failed to get response"),
				},
			},
			args:    args{},
			wantErr: true,
		},
		{
			name: "error case - response statu code more than 300",
			wp: &workerPool{
				client: mockHttpClient{
					DoResponse: &http.Response{
						Status:     "500",
						StatusCode: 500,
						Body:       ioutil.NopCloser(bytes.NewBufferString("")),
					},
					DoError: nil,
				},
			},
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.wp.RegisterToLorawan(tt.args.devEui); (err != nil) != tt.wantErr {
				t.Errorf("workerPool.RegisterToLorawan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type mockHttpClient struct {
	DoResponse *http.Response
	DoError    error
}

func (c mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return c.DoResponse, c.DoError
}

// go test -coverprofile cover.out
// go tool cover  -html="./cover.out"
