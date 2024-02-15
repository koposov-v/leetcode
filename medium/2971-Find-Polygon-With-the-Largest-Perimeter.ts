/**
 * @link https://leetcode.com/problems/find-polygon-with-the-largest-perimeter/?envType=daily-question&envId=2024-02-15
 * @param nums
 */
 const largestPerimeter = function(nums: number[]) {
  nums.sort((a, b) => a - b);
  let prefix = new Array(nums.length + 1).fill(0);
  let i = 1;

  while (i < prefix.length) {
    prefix[i] = prefix[i - 1] + nums[i - 1];
    i++;
  }
  console.log(prefix)

  i = nums.length - 1;

  while (i > 1) {
    if (prefix[i] > nums[i]) {
      return prefix[i + 2];
    }
    i--;
  }

  return -1;
};

const nums = [1, 12, 1, 2, 5, 50, 3]
// const nums = [1, 2, 1]
// const nums = [5, 5, 5]

console.log(largestPerimeter(nums))

export {}