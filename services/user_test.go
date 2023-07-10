package services

import (
	"reflect"
	"testing"

	"github.com/pavva91/gin-gorm-rest/models"
)

func Test_userServiceImpl_ListUsers(t *testing.T) {
	tests := []struct {
		name    string
		service userServiceImpl
		want    []models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := userServiceImpl{}
			got, err := service.ListUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("userServiceImpl.ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userServiceImpl.ListUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
