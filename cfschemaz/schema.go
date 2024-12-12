package cfschemaz

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ibrt/golang-utils/errorz"

	"github.com/ibrt/golang-cloudformation/cfspecz"
	"github.com/ibrt/golang-cloudformation/cfz"
)

// Schema describes the CloudFormation JSON schema.
type Schema struct {
	Resources map[string]*Resource
}

// NewSchemaFromBuffer parses and validates a CloudFormation JSON schema from the given buffer
// (i.e. "CloudFormationSchema.zip" file).
func NewSchemaFromBuffer(buf []byte) (*Schema, error) {
	zr, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return nil, errorz.Wrap(err)
	}

	s := &Schema{
		Resources: make(map[string]*Resource),
	}

	for _, f := range zr.File {
		utlr, err := parseResource(f)
		if err != nil {
			return nil, errorz.Wrap(err)
		}

		s.Resources[f.Name] = utlr
	}

	return s, nil
}

func parseResource(f *zip.File) (utlr *Resource, err error) {
	fr, err := f.Open()
	if err != nil {
		return nil, errorz.Wrap(err, fmt.Errorf(f.Name))
	}

	defer func() {
		if cErr := fr.Close(); cErr != nil && err == nil {
			err = errorz.Wrap(cErr, fmt.Errorf(f.Name))
		}
	}()

	d := json.NewDecoder(fr)
	d.DisallowUnknownFields()

	if err := d.Decode(&utlr); err != nil {
		return nil, errorz.Wrap(err, fmt.Errorf(f.Name))
	}

	return utlr, nil
}

// MaybeGetValidation returns the validation information for a given property, if possible.
func (s *Schema) MaybeGetValidation(pc *cfz.ProblemsCollector, specT *cfspecz.Type, specP *cfspecz.Property) *cfz.Validation {
	plt := cfz.NewProblemLocationTracker("schema")
	rName := specT.GetRelatedTopLevelResourceTypeName()
	r := s.Resources[rName]

	if r == nil {
		pc.Collect(plt, "missing resource: '%v'", rName)
		return nil
	}

	plt = plt.WithPathElements(fmt.Sprintf("resource[%v]", rName))

	if specT.IsTopLevelResourceType {
		return r.Properties[specP.Name].maybeGetValidation(
			pc, plt.WithPathElements(fmt.Sprintf("property[%v]", specP.Name)), r)
	}

	dName := specT.MaybeGetUnqualifiedStructuredTypeName()
	d := r.Definitions[dName]

	if d == nil {
		pc.Collect(plt, "missing definition: '%v'", dName)
		return nil
	}

	plt = plt.WithPathElements(fmt.Sprintf("definition[%v]", dName))

	return d.Properties[specP.Name].maybeGetValidation(
		pc, plt.WithPathElements(fmt.Sprintf("property[%v]", specP.Name)), r)
}
