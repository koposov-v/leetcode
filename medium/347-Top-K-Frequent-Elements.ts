/**
 * @link https://leetcode.com/problems/top-k-frequent-elements/description/
 */

function topKFrequent(nums: number[], k: number): number[] {

  const hashTable = new Map<number, number>()

  for (let i = 0; i < nums.length; i++) {
    let count = hashTable.get(nums[i])
    const value = (count ?? 0) + 1
    hashTable.set(nums[i], value)
  }

  const freq: Record<number, number[]> & Object = {}

  for (const [key, value] of hashTable.entries()) {
    const values = freq[value] ?? []
    freq[value] = values.concat(key)
  }
  const res: number[] = []

  for (let i = Object.keys(nums).length; i >= 0; i--) {
    const value = freq[i]
    if (!value) continue;
    res.push(...value)
    const countRes = res.length

    if (countRes == k) break;

    if (countRes > k) {
      deleteLastElement(res, k)
    }
  }


  return res
}

const deleteLastElement = (arr: number[], k: number) => {
  arr.pop()
  if (arr.length > k) {
    deleteLastElement(arr, k)
  }
}


// const nums = [1, 1, 1, 2, 2, 3], k = 2
const nums1 = [1, 2, 5, 5, 5], k = 2
const nums2 = [1], k2 = 1
topKFrequent(nums2, k2)
// Output: [1,2]