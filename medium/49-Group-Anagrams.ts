/**
 * @link https://leetcode.com/problems/group-anagrams/description/
 */


function groupAnagrams(strs: string[]): string[][] {
  const hashTable = new Map<string, number>()
  const newArr: any[] = []

  strs.forEach(str => {
    const strLol = str.split('').sort().join()
    const index = hashTable.get(strLol)
    if (index !== undefined) {
      newArr[index] = newArr[index].concat(str)
    } else {
      newArr.push([str])
      hashTable.set(strLol, newArr.length - 1)
    }

  })

  return newArr
}


const strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
groupAnagrams(strs)
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]
export {}