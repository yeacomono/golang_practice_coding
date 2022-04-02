package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
func main() {
	//libreria que permite usar input de la terminal os<.stdin
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var operacion string= scanner.Text()
	fmt.Println(operacion)
	 valores := strings.Split(operacion, "+")

	 valorOne, _  := strconv.Atoi( valores[0])
	 valorTwo, _  :=  strconv.Atoi( valores[1])
	
   // resultOperation := (valorOne+valorTwo)
    fmt.Println(valorOne+valorTwo)
}