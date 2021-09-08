package coverbadge

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	logging "github.com/sirupsen/logrus"
)

type badge struct {
	Style          string
	ImageExtension string
}

func (badge badge) generateBadgeBadgeURL(coverageFloat float64) string {
	const (
		bitsize   int = 64
		badgeName     = "coverage"
	)

	url := fmt.Sprintf(
		"https://img.shields.io/badge/%s-%s%%25-brightgreen%s?longCache=true&style=%s",
		badgeName,
		strconv.FormatFloat(coverageFloat, 'G', -1, bitsize),
		badge.ImageExtension,
		badge.Style,
	)

	return url
}

var (
	regex = regexp.MustCompile(`(<a href=.*>)?\!\[coverbadger-tag-do-not-edit\]\(.*\)(<\/a>)?`)
)

func (badge badge) writeBadgeToMd(fpath string, cov float64) error {
	badgeURL := badge.generateBadgeBadgeURL(cov)
	newImageTag := fmt.Sprintf("![coverbadger-tag-do-not-edit](%s)", badgeURL)

	filedata, err := ioutil.ReadFile(filepath.Clean(fpath))
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	var markdownData string
	if string(filedata) == "" {
		markdownData = newImageTag
	} else {
		if !regex.MatchString(string(filedata)) {
			// try to add badge to the top of Markdown
			markdownData = newImageTag + "\n\n" + string(filedata)
		} else {
			markdownData = regex.ReplaceAllString(string(filedata), newImageTag)
		}
	}

	err = ioutil.WriteFile(fpath, []byte(markdownData), os.ModePerm)
	if err != nil {
		return fmt.Errorf("update markdown file: %w", err)
	}

	logging.WithFields(logging.Fields{
		"regex": regex.String(),
	}).Info("Wrote badge image to markdown file")

	return nil
}
