package main

import (
    "bytes"
    "compress/zlib"
    //"encoding/base64"
    "fmt"
	"flag"
    "bufio"
    "os"
	"strings"
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

type UMLGroups struct {
	Label string
	Keywords []Keyword
}

type Keyword struct {
	Word string
	Extension string
	Extends bool 
	Recursive bool
	Inherits bool
}

type UMLTypes struct {
	Connecting string
}

func LoadSyntaxFile(path string, syntax *SyntaxFile) (error) {
    file, err := ioutil.ReadFile(path)

    if err != nil {
        return err
    }

    return json.Unmarshal(file, syntax)
}

func LoadUMLTypes(uml_type string, umltypes *UMLTypes) (error) {
	var path string
	switch uml_type {
		case "blockdiag":
			path = "./res/language/blockdiag.json"
	}
	file, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	return json.Unmarshal(file, umltypes)
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
	var umltypes UMLTypes
    var lines []string
	var output []string

    file_path := flag.String("path", "", "The file to parse")
	syntax_path := flag.String("syntax", "./res/syntax/default_py_blockdiag.json", "The syntax file used to generate UML")

	flag.Parse()
	
    lines, err = ReadFile(*file_path)

	CheckErr(err)

	err = LoadSyntaxFile(*syntax_path, &syntax)

	CheckErr(err)

	err = LoadUMLTypes(syntax.UML_Type, &umltypes)

	CheckErr(err)

	fmt.Printf("\nPrinting Syntax File Diagram\n")
	for i, group := range syntax.Diagram {
		fmt.Printf("%d : %+v\n", i, group)
	}

	fmt.Printf("\nParsing Input File\n")
	for _, line := range lines {
		_, err := Encode(line, &buf)

		CheckErr(err)

    	//result := base64.URLEncoding.EncodeToString(buf.Bytes())

		words := strings.Split(strings.TrimLeft(line, " "), " ")
		for _, group := range syntax.Diagram {
			for _, kword := range group.Keywords {
				switch words[0] {
					case kword.Word:
						var parsed_string bytes.Buffer
						var recursive_string []string
						
						if !kword.Recursive {
							if kword.Extension == "" {
								parsed_string.WriteString(words[1])
							} else {
								parsed_string.WriteString(strings.ReplaceAll(words[1], kword.Extension, umltypes.Connecting))		
							}	
						} else {
							recursive_string = strings.Split(words[1], kword.Extension)
							for _, recursive_word := range recursive_string {
								parsed_string.WriteString(recursive_word)
							}
						}

						fmt.Printf("\nFOUND KEYWORD: %s FROM GROUP: %+v IN LINE: %v WILL DISPLAY AS: %s\n", kword.Word, group, words, parsed_string.String())

						if len(group.Keywords) > 1 {
							sub_words := words[2:]
							for i, sub_word := range sub_words {
								for _, sub_kword := range group.Keywords[1:] {
									switch sub_word {
										case sub_kword.Word:
											if !sub_kword.Recursive {
												if sub_kword.Extension == "" {
													parsed_string.WriteString(fmt.Sprintf("%s%s", umltypes.Connecting, sub_words[i+1]))
												} else {
													parsed_string.WriteString(fmt.Sprintf("%s%s", umltypes.Connecting, strings.ReplaceAll(sub_words[i+1], sub_kword.Extension, " -> ")))		
												}	
											} else {
												current_string := parsed_string
												for n, recursive_word := range sub_words[i+1:] {
													if n >= 1 {
														parsed_string.WriteString(fmt.Sprintf("%s%s%s\n", current_string.String(), umltypes.Connecting, recursive_word))
													} else {
														parsed_string.WriteString(fmt.Sprintf("%s%s\n", umltypes.Connecting, recursive_word))
													}
												}

											}
											fmt.Printf("\nFOUND KEYWORD: %s FROM GROUP: %+v IN LINE: %v WILL DISPLAY AS: %s\n", sub_kword.Word, group, sub_words, sub_words[i+1])
									}
							}
						}
					}
					output = append(output, fmt.Sprintf("%s\n", parsed_string.String())) 
				}
			}
		}
    	//fmt.Printf("%d : %s\n", i, line)
    }
	fmt.Printf("OUTPUT: %v", output)
}
