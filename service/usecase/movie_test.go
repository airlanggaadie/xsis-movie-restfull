package usecase

import (
	"context"
	"reflect"
	"testing"
	"xsis/assignment-test/model"
	"xsis/assignment-test/service"

	"github.com/google/uuid"
)

func TestNewMovieUsecase(t *testing.T) {
	movieRepositoryMock := movieRepositorySuccessMock{}
	type args struct {
		movieRepository service.MovieRepository
	}
	tests := []struct {
		name string
		args args
		want service.MovieUsecase
	}{
		{
			name: "success",
			args: args{
				movieRepository: movieRepositoryMock,
			},
			want: movie{
				movieRepository: movieRepositoryMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMovieUsecase(tt.args.movieRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMovieUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_movie_GetMoviesPaginate(t *testing.T) {
	movieRepositorySuccessMock := movieRepositorySuccessMock{}
	movieRepositoryFailMock := movieRepositoryFailMock{}
	type fields struct {
		movieRepository service.MovieRepository
	}
	type args struct {
		ctx   context.Context
		page  int
		limit int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.ListMovieResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				movieRepository: movieRepositorySuccessMock,
			},
			args: args{
				ctx:   context.Background(),
				page:  1,
				limit: 10,
			},
			want: model.ListMovieResponse{
				Data:  []model.Movie{},
				Total: 0,
			},
			wantErr: false,
		},
		{
			name: "success with default page and limit",
			fields: fields{
				movieRepository: movieRepositorySuccessMock,
			},
			args: args{
				ctx:   context.Background(),
				page:  0,
				limit: 0,
			},
			want: model.ListMovieResponse{
				Data:  []model.Movie{},
				Total: 0,
			},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				movieRepository: movieRepositoryFailMock,
			},
			args: args{
				ctx:   context.Background(),
				page:  0,
				limit: 0,
			},
			want:    model.ListMovieResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := movie{
				movieRepository: tt.fields.movieRepository,
			}
			got, err := m.GetMoviesPaginate(tt.args.ctx, "", tt.args.page, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("movie.GetMoviesPaginate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("movie.GetMoviesPaginate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_movie_GetMovie(t *testing.T) {
	movieRepositorySuccessMock := movieRepositorySuccessMock{}
	movieRepositoryFailMock := movieRepositoryFailMock{}
	type fields struct {
		movieRepository service.MovieRepository
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.MovieDetailResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				movieRepository: movieRepositorySuccessMock,
			},
			args: args{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			want: model.MovieDetailResponse{
				Movie: model.Movie{},
			},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				movieRepository: movieRepositoryFailMock,
			},
			args: args{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			want:    model.MovieDetailResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := movie{
				movieRepository: tt.fields.movieRepository,
			}
			got, err := m.GetMovie(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("movie.GetMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("movie.GetMovie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_movie_AddNewMovie(t *testing.T) {
	movieRepositorySuccessMock := movieRepositorySuccessMock{}
	movieRepositoryFailMock := movieRepositoryFailMock{}
	idUUID, _ := uuid.Parse("fe9ba0d7-3f78-472d-9c75-5c5ade522b84")
	type fields struct {
		movieRepository service.MovieRepository
	}
	type args struct {
		ctx     context.Context
		request model.NewMovieRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.MovieDetailResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				movieRepository: movieRepositorySuccessMock,
			},
			args: args{
				ctx: context.Background(),
				request: model.NewMovieRequest{
					Title:       "test",
					Description: "description test",
					Rating:      5,
					Image:       "image test",
				},
			},
			want: model.MovieDetailResponse{
				Movie: model.Movie{
					Id:          idUUID,
					Title:       "test",
					Description: "description test",
					Rating:      5,
					Image:       "image test",
				},
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				movieRepository: movieRepositoryFailMock,
			},
			args: args{
				ctx: context.Background(),
				request: model.NewMovieRequest{
					Title:       "test",
					Description: "description test",
					Rating:      5,
					Image:       "image test",
				},
			},
			want:    model.MovieDetailResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := movie{
				movieRepository: tt.fields.movieRepository,
			}
			got, err := m.AddNewMovie(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("movie.AddNewMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Title != tt.want.Title {
				t.Errorf("NewMovie() Title = %v, want %v", got.Title, tt.want.Title)
			}

			if got.Description != tt.want.Description {
				t.Errorf("NewMovie() Description = %v, want %v", got.Description, tt.want.Description)
			}

			if got.Rating != tt.want.Rating {
				t.Errorf("NewMovie() Rating = %v, want %v", got.Rating, tt.want.Rating)
			}

			if got.Image != tt.want.Image {
				t.Errorf("NewMovie() Image = %v, want %v", got.Image, tt.want.Image)
			}

			if _, err := uuid.Parse(got.Id.String()); err != nil {
				t.Errorf("NewMovie() Id = %v, want uuid format", got.Id)
			}
		})
	}
}

func Test_movie_UpdateMovie(t *testing.T) {
	movieRepositorySuccessMock := movieRepositorySuccessMock{}
	movieRepositoryFailMock := movieRepositoryFailMock{}
	idUUID, _ := uuid.Parse("fe9ba0d7-3f78-472d-9c75-5c5ade522b84")
	type fields struct {
		movieRepository service.MovieRepository
	}
	type args struct {
		ctx     context.Context
		id      uuid.UUID
		request model.NewMovieRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.MovieDetailResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				movieRepository: movieRepositorySuccessMock,
			},
			args: args{
				ctx: context.Background(),
				id:  uuid.New(),
				request: model.NewMovieRequest{
					Title:       "test",
					Description: "description test",
					Rating:      5,
					Image:       "image test",
				},
			},
			want: model.MovieDetailResponse{
				Movie: model.Movie{
					Id:          idUUID,
					Title:       "test",
					Description: "description test",
					Rating:      5,
					Image:       "image test",
				},
			},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				movieRepository: movieRepositoryFailMock,
			},
			args: args{
				ctx: context.Background(),
				id:  uuid.New(),
				request: model.NewMovieRequest{
					Title:       "test",
					Description: "description test",
					Rating:      5,
					Image:       "image test",
				},
			},
			want:    model.MovieDetailResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := movie{
				movieRepository: tt.fields.movieRepository,
			}
			got, err := m.UpdateMovie(tt.args.ctx, tt.args.id, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("movie.UpdateMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("movie.UpdateMovie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_movie_DeleteMovie(t *testing.T) {
	movieRepositorySuccessMock := movieRepositorySuccessMock{}
	movieRepositoryFailMock := movieRepositoryFailMock{}
	type fields struct {
		movieRepository service.MovieRepository
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				movieRepository: movieRepositorySuccessMock,
			},
			args: args{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				movieRepository: movieRepositoryFailMock,
			},
			args: args{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := movie{
				movieRepository: tt.fields.movieRepository,
			}
			if err := m.DeleteMovie(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("movie.DeleteMovie() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
