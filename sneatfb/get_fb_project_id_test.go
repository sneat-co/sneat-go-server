package sneatfb

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFirebaseProjectID(t *testing.T) {
	// Can't be parallel
	tests := []struct {
		name      string
		want      string
		osEnviron []string
	}{
		{
			name:      "empty",
			want:      "demo-sneat-app",
			osEnviron: []string{},
		},
		{
			name:      "ENV_FIREBASE_PROJECT_ID",
			want:      "sneat-fb-test",
			osEnviron: []string{firebaseProjEnvVarName + "=sneat-fb-test"},
		},
		{
			name:      "ENV_GAE_APPLICATION_with_region_prefix",
			want:      "sneat-gae-test",
			osEnviron: []string{gaeApplicationEnvVarName + "=e~sneat-gae-test"},
		},
		{
			name:      "ENV_GAE_APPLICATION_without_prefix",
			want:      "gae-test-app",
			osEnviron: []string{gaeApplicationEnvVarName + "=gae-test-app"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Can't be parallel
			backupEnv := make(map[string]string, len(tt.osEnviron))
			defer func() {
				for k, v := range backupEnv {
					if v == "" {
						_ = os.Unsetenv(k)
					} else {
						_ = os.Setenv(k, v)
					}
				}
			}()
			for _, ev := range tt.osEnviron {
				vals := strings.Split(ev, "=")
				k, v := vals[0], vals[1]
				backupEnv[k] = os.Getenv(k)
				if err := os.Setenv(k, v); err != nil {
					t.Fatalf("failed to set env var %s=%s: %v", k, v, err)
				}
			}
			firebaseProjectID := GetFirebaseProjectID()
			assert.Equal(t, tt.want, firebaseProjectID)
		})
	}
}
