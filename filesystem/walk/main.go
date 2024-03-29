package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type config struct {
	// extension to filter out
	ext string
	// min file size
	size int64
	// list files
	list bool
}

func main() {
	// Parsing command line flags
	root := flag.String("root", ".", "Root directory to start")
	//Action Options
	list := flag.Bool("list", false, "List files only")
	// Filter Options
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minium file size")
	flag.Parse()

	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
	}
	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func run(root string, out io.Writer, cfg config) error {
	return filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filterOut(path, cfg.ext, cfg.size, info) {
				return nil
			}
			// if list was explicitly set, don't do anything else
			if cfg.list {
				return listFile(path, out)
			}
			return listFile(path, out)
		})

}
