package cfschemaz

import (
	"archive/zip"
	"bytes"
	"io"

	"github.com/ibrt/golang-utils/errorz"
)

// Schema describes the CloudFormation JSON schema.
type Schema struct {
	RawSchemas map[string][]byte
}

// NewSchemaFromBuffer parses, patches (using the given spec patch manager), and validates a CloudFormation JSON schema
// from the given buffer (i.e. "CloudFormationSchema.zip" file).
func NewSchemaFromBuffer(buf []byte) (*Schema, error) {
	zr, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return nil, errorz.Wrap(err)
	}

	s := &Schema{
		RawSchemas: make(map[string][]byte),
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

		s.RawSchemas[f.Name] = buf
	}

	return s, nil
}
