package cfgenz

import (
	"cmp"
	"fmt"
	"regexp"
	"strings"

	"github.com/ibrt/golang-utils/memz"
)

var (
	singleNewLinePlusRegexp = regexp.MustCompile(`\n+`)
	doubleNewLinePlusRegexp = regexp.MustCompile(`\n\n+`)
)

// Comment describes a comment.
type Comment struct {
	lines []string
	links map[string]string
}

// NewComment initializes a new comment.
func NewComment() *Comment {
	return &Comment{
		lines: make([]string, 0),
		links: make(map[string]string),
	}
}

// AddLinef formats a line and adds it to the comment.
func (c *Comment) AddLinef(format string, a ...any) *Comment {
	return c.AddLines(fmt.Sprintf(format, a...))
}

// AddLines adds the given lines to the comment.
func (c *Comment) AddLines(lines ...string) *Comment {
	for _, l := range lines {
		c.lines = append(
			c.lines,
			strings.TrimSpace(singleNewLinePlusRegexp.ReplaceAllString(l, " ")))
	}

	return c
}

// MaybeAddLink adds the given link to the comment if both label and url are not empty.
func (c *Comment) MaybeAddLink(label, url string) *Comment {
	if label != "" && url != "" {
		c.links[label] = url
	}
	return c
}

// String implements the fmt.Stringer interface.
func (c *Comment) String() string {
	lines := append(memz.ShallowCopySlice(c.lines), "")

	for _, label := range memz.GetSortedMapKeys(c.links, cmp.Less) {
		lines = append(lines, fmt.Sprintf("[%v]: %v", label, c.links[label]))
	}

	body := doubleNewLinePlusRegexp.ReplaceAllString(
		strings.TrimSpace(strings.Join(lines, "\n")),
		"\n\n")

	return strings.Join(
		memz.TransformSlice(
			strings.Split(body, "\n"),
			func(l string) string {
				return fmt.Sprintf("// %v", l)
			}),
		"\n")
}
