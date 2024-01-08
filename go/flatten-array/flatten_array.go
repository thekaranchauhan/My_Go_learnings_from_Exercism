package flatten

func Flatten(in any) (out []any) {
	if slice, ok := in.([]any); ok {
		for _, e := range slice {
			out = append(out, Flatten(e)...)
		}
	} else if in != nil {
		out = append(out, in)
	}
	return append([]any{}, out...)
}
