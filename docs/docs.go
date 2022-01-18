package main

import (
	"fmt"
	"github.com/cloudquery/cq-provider-msgraph/resources/provider"
	"io/ioutil"
	"os"
	"path"

	"github.com/cloudquery/cq-provider-sdk/provider/docs"
)

func main() {
	outputPath := "./docs"
	dir, err := ioutil.ReadDir(outputPath + "/tables")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
		os.Exit(1)
	}
	for _, d := range dir {
		if err := os.RemoveAll(path.Join([]string{outputPath + "/tables", d.Name()}...)); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
			os.Exit(1)
		}

	}

	if err = docs.GenerateDocs(provider.Provider(), outputPath); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
	}
}
