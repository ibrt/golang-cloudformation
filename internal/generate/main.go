package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/filez"
	"github.com/ibrt/golang-utils/ioz"

	"github.com/ibrt/golang-cloudformation/cfspecz"
)

const (
	specURL = "https://d1uauaxba7bl26.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json"
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
	_, err = cfspecz.NewSpecFromBuffer(filez.MustReadFile(specFilePath), cfspecz.NewDefaultSpecPatchManager())
	errorz.MaybeMustWrap(err)
}
