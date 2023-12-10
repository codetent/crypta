package cli

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestAskInput(t *testing.T) {
	type args struct {
		r      io.Reader
		prompt string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantW   string
		wantErr bool
	}{
		{
			name: "no input",
			args: args{
				prompt: "foo",
				r:      strings.NewReader("\n"),
			},
			want:    "",
			wantW:   "foo: ",
			wantErr: false,
		},
		{
			name: "text input",
			args: args{
				prompt: "foo",
				r:      strings.NewReader("bar\n"),
			},
			want:    "bar",
			wantW:   "foo: ",
			wantErr: false,
		},
		{
			name: "unexpected eof",
			args: args{
				prompt: "foo",
				r:      strings.NewReader(""),
			},
			want:    "",
			wantW:   "foo: ",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			got, err := AskInput(tt.args.r, w, tt.args.prompt)
			if (err != nil) != tt.wantErr {
				t.Errorf("AskInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AskInput() = %v, want %v", got, tt.want)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("AskInput() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
