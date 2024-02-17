package example

import "testing"

func TestUseCase_CreateUser(t *testing.T) {
	type fields struct {
		Repo UserRepo
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				Repo: StubUserRepo{
					NextIdRes0: 1,
					NextIdRes1: nil,
					StoreRes0:  nil,
				},
			},
			args: args{
				name: "user",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				Repo: tt.fields.Repo,
			}
			if err := u.CreateUser(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
