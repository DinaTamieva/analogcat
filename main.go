package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func getContentFile(filePath string) []byte {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer f.Close()

	fStat, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	buf := make([]byte, fStat.Size())
	if _, err := io.ReadFull(f, buf); err != nil {
		fmt.Println(err)
		return nil
	}
	return buf
}

func writeFile(fileName string, content1, content2 []byte) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	writer := bufio.NewWriter(f)
	writer.Write(content1)
	writer.WriteString("\n")
	writer.Write(content2)
	writer.Flush()

}

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Println()
		fmt.Println("Аналог программы cat")
		fmt.Println()
		fmt.Println("Необходимо передать аргументы!")
		fmt.Println()
		fmt.Println("Для проверки задания выполните go run main.go first.txt second.txt \n или \n go run main.go first.txt second.txt result.txt \n или \n go run main.go first.txt")
	case 2: //один аргумент
		filePath := os.Args[1]
		fmt.Println(filePath)
		text := getContentFile(filePath)
		fmt.Println(text)
	case 3: //два агумента
		filePath1 := os.Args[1]
		text1 := getContentFile(filePath1)
		filePath2 := os.Args[2]
		text2 := getContentFile(filePath2)
		unionstring := strings.Join([]string{string(text1), string(text2)}, "\n")
		fmt.Println("Содержимое двух файлов: \n", unionstring)
	case 4: // 3 аргумента
		filePath1 := os.Args[1]
		text1 := getContentFile(filePath1)
		filePath2 := os.Args[2]
		text2 := getContentFile(filePath2)
		fmt.Printf("Содержимое %s и %s записано в файл %s", filePath1, filePath2, os.Args[3])

		writeFile(string(os.Args[3]), text1, text2)
	default:
		fmt.Println("Too many arguments")
	}
}
