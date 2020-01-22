package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/reconquest/lstags/pkg/gotags"
	"github.com/reconquest/lstags/pkg/lstags"
	"github.com/reconquest/pkg/log"

	"github.com/docopt/docopt-go"
)

var version = "[manual build]"

var usage = `lstags - list tags symbols

Usage:
  lstags -h | --help
  lstags [options]

Options:
  -x --suffix <ext>   Filter files by specified suffix. [default: .go]
  -m --method <name>  Use the following tags method.
                       Supported only the following values:
                       * gotags
                       [default: gotags]
  -h --help           Show this help.
  --version           Show version.
`

func main() {
	args, err := docopt.ParseArgs(usage, nil, "lstags "+version)
	if err != nil {
		panic(err)
	}

	var opts struct {
		Suffix string `docopt:"--suffix"`
		Method string `docopt:"--method"`
	}

	err = args.Bind(&opts)
	if err != nil {
		log.Fatalf(err, "unable to bind arguments")
	}

	files, err := walkDir(".", opts.Suffix)
	if err != nil {
		log.Fatalf(err, "unable to walk directory")
	}

	var tags []lstags.Tag
	for _, filename := range files {
		if opts.Method == "gotags" {
			filetags, err := gotags.Parse(filename)
			if err != nil {
				log.Errorf(err, "unable to parse %s", filename)
			}

			tags = append(tags, filetags...)
		}
	}

	for _, tag := range tags {
		fmt.Printf(
			"%s:%d:%d:%s\n",
			tag.GetFilename(),
			tag.GetLine(),
			tag.GetColumn(),
			tag.GetSignature(),
		)
	}
}

func walkDir(dir string, suffix string) ([]string, error) {
	var names []string
	err := filepath.Walk(dir, func(path string, finfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, suffix) && !finfo.IsDir() {
			names = append(names, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return names, nil
}
