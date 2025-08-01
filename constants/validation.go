package constants

const (
	FILE_NAME_MAX_LENGTH = 80
)

var (
	RESERVED_HOOKS_NAMES = []string{
		"usestate", "useeffect", "usecontext", "usereducer", "usecallback",
		"usememo", "useref", "useimperativehandle", "uselayouteffect",
		"usedebugvalue", "usedeferredvalue", "useid", "useinsertioneffect",
		"usesyncexternalstore", "usetransition", "useactionstate",
		"useformstatus", "useoptimistic",
	}
)
