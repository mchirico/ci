package static

import "strings"

func notes() string {
	s := `## Reference

Run the following command:

<trip>bash
	./run_ci.sh
<trip>

}

`
	return strings.ReplaceAll(s, "<trip>", "```")

}

func Notes() string {
	return notes()
}
