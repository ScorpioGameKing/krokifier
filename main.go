package main

import (
    "bytes"
    "compress/zlib"
    //"encoding/base64"
	"flag"
    "bufio"
    "os"
	"github.com/ScorpioGameKing/krokifier/parser"
)

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func Encode(input string, buf *bytes.Buffer) (int, error) {
    writer, err := zlib.NewWriterLevel(buf, zlib.BestCompression)

    if err != nil {
        return 0, err
    }

    defer writer.Close()

    n, err := writer.Write([]byte(input))
    if err != nil {
        return 0, err
    }

    return n, nil
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	var syntax parser.SyntaxFile 
	var umltypes parser.UMLTypes
    var lines []string

    file_path := flag.String("path", "", "The file to parse")
	syntax_path := flag.String("syntax", "", "The syntax file used to generate UML")

	flag.Parse()
	
    lines, err = ReadFile(*file_path)

	CheckErr(err)

	err = parser.LoadSyntaxFile(*syntax_path, &syntax)

	CheckErr(err)

	err = parser.LoadUMLTypes(syntax.UML_Type, &umltypes)

	CheckErr(err)

	_, err = parser.ParseFile(lines, &syntax, &umltypes)

}
