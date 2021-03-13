package config

import (
	"errors"
	"fmt"

	"github.com/valonekowd/togo/util/helper"
)

var (
	ErrEmptyFileName       = errors.New("empty file name")
	ErrUnsupportedFileExt  = errors.New("unsupported file ext")
	ErrNoFilePathsProvided = errors.New("no file paths provided")
)

var SupportedFileExts = []string{"json", "yml"}

type File struct {
	Name  string
	Type  string
	Paths []string
}

func NewFile(name, fileType string, paths []string) (*File, error) {
	f := &File{
		Name:  name,
		Type:  fileType,
		Paths: paths,
	}

	if f.Name == "" {
		return nil, fmt.Errorf("new config file: %w", ErrEmptyFileName)
	}

	if !helper.StringInSlice(f.Type, SupportedFileExts) {
		return nil, fmt.Errorf("new config file: %w", ErrUnsupportedFileExt)
	}

	for _, p := range paths {
		if p != "" {
			return f, nil
		}
	}

	return nil, fmt.Errorf("new config file: %w", ErrNoFilePathsProvided)
}
