// https://leetcode.com/problems/unique-morse-code-words/

func uniqueMorseRepresentations(words []string) int {
    morseList := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    
    uniqueWord := make(map[string]int)
    for _, word := range words {
        var newWord string
        for _, alpha := range []rune(word) {
            newWord = newWord + morseList[alpha - 97]
        }
        uniqueWord[newWord] = 1
    }
    return len(uniqueWord)
}