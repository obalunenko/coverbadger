package coverbadge_test

import (
	"context"
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
	err := coverbadge.Badger(context.Background(), defaultConfig)
	require.NoError(t, err)
}
