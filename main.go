package main

import (
    "bytes"
    "compress/zlib"
    //"encoding/base64"
    "fmt"
	"flag"
    "bufio"
    "os"
	"io/ioutil"
	"encoding/json"
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

type SyntaxFile struct {
	UML_Type string
	Language string
	Diagram []UMLGroups
}

type LangFile struct {
	Keywords []Keyword
}

type Keyword struct {
	Word string
	Extends string
}

type UMLGroups struct {
	Label string
}

func LoadSyntaxFile(path string, syntax *SyntaxFile) (error) {
    file, err := ioutil.ReadFile(path)

    if err != nil {
        return err
    }

    return json.Unmarshal(file, syntax)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
    var buf bytes.Buffer
	var syntax SyntaxFile 
    var lines []string

    file_path := flag.String("path", "", "The file to parse")
	syntax_path := flag.String("syntax", "./res/syntax/default_py_blockdiag.json", "The syntax file used to generate UML")

	flag.Parse()
	
    lines, err = ReadFile(*file_path)

	CheckErr(err)

	err = LoadSyntaxFile(*syntax_path, &syntax)

	CheckErr(err)

	fmt.Print(syntax.UML_Type)

	for i, line := range lines {
		_, err := Encode(line, &buf)

		CheckErr(err)

    	//result := base64.URLEncoding.EncodeToString(buf.Bytes())

    	fmt.Printf("%d : %s\n", i, line)
    }
}
