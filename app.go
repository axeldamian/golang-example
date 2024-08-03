package main
import "fmt"

func translado(i, j int) (int , int) {
	return i + 1 , j + 1
}

func main() {
  x , y := translado( 1 , 2 )
	s := fmt.Sprintf("translado is %v , %v", x , y )
	fmt.Println(s)
}
