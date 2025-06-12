package main

import (
    "bytes"
    "compress/zlib"
    // "encoding/base64"
    "fmt"
    "bufio"
    "os"
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

func main() {
    var buf bytes.Buffer
    var lines []string
    file_path := "./tests/_nirn_weaver.py"
    lines, err := ReadFile(file_path)

    if err != nil {
    	panic(err)
    }
    
    for i, line := range lines {
    	_, err := Encode(line, &buf)

		if err != nil {
	    	panic(err)
	    }

    	// result := base64.URLEncoding.EncodeToString(buf.Bytes())

    	fmt.Printf("%d : %s\n", i, line)
    }
}
