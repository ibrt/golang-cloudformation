package cfspecz

import (
	"fmt"
)

// ProblemLocation describes a location in the spec (i.e. spec, type, property, or attribute).
type ProblemLocation interface {
	GetDisplayPath() string
}

var (
	_ ProblemLocation = (*Spec)(nil)
	_ ProblemLocation = (*Type)(nil)
	_ ProblemLocation = (*Property)(nil)
	_ ProblemLocation = (*Attribute)(nil)
)

// ProblemsFoundError describes a set of problems within a spec.
type ProblemsFoundError struct {
	ProblemsByDisplayPath map[string][]string `json:"problemsByDisplayPath,omitempty"`
}

// ProblemsFoundError implements the error interface.
func (e *ProblemsFoundError) Error() string {
	return fmt.Sprintf("found problems in %v paths", len(e.ProblemsByDisplayPath))
}

// ProblemsCollector is a utility to collect problems.
type ProblemsCollector struct {
	problemsByLocation map[ProblemLocation][]string
}

// NewProblemsCollector initializes a new problems collector.
func NewProblemsCollector() *ProblemsCollector {
	return &ProblemsCollector{
		problemsByLocation: make(map[ProblemLocation][]string),
	}
}

// Collect collects a problem.
func (c *ProblemsCollector) Collect(sc ProblemLocation, format string, a ...any) *ProblemsCollector {
	problems := c.problemsByLocation[sc]
	problems = append(problems, fmt.Sprintf(format, a...))
	c.problemsByLocation[sc] = problems
	return c
}

// MaybeCollect collects a problem if "cond" is true.
func (c *ProblemsCollector) MaybeCollect(sc ProblemLocation, cond bool, format string, a ...any) *ProblemsCollector {
	if cond {
		return c.Collect(sc, format, a...)
	}
	return c
}

// ToError return an error if any problems were collected, nil otherwise.
func (c *ProblemsCollector) ToError() error {
	if len(c.problemsByLocation) == 0 {
		return nil
	}

	problemsByDisplayPath := make(map[string][]string, len(c.problemsByLocation))

	for k, v := range c.problemsByLocation {
		problemsByDisplayPath[k.GetDisplayPath()] = v
	}

	return &ProblemsFoundError{
		ProblemsByDisplayPath: problemsByDisplayPath,
	}
}
