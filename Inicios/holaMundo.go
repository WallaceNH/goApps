package Inicios

import (
	"fmt"
	"strconv"
)

// hola mundo
func main0() {

	//fmt.Println("Destructor de mundos")
	i := 0
	for i >= 0 {
		i++
		if i >= 60 {
			i = -1
			fmt.Println("aqui muere todo")
		} else {
			fmt.Println("contador en : " + strconv.Itoa(i))
		}
	}
}
