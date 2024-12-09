package cfschemaz

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ibrt/golang-utils/errorz"
)

// RootSchema describes the CloudFormation JSON schema.
type RootSchema struct {
	ParsedSchemasByFileName map[string]*ParsedSchema
}

// NewRootSchemaFromBuffer parses and validates a CloudFormation JSON schema from the given buffer
// (i.e. "CloudFormationSchema.zip" file).
func NewRootSchemaFromBuffer(buf []byte) (*RootSchema, error) {
	zr, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return nil, errorz.Wrap(err)
	}

	s := &RootSchema{
		ParsedSchemasByFileName: make(map[string]*ParsedSchema),
	}

	for _, f := range zr.File {
		pr, err := newParsedSchemaFromFile(f)
		if err != nil {
			return nil, errorz.Wrap(err)
		}

		s.ParsedSchemasByFileName[f.Name] = pr
	}

	return s, nil
}

func newParsedSchemaFromFile(f *zip.File) (ps *ParsedSchema, err error) {
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
