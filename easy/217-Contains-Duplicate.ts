/**
 * @link https://leetcode.com/problems/contains-duplicate/
 */


function containsDuplicate(nums: number[]): boolean {

  const set = new Set<number>(nums)
  console.log(set.size, nums.length )
  return set.size !== nums.length
}
