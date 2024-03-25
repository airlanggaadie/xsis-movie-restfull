package model_test

import (
	"reflect"
	"testing"
	"xsis/assignment-test/model"

	"github.com/google/uuid"
)

func TestNewMovie(t *testing.T) {
	type args struct {
		request model.NewMovieRequest
	}
	tests := []struct {
		name    string
		args    args
		want    model.Movie
		wantErr bool
	}{
		{
			name: "success creation",
			args: args{
				model.NewMovieRequest{
					Title:       "test title",
					Description: "test description",
					Rating:      5,
					Image:       "test image",
				},
			},
			want: model.Movie{
				Title:       "test title",
				Description: "test description",
				Rating:      5,
				Image:       "test image",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := model.NewMovie(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMovie() error = %v, wantErr %v", err, tt.wantErr)
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

func TestNewUpdateMovie(t *testing.T) {
	testUUID := uuid.New()
	type args struct {
		id      uuid.UUID
		request model.NewMovieRequest
	}
	tests := []struct {
		name string
		args args
		want model.Movie
	}{
		{
			name: "success creation",
			args: args{
				testUUID,
				model.NewMovieRequest{
					Title:       "test title",
					Description: "test description",
					Rating:      5,
					Image:       "test image",
				},
			},
			want: model.Movie{
				Id:          testUUID,
				Title:       "test title",
				Description: "test description",
				Rating:      5,
				Image:       "test image",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := model.NewUpdateMovie(tt.args.id, tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUpdateMovie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMovieRequest_Validate(t *testing.T) {
	type fields struct {
		Title       string
		Description string
		Rating      float64
		Image       string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				Title:       "test title",
				Description: "test description",
				Rating:      5,
				Image:       "test image",
			},
			wantErr: false,
		},
		{
			name: "empty title",
			fields: fields{
				Title:       "",
				Description: "test description",
				Rating:      5,
				Image:       "test image",
			},
			wantErr: true,
		},
		{
			name: "invalid rating under 0",
			fields: fields{
				Title:       "test title",
				Description: "",
				Rating:      -2,
				Image:       "test image",
			},
			wantErr: true,
		},
		{
			name: "invalid rating more than 10",
			fields: fields{
				Title:       "test title",
				Description: "",
				Rating:      12,
				Image:       "test image",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &model.NewMovieRequest{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Rating:      tt.fields.Rating,
				Image:       tt.fields.Image,
			}
			if err := r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("NewMovieRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
