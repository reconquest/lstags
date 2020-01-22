package gotags

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	lib "github.com/kovetskiy/gotags/pkg/gotags"
	"github.com/reconquest/karma-go"
	"github.com/reconquest/lstags/pkg/lstags"
)

type Tag struct {
	lib.Tag
}

func (tag *Tag) GetColumn() int {
	value, err := strconv.Atoi(tag.Fields["column"])
	if err != nil {
		panic(karma.Format(
			err,
			"column is not a number: %#v", tag,
		))
	}

	return value
}

func (tag *Tag) GetLine() int {
	value, err := strconv.Atoi(tag.Fields["line"])
	if err != nil {
		panic(karma.Format(
			err,
			"column is not a number: %#v", tag,
		))
	}

	return value
}

func (tag *Tag) GetFilename() string {
	return tag.File
}

func (tag *Tag) GetSignature() string {
	wrap := func(x string) string {
		if strings.Contains(x, ",") {
			return "(" + x + ")"
		}

		return x
	}

	switch tag.Type {
	case lib.Method:
		return fmt.Sprintf(
			"(%s) %s%s %s",
			tag.Fields["ctype"],
			tag.Name,
			tag.Fields["signature"],
			wrap(tag.Fields["type"]),
		)

	case lib.Function:
		return fmt.Sprintf(
			"%s%s %s",
			tag.Name,
			tag.Fields["signature"],
			wrap(tag.Fields["type"]),
		)
	}

	return ""
}

func Parse(filename string) ([]lstags.Tag, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	tags, err := lib.Parse(filename, true, cwd)
	if err != nil {
		return nil, err
	}

	var result []lstags.Tag
	for _, tag := range tags {
		if tag.Type != lib.Function && tag.Type != lib.Method {
			continue
		}
		fmt.Fprintf(os.Stderr, "XXXXXX gotags.go:64 tag: %#v\n", tag)
		result = append(result, &Tag{tag})
	}

	return result, nil
}
