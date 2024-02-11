package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){

	// membuat agar user bisa melakukan fungsi input di command line
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nilai n: ")
	input, _ := reader.ReadString('\n') // membaca input hingga newline
	input = strings.TrimSpace(input) //menghapus spasi dan newline

	// konversi string ke number
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error: Input bukan angka valid")
		return
	}

	for i := 1; i<=n; i++ {
		var isSquare, isCube = false, false
		// fmt.Printf("Iterasi  ke -%d\n", i)
		for j:= 1; j*j<=i; j++{
			if j*j == i{
				isSquare = true
				break;
			}
		}
		
		for j:= 1; j*j*j <= i; j++{
			if j*j*j == i{
				isCube = true;
				break
			}
		}

		if isSquare && isCube {
			fmt.Println("SquareCube")
		}else if isSquare{
			fmt.Println("Square")
		}else if isCube {
			fmt.Println("Cube")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
	

}