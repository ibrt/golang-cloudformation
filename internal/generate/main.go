package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/filez"
	"github.com/ibrt/golang-utils/ioz"

	"github.com/ibrt/golang-cloudformation/cfgenz"
	"github.com/ibrt/golang-cloudformation/cfschemaz"
	"github.com/ibrt/golang-cloudformation/cfspecz"
)

const (
	specURL   = "https://d1uauaxba7bl26.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json"
	schemaURL = "https://schema.cloudformation.us-east-1.amazonaws.com/CloudformationSchema.zip"
)

func main() {
	defer func() {
		if err := errorz.MaybeWrapRecover(recover()); err != nil {
			fmt.Println(errorz.SDump(err))
		}
	}()

	fmt.Println("Preparing temporary folder...")
	tmpDirPath := filepath.Join(".build", "generate")
	filez.MustPrepareDir(tmpDirPath, 0777)

	fmt.Println("Downloading CloudFormation spec...")
	specFilePath := filepath.Join(tmpDirPath, "CloudFormationResourceSpecification.json")
	resp, err := http.Get(specURL)
	errorz.MaybeMustWrap(err)
	errorz.Assertf(resp.StatusCode == http.StatusOK, "unexpected status code: %v", resp.StatusCode)
	filez.MustWriteFile(specFilePath, 0777, 0666, ioz.MustReadAllAndClose(resp.Body))

	fmt.Println("Parsing and validating CloudFormation spec...")
	spec, err := cfspecz.NewSpecFromBuffer(filez.MustReadFile(specFilePath), cfspecz.NewDefaultPatchManager())
	errorz.MaybeMustWrap(err)

	fmt.Println("Downloading CloudFormation schema...")
	schemaFilePath := filepath.Join(tmpDirPath, "CloudFormationSchema.zip")
	resp, err = http.Get(schemaURL)
	errorz.MaybeMustWrap(err)
	errorz.Assertf(resp.StatusCode == http.StatusOK, "unexpected status code: %v", resp.StatusCode)
	filez.MustWriteFile(schemaFilePath, 0777, 0666, ioz.MustReadAllAndClose(resp.Body))

	fmt.Println("Parsing and validating CloudFormation schema...")
	schema, err := cfschemaz.NewSchemaFromBuffer(filez.MustReadFile(schemaFilePath))
	errorz.MaybeMustWrap(err)

	g := cfgenz.NewGenerator(cfgenz.NewDefaultGeneratorOptions(), spec, schema)
	errorz.MaybeMustWrap(g.Generate(filez.MustAbs(filepath.Join("cfz", "gen"))))
}
