//import "slices"

func isAnagram(s string, t string) bool {
// Create frequency maps for both strings
    freqS := make(map[rune]int)
    freqT := make(map[rune]int)

    // Count frequencies of characters in s
    for _, char := range s {
        freqS[char]++
    }

    // Count frequencies of characters in t
    for _, char := range t {
        freqT[char]++
    }

    // Check if the frequency maps are equal
    if len(freqS) != len(freqT) {
        return false
    }

    for key, val := range freqS {
        if freqT[key] != val {
            return false
        }
    }

    return true
}
