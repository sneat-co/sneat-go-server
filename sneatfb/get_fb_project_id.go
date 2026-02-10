package sneatfb

import (
	"os"
	"strings"
)

const (
	firebaseProjEnvVarName   = "FIREBASE_PROJECT_ID"
	gaeApplicationEnvVarName = "GAE_APPLICATION"
)

func GetFirebaseProjectID() string {
	if fbProjID := os.Getenv(firebaseProjEnvVarName); fbProjID != "" {
		return fbProjID
	}
	if gapAppID := os.Getenv(gaeApplicationEnvVarName); gapAppID != "" {
		i := strings.Index(gapAppID, "~")
		if i >= 0 {
			return gapAppID[i+1:]
		}
		return gapAppID
	}
	return "demo-sneat-app"
}
