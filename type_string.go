// Code generated by "stringer -type=Type"; DO NOT EDIT

package j

import "fmt"

const _Type_name = "InvalidTypeObjectTypeMemberTypeArrayTypeBoolTypeNumberTypeStringTypeNullType"

var _Type_index = [...]uint8{0, 11, 21, 31, 40, 48, 58, 68, 76}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
