// Code generated by "stringer -type=state"; DO NOT EDIT

package decoder

import "fmt"

const _state_name = "stateBeginJSONstateEndJSONstateBeginObjectstateInObjectstateBeginMemberstateInMemberstateBeginMemberKeystateEndMemberKeystateInMemberValuestateBeginArraystateInArraystateBeginArrayValuestateEndArraystateBeginStringstateInStringstateEndStringstateInStringEscapestateInStringEscapeUstateInStringEscapeUxstateInStringEscapeUxxstateInStringEscapeUxxxstateInStringEscapeUxxxxstateBeginNumberstateInNumberstateBeginNumberFracstateInNumberFracstateBeginNumberExpstateInNumberExpstateInFalsestateInTruestateEndBoolstateInNullstateError"

var _state_index = [...]uint16{0, 14, 26, 42, 55, 71, 84, 103, 120, 138, 153, 165, 185, 198, 214, 227, 241, 260, 280, 301, 323, 346, 370, 386, 399, 419, 436, 455, 471, 483, 494, 506, 517, 527}

func (i state) String() string {
	if i < 0 || i >= state(len(_state_index)-1) {
		return fmt.Sprintf("state(%d)", i)
	}
	return _state_name[_state_index[i]:_state_index[i+1]]
}
