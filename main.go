package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите текст для анализа корректности расстановки скобок: ")
	text, _ := reader.ReadString('\n')
	str := StringToParse(text)
	err := str.Parse()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Проблем со скобками не обнаружено")
	}
}
