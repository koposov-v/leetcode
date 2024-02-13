/**
 * @link https://leetcode.com/problems/top-k-frequent-elements/description/
 */

function topKFrequent(nums: number[], k: number): number[] {

  const hashTable = new Map<number, number>()

  for (let i = 0; i < nums.length; i++) {
    let count = hashTable.get(nums[i])
    hashTable.set(nums[i], (count ?? 0) + 1)
  }

  const newArr: number[] = []
  hashTable.forEach((value, key) => {
    if (value >= k) {
      newArr.push(key)
    }
  })

  console.log(newArr)
  return newArr
}
// const nums = [1, 1, 1, 2, 2, 3], k = 2
const nums = [1,2], k = 2
topKFrequent(nums, k)
// Output: [1,2]

export {}
