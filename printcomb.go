package main 

import "github.com/01-edu/z01"

func printcomb()  {
	for x:='0'; x<= 7; x++ {
		for y:= '1' + x; y <=8; y++{
			for z:= '1' + y; z<= 9 z++ {

				z01.PrintRune(x)
				z01.PrintRune(y)
				z01.PrintRune(z)
				z01.PrintRune(',')

				if x=='7' && y =='8' && z == '9'{
					z01.print
				}

			}
		}
	}
}

func main(){


}