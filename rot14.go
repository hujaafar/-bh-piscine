package piscine

func Rot14(s string) string {
	s1:=[]rune(s)
	for i, chr:=range s{
		if chr >='a' && chr<='z' || chr >='A' && chr <='Z'{
			s1[i]+=14
		}
	}
	return string(s1)
}