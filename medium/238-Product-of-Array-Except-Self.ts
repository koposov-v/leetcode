/**
 * @link https://leetcode.com/problems/product-of-array-except-self/description/
 * @param nums
 */
function productExceptSelf(nums: number[]): number[] {
  let post = 1
  let prefix = 1

  const res: number[] = []

  for (let i = 0; i < nums.length; i++) {
    res[i] = prefix
    prefix *= nums[i]
  }

  for (let i = nums.length - 1; i >= 0; i--) {
    res[i] *= post
    post *= nums[i]
  }

  return res
}

const nums = [1, 2, 3, 4]
productExceptSelf(nums)
// Output: [24,12,8,6]
export {}
