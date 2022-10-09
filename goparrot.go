package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	parseFile(parse())
}

func parse() (file_name string, line_number int, p string) {
	flag.StringVar(&p, "p", "", "")
	flag.Parse()
	file_name, b := os.LookupEnv("GOFILE")
	if !b {
		return
	}
	line, b2 := os.LookupEnv("GOLINE")
	if !b2 {
		return
	}
	line_number, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return file_name, line_number, p
}

func parseFile(filename string, line int, args string) {
	file_set := token.NewFileSet()
	file, err := parser.ParseFile(file_set, filename, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	for _, d := range file.Decls {
		if func_decl, ok := d.(*ast.FuncDecl); ok {
			doc := func_decl.Doc
			if doc == nil {
				continue
			}
			token_file := file_set.File(doc.Pos())
			if token_file.Line(doc.Pos()) == line {
				buf := bytes.NewBuffer(make([]byte, 0, func_decl.End()-func_decl.Type.Pos()+1))
				func_decl.Doc = nil
				printer.Fprint(buf, file_set, func_decl)
				print(buf.String(), args)
			}
		}
	}
}

func print(s string, args string) {
	//["Time;Open;High", "v[0];v[1];v[2]"]
	rs := strings.Split(args, "&&")
	column_length := len(rs)
	type row []string
	type rows []row
	matrix := make(rows, column_length)
	// str := s
	for i := range rs {
		matrix[i] = strings.Split(rs[i], ";")
	}
	// matrix[0] = ["Time","Open","High"]
	loop_count := len(matrix[0])
	for i := 0; i < loop_count; i++ {
		tar := s
		for j := 0; j < len(matrix); j++ {
			if old, new := matrix[j][0], matrix[j][i]; new != "" && new != old {
				tar = ReplaceAllByCaseSensitive(tar, old, new)
			}
		}
		fmt.Println(tar)
	}
}

func ReplaceAllByCaseSensitive(s, old, new string) string {
	tar := s
	tar = strings.ReplaceAll(tar, old, new)
	if unicode.IsLower([]rune(old)[0]) {
		tar = strings.ReplaceAll(tar, strings.ToTitle(old), strings.ToTitle(new))
	} else {
		old_rune := []rune(old)
		old_rune[0] = unicode.ToLower(old_rune[0])
		new_rune := []rune(new)
		new_rune[0] = unicode.ToLower(new_rune[0])
		tar = strings.ReplaceAll(tar, string(old_rune), string(new_rune))
	}
	return tar
}
