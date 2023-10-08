package usecase

import (
	"hackathon/model"
	"net/http"
	"testing"
)

func Test_validateInput(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		u model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{
			name: "正常系",
			args: args{
				w: nil,
				u: model.User{
					Name: "test",
					Age:  20,
				}},
			wantErr: "",
		}, {
			name: "異常系: 名前が空文字",
			args: args{
				w: nil,
				u: model.User{
					Name: "",
					Age:  20,
				}},
			wantErr: "fail: invalid user name",
		}, {
			name: "異常系: 名前が50文字以上",
			args: args{
				w: nil,
				u: model.User{
					Name: "123456789012345678901234567890123456789012345678901",
					Age:  20,
				}},
			wantErr: "fail: invalid user name",
		}, {
			name: "異常系: 年齢が20未満",
			args: args{
				w: nil,
				u: model.User{
					Name: "test",
					Age:  19,
				}},
			wantErr: "fail: invalid user age",
		}, {
			name: "異常系: 年齢が80より大きい",
			args: args{
				w: nil,
				u: model.User{
					Name: "test",
					Age:  81,
				}},
			wantErr: "fail: invalid user age",
		}, {
			name: "異常系: 名前が空文字, 年齢が20未満",
			args: args{
				w: nil,
				u: model.User{
					Name: "",
					Age:  19,
				}},
			wantErr: "fail: invalid user name",
		}, {
			name: "異常系: 名前が50文字以上, 年齢が80より大きい",
			args: args{
				w: nil,
				u: model.User{
					Name: "123456789012345678901234567890123456789012345678901",
					Age:  81,
				}},
			wantErr: "fail: invalid user name",
		}, {
			name: "異常系: 名前が空文字, 年齢が80より大きい",
			args: args{
				w: nil,
				u: model.User{
					Name: "",
					Age:  81,
				}},
			wantErr: "fail: invalid user name",
		}, {
			name: "異常系: 名前が50文字以上, 年齢が20未満",
			args: args{
				w: nil,
				u: model.User{
					Name: "123456789012345678901234567890123456789012345678901",
					Age:  19,
				}},
			wantErr: "fail: invalid user name",
		}, {
			name: "異常系: 名前が空文字, 年齢が20未満",
			args: args{
				w: nil,
				u: model.User{
					Name: "",
					Age:  19,
				}},
			wantErr: "fail: invalid user name",
		}, {
			name: "異常系: 名前が50文字以上, 年齢が80より大きい",
			args: args{
				w: nil,
				u: model.User{
					Name: "123456789012345678901234567890123456789012345678901",
					Age:  81,
				}},
			wantErr: "fail: invalid user name",
		}, {
			name: "異常系: 名前が空文字, 年齢が80より大きい",
			args: args{
				w: nil,
				u: model.User{
					Name: "",
					Age:  81,
				}},
			wantErr: "fail: invalid user name",
		}, {
			name: "異常系: 名前が50文字以上, 年齢が20未満",
			args: args{
				w: nil,
				u: model.User{
					Name: "123456789012345678901234567890123456789012345678901",
					Age:  19,
				}},
			wantErr: "fail: invalid user name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotErr := validateInput(tt.args.w, tt.args.u); gotErr != tt.wantErr {
				t.Errorf("validateInput() = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
