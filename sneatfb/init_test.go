package sneatfb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitFirebaseForSneat(t *testing.T) {
	t.Run("panic_on_empty_project_id", func(t *testing.T) {
		assert.Panics(t, func() {
			InitFirebaseForSneat("", "db-id")
		})
	})
	t.Run("panic_on_empty_db_id", func(t *testing.T) {
		assert.Panics(t, func() {
			InitFirebaseForSneat("project-id", "")
		})
	})
}
