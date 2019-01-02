package route

import "strings"

func Method(method ...string) []string {
	var methods []string
	if len(method) == 0 {
		return methods
	}
	for _, m := range method {
		methods = append(methods, strings.ToUpper(m))
	}
	return methods
}
