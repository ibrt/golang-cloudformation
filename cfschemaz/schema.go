package cfschemaz

import (
	"archive/zip"
	"bytes"
	"cmp"
	"encoding/json"
	"fmt"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/memz"

	"github.com/ibrt/golang-cloudformation/cfz"
)

// Schema describes the CloudFormation JSON schema.
type Schema struct {
	UnprocessedByFileName map[string]*UnprocessedTopLevelResource
	TopLevelResourceTypes map[string]*Type
	StructuredTypes       map[string]*Type
}

// NewSchemaFromBuffer parses and validates a CloudFormation JSON schema from the given buffer
// (i.e. "CloudFormationSchema.zip" file).
func NewSchemaFromBuffer(buf []byte) (*Schema, error) {
	zr, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return nil, errorz.Wrap(err)
	}

	s := &Schema{
		UnprocessedByFileName: make(map[string]*UnprocessedTopLevelResource),
		TopLevelResourceTypes: make(map[string]*Type),
		StructuredTypes:       make(map[string]*Type),
	}

	for _, f := range zr.File {
		utlr, err := parseUnprocessedTopLevelResource(f)
		if err != nil {
			return nil, errorz.Wrap(err)
		}

		s.UnprocessedByFileName[f.Name] = utlr
	}

	if err := s.process(); err != nil {
		return nil, errorz.Wrap(err)
	}

	return s, nil
}

func parseUnprocessedTopLevelResource(f *zip.File) (utlr *UnprocessedTopLevelResource, err error) {
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

func (s *Schema) process() error {
	plt := cfz.NewProblemLocationTracker("schema")

	for _, fileName := range memz.GetSortedMapKeys(s.UnprocessedByFileName, cmp.Less) {
		utlr := s.UnprocessedByFileName[fileName]

		tlr, sts, err := utlr.toTypes(plt.WithPathElements(fmt.Sprintf("fileName[%v]", fileName)))
		if err != nil {
			return errorz.Wrap(err, errorz.Errorf(fileName))
		}

		s.TopLevelResourceTypes[tlr.Name] = tlr

		for _, st := range sts {
			s.StructuredTypes[st.Name] = st
		}
	}

	return nil
}
