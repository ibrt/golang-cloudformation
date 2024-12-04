package specz

import (
	"fmt"
	"sort"

	"github.com/ibrt/golang-utils/memz"
)

// SpecValidationErrorEntry describes an entry in the spec validation error.
type SpecValidationErrorEntry struct {
	Path   string
	Issues []string
}

// SpecValidationError describes a spec validation error.
type SpecValidationError struct {
	message string
	entries []*SpecValidationErrorEntry
}

// NewSpecValidationErrorByMessage initializes a new spec validation error.
func NewSpecValidationErrorByMessage(format string, a ...any) *SpecValidationError {
	return &SpecValidationError{
		message: fmt.Sprintf("spec validation error: %v", fmt.Sprintf(format, a...)),
	}
}

// NewSpecValidationErrorByEntries initializes a new spec validation error.
func NewSpecValidationErrorByEntries(entries []*SpecValidationErrorEntry) *SpecValidationError {
	entries = memz.ShallowCopySlice(entries)
	sort.Slice(entries, func(i, j int) bool { return entries[i].Path < entries[j].Path })

	return &SpecValidationError{
		message: fmt.Sprintf("validation error: found %v entries with issues", len(entries)),
		entries: entries,
	}
}

// GetEntries returns the entries in the spec validation error.
func (e *SpecValidationError) GetEntries() []*SpecValidationErrorEntry {
	return e.entries
}

// Error implements the error interface.
func (e *SpecValidationError) Error() string {
	return e.message
}

type issuesCollectorTarget interface {
	getDisplayPath() string
}

var (
	_ issuesCollectorTarget = (*Type)(nil)
	_ issuesCollectorTarget = (*Attribute)(nil)
	_ issuesCollectorTarget = (*Property)(nil)
)

type issuesCollector struct {
	issuesByTarget map[issuesCollectorTarget][]string
}

func newIssuesCollector() *issuesCollector {
	return &issuesCollector{
		issuesByTarget: make(map[issuesCollectorTarget][]string),
	}
}

func (i *issuesCollector) collectIssue(target issuesCollectorTarget, format string, a ...any) {
	issues := i.issuesByTarget[target]
	issues = append(issues, fmt.Sprintf(format, a...))
	i.issuesByTarget[target] = issues
}

func (i *issuesCollector) maybeToError() error {
	if len(i.issuesByTarget) == 0 {
		return nil
	}

	entries := make([]*SpecValidationErrorEntry, 0, len(i.issuesByTarget))

	for target, issues := range i.issuesByTarget {
		entries = append(entries, &SpecValidationErrorEntry{
			Path:   target.getDisplayPath(),
			Issues: issues,
		})
	}

	return NewSpecValidationErrorByEntries(entries)
}
