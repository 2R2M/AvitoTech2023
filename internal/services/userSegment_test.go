package services

import (
	"avitoTech/internal/model"
	"avitoTech/internal/store"
	"avitoTech/internal/store/mocks"
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
)

func TestUserSegmentService_CreateUser(t *testing.T) {

	type args struct {
		ctx context.Context
		in  *model.User
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "base case",
			args: args{
				ctx: context.Background(),
				in: &model.User{
					ID: "1",
				},
			},
			want: &model.User{
				ID: "1",
			},
			wantErr: false,
		},
		{
			name: "missing ID",
			args: args{
				ctx: context.Background(),
				in: &model.User{
					ID: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid input, repository error",
			args: args{
				ctx: context.Background(),
				in: &model.User{
					ID: "123",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userSegmentRepoMocks := mocks.NewUserSegmentRepository(t)
			switch tt.name {
			case "base case":
				userSegmentRepoMocks.
					On("CreateUser", mock.Anything, tt.args.in).
					Return(tt.want, nil)
			case "missing ID":
				userSegmentRepoMocks.
					On("CreateUser", mock.Anything, tt.args.in).
					Return(nil, fmt.Errorf("expected error"))
			case "valid input, repository error":
				userSegmentRepoMocks.
					On("CreateUser", mock.Anything, tt.args.in).
					Return(nil, fmt.Errorf("expected repository error"))
			}

			s := UserSegmentService{
				userSegmentRepo: userSegmentRepoMocks,
			}

			got, err := s.CreateUser(tt.args.ctx, tt.args.in)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err == nil {
				t.Errorf("CreateUser() expected error, but got nil")
				return
			}

			if !tt.wantErr && got != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSegmentService_CreateSegment(t *testing.T) {

	type args struct {
		ctx context.Context
		in  *model.Segment
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Segment
		wantErr bool
	}{
		{
			name: "base case",
			args: args{
				in: &model.Segment{
					Slug: "TEST",
				},
			},
			want: &model.Segment{
				Slug: "TEST",
			},
			wantErr: false,
		},
		{
			name: "missing ID",
			args: args{
				in: &model.Segment{
					Slug: "TEST",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid input, repository error",
			args: args{
				in: &model.Segment{
					Slug: "TEST",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userSegmentRepoMocks := mocks.NewUserSegmentRepository(t)
			switch tt.name {
			case "base case":
				userSegmentRepoMocks.
					On("CreateSegment", mock.Anything, tt.args.in).
					Return(tt.want, nil)
			case "missing ID":
				userSegmentRepoMocks.
					On("CreateSegment", mock.Anything, tt.args.in).
					Return(nil, fmt.Errorf("expected error"))
			case "valid input, repository error":
				userSegmentRepoMocks.
					On("CreateSegment", mock.Anything, tt.args.in).
					Return(nil, fmt.Errorf("expected repository error"))
			}

			s := UserSegmentService{
				userSegmentRepo: userSegmentRepoMocks,
			}

			got, err := s.CreateSegment(tt.args.ctx, tt.args.in)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSegment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err == nil {
				t.Errorf("CreateSegment() expected error, but got nil")
				return
			}

			if !tt.wantErr && got != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSegment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSegmentService_GetUsersSegments(t *testing.T) {
	type fields struct {
		userSegmentRepo store.UserSegmentRepository
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Segment
		wantErr bool
	}{
		{
			name: "base case",
			fields: fields{
				userSegmentRepo: &mocks.UserSegmentRepository{},
			},
			args: args{
				ctx: context.TODO(),
				id:  "1",
			},
			want: []*model.Segment{
				&model.Segment{Slug: "TEST"},
				&model.Segment{Slug: "TEST1"},
			},
			wantErr: false,
		},
		{
			name: "invalid ID",
			fields: fields{
				userSegmentRepo: &mocks.UserSegmentRepository{},
			},
			args: args{
				ctx: context.TODO(),
				id:  "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "repository error",
			fields: fields{
				userSegmentRepo: &mocks.UserSegmentRepository{},
			},
			args: args{
				ctx: context.TODO(),
				id:  "123",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userSegmentRepoMocks := tt.fields.userSegmentRepo.(*mocks.UserSegmentRepository)
			switch tt.name {
			case "base case":
				userSegmentRepoMocks.
					On("GetUsersSegments", mock.Anything, tt.args.id).
					Return(tt.want, nil)
			case "invalid ID":
				userSegmentRepoMocks.
					On("GetUsersSegments", mock.Anything, tt.args.id).
					Return(nil, fmt.Errorf("invalid ID"))
			case "repository error":
				userSegmentRepoMocks.
					On("GetUsersSegments", mock.Anything, tt.args.id).
					Return(nil, fmt.Errorf("expected repository error"))
			}

			s := UserSegmentService{
				userSegmentRepo: userSegmentRepoMocks,
			}

			got, err := s.GetUsersSegments(tt.args.ctx, tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsersSegments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err == nil {
				t.Errorf("GetUsersSegments() expected error, but got nil")
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsersSegments() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSegmentService_GetReport(t *testing.T) {

	type args struct {
		ctx   context.Context
		month string
		year  string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "base case",
			args: args{
				ctx:   context.TODO(),
				month: "08",
				year:  "2023",
			},
			want:    []byte("report data"),
			wantErr: false,
		},
		{
			name: "invalid input",
			args: args{
				ctx:   context.TODO(),
				month: "",
				year:  "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "repository error",

			args: args{
				ctx:   context.TODO(),
				month: "08",
				year:  "2023",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userSegmentRepoMocks := mocks.NewUserSegmentRepository(t)
			switch tt.name {
			case "base case":
				userSegmentRepoMocks.
					On("GetReport", tt.args.ctx, tt.args.month, tt.args.year).
					Return([]byte("report data"), nil)
			case "invalid input":
				userSegmentRepoMocks.
					On("GetReport", tt.args.ctx, tt.args.month, tt.args.year).
					Return(nil, fmt.Errorf("invalid input"))
			case "repository error":
				userSegmentRepoMocks.
					On("GetReport", tt.args.ctx, tt.args.month, tt.args.year).
					Return(nil, fmt.Errorf("expected repository error"))
			}

			s := UserSegmentService{
				userSegmentRepo: userSegmentRepoMocks,
			}

			got, err := s.GetReport(tt.args.ctx, tt.args.month, tt.args.year)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err == nil {
				t.Errorf("GetReport() expected error, but got nil")
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReport() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSegmentService_AddUsersSegment(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *model.Operation
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Operation
		wantErr bool
	}{
		{
			name: "base case",
			args: args{
				ctx: context.Background(),
				in:  model.NewOperation("user1", []string{"segment1"}, "2023-12-31"),
			},
			want: &model.Operation{
				UserId:    "user1",
				Segment:   []string{"segment1"},
				ExpiredAt: "2023-12-31",
			},
			wantErr: false,
		},

		{
			name: "invalid segment",
			args: args{
				ctx: context.Background(),
				in:  model.NewOperation("user2", nil, "2023-12-31"),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userSegmentRepoMocks := mocks.NewUserSegmentRepository(t)
			switch tt.name {
			case "base case":
				userSegmentRepoMocks.
					On("AddUsersSegment", tt.args.ctx, tt.args.in).
					Return(model.Operation{ID: "", UserId: "fdfdf"}, nil)
			case "invalid segment":
				userSegmentRepoMocks.
					On("AddUsersSegment", tt.args.ctx, tt.args.in).
					Return(nil, fmt.Errorf("invalid input"))
			case "repository error":
				userSegmentRepoMocks.
					On("AddUsersSegment", tt.args.ctx, tt.args.in).
					Return(nil, fmt.Errorf("expected repository error"))
			}
			s := UserSegmentService{
				userSegmentRepo: userSegmentRepoMocks,
			}
			got, err := s.AddUsersSegment(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUsersSegment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddUsersSegment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSegmentService_DeleteSegment(t *testing.T) {
	type args struct {
		ctx  context.Context
		slug string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "base case",
			args: args{
				ctx:  context.Background(),
				slug: "EXISTING_SEGMENT",
			},

			wantErr: false,
		},
		{
			name: "empty slug",
			args: args{
				ctx:  context.Background(),
				slug: "",
			},
			wantErr: true,
		},
		{
			name: "repository error",
			args: args{
				ctx:  context.Background(),
				slug: "ERROR_SEGMENT",
			},

			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userSegmentRepoMocks := mocks.NewUserSegmentRepository(t)
			switch tt.name {
			case "base case":
				userSegmentRepoMocks.
					On("DeleteSegment", tt.args.ctx, tt.args.slug).
					Return(nil)
			case "empty slug":
				userSegmentRepoMocks.On("DeleteSegment", tt.args.ctx, tt.args.slug).
					Return(errors.New("empty slug"))
			case "repository error":
				userSegmentRepoMocks.On("DeleteSegment", tt.args.ctx, tt.args.slug).
					Return(errors.New("repository error"))

			}
			s := UserSegmentService{
				userSegmentRepo: userSegmentRepoMocks,
			}
			if err := s.DeleteSegment(tt.args.ctx, tt.args.slug); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSegment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
