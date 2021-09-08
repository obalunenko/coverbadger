package coverbadge

import "fmt"

//go:generate stringer -type=badgeStyle -linecomment

// badgeStyle is an enumeration of badge style values.
type badgeStyle uint

const (
	badgeStyleUnknown     badgeStyle = iota // unknown
	BadgeStylePlastic                       // plastic
	BadgeStyleFlat                          // flat
	BadgeStyleFlatSquare                    // flat-square
	BadgeStyleForTheBadge                   // for-the-badge
	BadgeStyleSocial                        // social
	badgeStyleSentinel                      // sentinel
)

var badgeStyleDict = func() map[string]badgeStyle {
	res := make(map[string]badgeStyle)

	for i := badgeStyleUnknown; i < badgeStyleSentinel; i++ {
		res[i.String()] = i
	}

	return res
}()

func parseBadgeStyle(v string) (badgeStyle, error) {
	style, ok := badgeStyleDict[v]
	if !ok {
		return badgeStyleUnknown, fmt.Errorf("invalid badge stle value")
	}

	return style, nil
}

func (b badgeStyle) IsValid() bool {
	return b > badgeStyleUnknown && b < badgeStyleSentinel
}

func BadgeStyleNames() []string {
	res := make([]string, 0, badgeStyleSentinel-1-1)

	for i := badgeStyleUnknown; i < badgeStyleSentinel; i++ {
		if i.IsValid() {
			res = append(res, i.String())
		}
	}

	return res
}
