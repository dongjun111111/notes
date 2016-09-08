package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

var (
	begin int
	step  int
	width int
)

func init() {
	flag.IntVar(&begin, "b", 0, "iterator to begin")
	flag.IntVar(&step, "s", 1, "add each time to iterator")
	flag.IntVar(&width, "w", 1, "width translate iterator to string")
	flag.Parse()
}

func main() {
	if flag.NArg() < 2 {
		fmt.Println("usage: executable match-pattern name-pattern")
		return
	}

	reg, err := regexp.Compile(flag.Arg(0))
	if err != nil {
		fmt.Println("syntax error in match-pattern")
		return
	}
	tgt, err := template.New("only").Parse(strings.Join(flag.Args()[1:], " "))
	if err != nil {
		fmt.Println("syntax error in name-pattern")
		return
	}
	buf := bytes.NewBuffer(nil)

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if info.IsDir() {
			if path == "" || path == "." {
				return nil
			}
			return filepath.SkipDir
		}
		ans := reg.FindStringSubmatch(path)
		if ans == nil {
			return nil
		}
		ans[0] = fmt.Sprintf("%*d", width, begin)
		begin += step
		err = tgt.Execute(buf, ans)
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = os.Rename(path, buf.String())
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(buf.String())
		buf.Reset()
		return nil
	})
}
