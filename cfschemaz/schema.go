package cfschemaz

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ibrt/golang-utils/errorz"
)

// Schema describes the CloudFormation JSON schema.
type Schema struct {
	UnprocessedTopLevelResourcesByFileName map[string]*UnprocessedTopLevelResource
	ResourceTypes                          map[string]*Type
	StructuredTypes                        map[string]*Type
}

// NewSchemaFromBuffer parses and validates a CloudFormation JSON schema from the given buffer
// (i.e. "CloudFormationSchema.zip" file).
func NewSchemaFromBuffer(buf []byte) (*Schema, error) {
	zr, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return nil, errorz.Wrap(err)
	}

	s := &Schema{
		UnprocessedTopLevelResourcesByFileName: make(map[string]*UnprocessedTopLevelResource),
	}

	for _, f := range zr.File {
		pr, err := parseUnprocessedTopLevelResource(f)
		if err != nil {
			return nil, errorz.Wrap(err)
		}

		s.UnprocessedTopLevelResourcesByFileName[f.Name] = pr
	}

	s.preProcess()
	return s, nil
}

func parseUnprocessedTopLevelResource(f *zip.File) (ps *UnprocessedTopLevelResource, err error) {
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

	if err := d.Decode(&ps); err != nil {
		return nil, errorz.Wrap(err, fmt.Errorf(f.Name))
	}

	return ps, nil
}

func (s *Schema) preProcess() {
	for _, ur := range s.UnprocessedTopLevelResourcesByFileName {
		s.ResourceTypes[ur.TypeName] = &Type{
			IsTopLevelResourceType: true,
			Name:                   ur.TypeName,
			Description:            ur.Description,
		}
	}
}
