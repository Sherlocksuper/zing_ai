package config

import "testing"

func TestGetConfig(t *testing.T) {
	t.Run("GetConfig", func(t *testing.T) {
		got := GetConfig()
		if got == nil {
			t.Errorf("GetConfig() = %v, want %v", got, nil)
		} else {
			t.Logf("GetConfig() = %v", got.OpenAI.Model)
		}
	})
}
