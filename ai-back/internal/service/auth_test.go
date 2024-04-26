package service

import "testing"

func TestAuthService_Backup(t *testing.T) {
	authService := NewAuthService()
	err := authService.Backup()
	if err != nil {
		t.Error(err)
	}
}
