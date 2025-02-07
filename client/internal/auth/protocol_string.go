// Code generated by "stringer -type=Protocol -linecomment"; DO NOT EDIT.

package auth

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unset-0]
	_ = x[AndroidPhone-1]
	_ = x[AndroidWatch-2]
	_ = x[MacOS-3]
	_ = x[QiDian-4]
	_ = x[IPad-5]
	_ = x[AndroidPad-6]
}

const _Protocol_name = "UnsetAndroid PhoneAndroid WatchMacOS企点iPadAndroid Pad"

var _Protocol_index = [...]uint8{0, 5, 18, 31, 36, 42, 46, 57}

func (i Protocol) String() string {
	if i < 0 || i >= Protocol(len(_Protocol_index)-1) {
		return "Protocol(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Protocol_name[_Protocol_index[i]:_Protocol_index[i+1]]
}
