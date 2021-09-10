package coverbadge

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var coverageBadge = badge{
	Style:          "flat",
	ImageExtension: ".png",
}

func TestWriteBadgeToMd(t *testing.T) {
	mdpath := filepath.Join("testdata", "coverage_test.md")

	err := coverageBadge.writeBadgeToMd(mdpath, 22)
	require.NoError(t, err)
}

func TestGenerateBadgeBadgeURL(t *testing.T) {
	type args struct {
		cov float64
	}

	type expected struct {
		wantErr bool
		url     string
	}

	var tests = []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "",
			args: args{
				cov: 1.12,
			},
			expected: expected{
				wantErr: false,
				url:     "https://img.shields.io/badge/coverage-1.12%25-brightgreen.png?longCache=true&style=flat",
			},
		},

		{
			name: "",
			args: args{
				cov: 10.1,
			},
			expected: expected{
				wantErr: false,
				url:     "https://img.shields.io/badge/coverage-10.1%25-brightgreen.png?longCache=true&style=flat",
			},
		},
		{
			name: "",
			args: args{
				cov: 100.1,
			},
			expected: expected{
				wantErr: true,
				url:     "",
			},
		},
		{
			name: "",
			args: args{
				cov: 100.00,
			},
			expected: expected{
				wantErr: false,
				url:     "https://img.shields.io/badge/coverage-100%25-brightgreen.png?longCache=true&style=flat",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := coverageBadge.generateBadgeBadgeURL(tt.args.cov)
			if tt.expected.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected.url, got)
		})
	}
}
