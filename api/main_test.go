package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/kami988/go-1.18-sample/helloworld/def"
)

func Test_server_SayHello(t *testing.T) {
	type fields struct {
		UnimplementedGreeterServer def.UnimplementedGreeterServer
	}
	type args struct {
		ctx context.Context
		in  *def.HelloRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *def.HelloReply
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				UnimplementedGreeterServer: def.UnimplementedGreeterServer{},
			},
			args: args{
				ctx: context.Background(),
				in: &def.HelloRequest{
					Name: "Test",
				},
			},
			want: &def.HelloReply{
				Message: "Hello Test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedGreeterServer: tt.fields.UnimplementedGreeterServer,
			}
			got, err := s.SayHello(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.SayHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.SayHello() = %v, want %v", got, tt.want)
			}
		})
	}
}
