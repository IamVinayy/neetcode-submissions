import "slices"

func isAnagram(s string, t string) bool {
  //convert the string to slices of runes
	sRune := []rune(s)
	tRune := []rune(t)
	//sort them 
	slices.Sort(sRune)
	slices.Sort(tRune)
	
  //compare them after sorting if they are equal. 
	return string(sRune) == string(tRune)
}
