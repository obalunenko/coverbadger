package coverbadge

import (
	"fmt"
	"strings"
)

// Params holds Badger parameters.
type Params struct {
	BadgeStyle     string
	UpdateMdFiles  string
	ManualCoverage float64
	ImgExt         string
}

// Badger updates cover badge according tp Params.
func Badger(p Params) error {
	style, err := parseBadgeStyle(p.BadgeStyle)
	if err != nil {
		return fmt.Errorf("invalid badge style flag: %w", err)
	}

	b := badge{
		Style:          style.String(),
		ImageExtension: p.ImgExt,
	}

	cov := p.ManualCoverage

	if p.UpdateMdFiles != "" {
		for _, filepath := range strings.Split(p.UpdateMdFiles, ",") {
			if err := b.writeBadgeToMd(filepath, cov); err != nil {
				return fmt.Errorf("write badge to markdown[%s]: %w", filepath, err)
			}
		}
	}

	return nil
}
