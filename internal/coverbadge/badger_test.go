package coverbadge_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/obalunenko/coverbadger/internal/coverbadge"
)

func TestBadger(t *testing.T) {
	defaultConfig := coverbadge.Params{
		BadgeStyle:     "flat",
		UpdateMdFiles:  filepath.Join("testdata", "coverage_test.md"),
		ManualCoverage: 90,
	}
	err := coverbadge.Badger(defaultConfig)
	require.NoError(t, err)
}
