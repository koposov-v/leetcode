function firstPalindrome(words: string[]): string {
  for (const word of words) {
    if (isPalindrome(word)) {
      return word
    }
  }

  return ''
}

const isPalindrome = (word: string): boolean => {
  const reverseWord = word.split('').reverse().join('')
  return reverseWord === word
}


const words = ["abc", "car", "ada", "racecar", "cool"]
// Output: "ada"
const a = firstPalindrome(words)
console.log(a)

export {}