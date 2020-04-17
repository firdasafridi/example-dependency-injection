package main

import (
	"errors"
	"reflect"
	"testing"

	"github.com/firdasafridi/example-dependency-injection/cmd/_mocks"
	"github.com/firdasafridi/example-dependency-injection/country"
	"github.com/golang/mock/gomock"
)

func TestHandler_printAllCountry(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := _mocks.NewMockRepository(mockCtrl)

	// Case 1 - success to print
	result := []*country.Country{
		{
			Name: "Indonesia",
		},
	}
	mockRepository.EXPECT().GetAll().Return(result, nil).Times(1)

	// Case 2 - error get list country
	mockRepository.EXPECT().GetAll().Return(nil, errors.New("GetAll return error")).Times(1)

	type fields struct {
		Country country.Repository
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Case 1 - Success to get list country",
			fields: fields{
				Country: mockRepository,
			},
		},
		{
			name: "Case 2 - Failed to get list country",
			fields: fields{
				Country: mockRepository,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &Handler{
				Country: tt.fields.Country,
			}
			handler.printAllCountry()
		})
	}
}

func Test_newHandler(t *testing.T) {
	type args struct {
		repoCountry country.Repository
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{
			name: "Set nil new handler",
			args: args{
				repoCountry: nil,
			},
			want: &Handler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newHandler(tt.args.repoCountry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
