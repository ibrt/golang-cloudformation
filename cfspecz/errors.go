package cfspecz

import (
	"fmt"
)

// SpecError describes a set of issues with a spec.
type SpecError struct {
	IssuesByPath map[string][]string `json:"issuesByPath,omitempty"`
}

// Error implements the error interface.
func (e *SpecError) Error() string {
	return fmt.Sprintf("spec: found issues in %v elements", len(e.IssuesByPath))
}

// SpecIssueCollector is a utility to collect issues with a spec.
type SpecIssueCollector struct {
	issuesByElement map[SpecContext][]string
}

// NewSpecIssueCollector initializes a new spec issue collector.
func NewSpecIssueCollector() *SpecIssueCollector {
	return &SpecIssueCollector{
		issuesByElement: make(map[SpecContext][]string),
	}
}

// CollectIssue collects an issue.
func (c *SpecIssueCollector) CollectIssue(e SpecContext, format string, a ...any) *SpecIssueCollector {
	issues := c.issuesByElement[e]
	issues = append(issues, fmt.Sprintf(format, a...))
	c.issuesByElement[e] = issues
	return c
}

// MaybeCollectIssue collects an issue if "cond" is true.
func (c *SpecIssueCollector) MaybeCollectIssue(e SpecContext, cond bool, format string, a ...any) *SpecIssueCollector {
	if cond {
		return c.CollectIssue(e, format, a...)
	}
	return c
}

// MaybeToError converts the spec issue collector to an error, if any issues were collected.
func (c *SpecIssueCollector) MaybeToError() error {
	if len(c.issuesByElement) == 0 {
		return nil
	}

	issuesByPath := make(map[string][]string, len(c.issuesByElement))

	for k, v := range c.issuesByElement {
		issuesByPath[k.GetDisplayPath()] = v
	}

	return &SpecError{
		IssuesByPath: issuesByPath,
	}
}
