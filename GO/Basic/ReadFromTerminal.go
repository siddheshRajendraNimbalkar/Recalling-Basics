package basic

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFromTerminal() {

	var rollNo int
	fmt.Println("Enter the Name :")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("ERROR:", err)
	}
	name = strings.TrimSpace(name)
	fmt.Println("Enter the RollNo. :")

	fmt.Scanln(&rollNo)
	fmt.Println(name, " ", rollNo)
}
