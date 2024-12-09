package cfschemaz

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/ibrt/golang-utils/errorz"
)

// Schema describes the CloudFormation JSON schema.
type Schema struct {
	RawSchemas      map[string][]byte
	ResourceSchemas map[string]*Resource
}

// NewSchemaFromBuffer parses and validates a CloudFormation JSON schema from the given buffer
// (i.e. "CloudFormationSchema.zip" file).
func NewSchemaFromBuffer(buf []byte) (*Schema, error) {
	zr, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return nil, errorz.Wrap(err)
	}

	s := &Schema{
		RawSchemas:      make(map[string][]byte),
		ResourceSchemas: make(map[string]*Resource),
	}

	for _, f := range zr.File {
		fr, err := f.Open()
		if err != nil {
			return nil, errorz.Wrap(err)
		}

		buf, err := io.ReadAll(fr)
		if err != nil {
			return nil, errorz.Wrap(err)
		}

		d := json.NewDecoder(bytes.NewBuffer(buf))
		d.DisallowUnknownFields()

		var r *Resource

		if err := d.Decode(&r); err != nil {
			return nil, errorz.Wrap(err, fmt.Errorf(f.Name))
		}

		s.RawSchemas[f.Name] = buf
		s.ResourceSchemas[f.Name] = r
	}

	return s, nil
}
