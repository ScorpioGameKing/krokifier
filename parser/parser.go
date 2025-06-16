package parser

import (
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
)

// Used for loading the user given settings from JSON

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

// ------------------------------

// Used by krokifier to properly build output

type IStringMap = map[int][]string

type ParsedGroup struct {
	Label string
	RawWords IStringMap
	LeadingKey Keyword
	SubKeys []Keyword
}

type FileGroup struct {
	Groups map[string]ParsedGroup
}

// ------------------------------

// JSON File Loading

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

// ------------------------------

func ParseFile(lines []string, syntax *SyntaxFile, file *FileGroup) (error) {
	file.Groups = map[string]ParsedGroup{}

	fmt.Printf("\nPrinting Syntax File Diagram\n")

	for i, group := range syntax.Diagram {
		fmt.Printf("%d : %+v\n", i, group)
	}

	fmt.Printf("\nParsing Input File\n")

	for _, line := range lines {
		//fmt.Printf("%d : %s\n", i, line)
		
		words := strings.Split(strings.TrimLeft(line, " "), " ")
		
		for _, group := range syntax.Diagram {
			for _, kword := range group.Keywords {
				switch words[0] {
					case kword.Word:
						fmt.Printf("Found Group: %s\n", group.Label)
						
						if file_group, ok := file.Groups[group.Label]; ok {
							file_group.RawWords[len(file_group.RawWords)] = words
						} else {
							var new_group ParsedGroup
							var raw_words IStringMap

							raw_words = IStringMap{}
							raw_words[0] = words

							new_group.Label = group.Label
							new_group.RawWords = raw_words
							new_group.LeadingKey = kword
							
							if len(group.Keywords) > 1 {
								new_group.SubKeys = group.Keywords[1:]
							} else {
								var no_keys []Keyword
								new_group.SubKeys = no_keys
							}
							
							file.Groups[group.Label] = new_group
						}
				}
			}
		}
	}

	fmt.Printf("\nFile Out:\n%+v", file)
	return nil
}
