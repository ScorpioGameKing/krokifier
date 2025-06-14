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
	diagram []UMLGroups
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

func LoadSyntaxFile(path string) (*SyntaxFile, error) {
    file, err := ioutil.ReadFile(path)

    if err != nil {
       return nil, err
    }

	var syntax SyntaxFile 
	err = json.Unmarshal(file, &syntax)
	
	if err != nil {
		return nil, err
	}

	return &syntax, nil
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
    // var buf bytes.Buffer
	var syntax SyntaxFile 
    // var lines []string

    // file_path := flag.String("path", "./tests/_nirn_weaver.py", "The file to parse")
	syntax_path := flag.String("syntax", "./tests/default_py_blockdiag.json", "The syntax file used to generate UML")

	flag.Parse()
	
    // lines, err = ReadFile(*file_path)

	// CheckErr(err)

	syntax, err = LoadSyntaxFile(*syntax_path)

	CheckErr(err)

	fmt.Print(syntax)

	// for i, line := range lines {
		// _, err := Encode(line, &buf)

		// CheckErr(err)

    	// result := base64.URLEncoding.EncodeToString(buf.Bytes())

    	// fmt.Printf("%d : %s\n", i, line)
    // }
}
