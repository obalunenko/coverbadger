// Code generated by "stringer -type=badgeStyle -linecomment"; DO NOT EDIT.

package coverbadge

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[badgeStyleUnknown-0]
	_ = x[BadgeStylePlastic-1]
	_ = x[BadgeStyleFlat-2]
	_ = x[BadgeStyleFlatSquare-3]
	_ = x[BadgeStyleForTheBadge-4]
	_ = x[BadgeStyleSocial-5]
	_ = x[badgeStyleSentinel-6]
}

const _badgeStyle_name = "unknownplasticflatflat-squarefor-the-badgesocialsentinel"

var _badgeStyle_index = [...]uint8{0, 7, 14, 18, 29, 42, 48, 56}

func (i badgeStyle) String() string {
	if i >= badgeStyle(len(_badgeStyle_index)-1) {
		return "badgeStyle(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _badgeStyle_name[_badgeStyle_index[i]:_badgeStyle_index[i+1]]
}
