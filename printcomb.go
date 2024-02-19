package piscine
import "github.com/01-edu/z01"

func PrintComb(){
for i:='0' ; i<='9' ; i++{
	for b:=i+1 ; b<='9' ; b++{
		for c:=b+1 ; c<='9' ; c++{
z01.PrintRune(i)
z01.PrintRune(b)
z01.PrintRune(c)
if i=='7' && b=='8' && c=='9' {
	z01.PrintRune('\n')
}else{
z01.PrintRune(',')
z01.PrintRune(' ')
}
		}

	}
}

}