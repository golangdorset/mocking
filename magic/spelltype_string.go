// Code generated by "stringer -type SpellType"; DO NOT EDIT.

package magic

import "strconv"

const _SpellType_name = "DarkLightFireWaterAirEarth"

var _SpellType_index = [...]uint8{0, 4, 9, 13, 18, 21, 26}

func (i SpellType) String() string {
	if i < 0 || i >= SpellType(len(_SpellType_index)-1) {
		return "SpellType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SpellType_name[_SpellType_index[i]:_SpellType_index[i+1]]
}
