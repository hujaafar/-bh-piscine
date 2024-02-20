a="$(ls -lR | Egrep -c '^-|^d')"
((a++))
echo $a