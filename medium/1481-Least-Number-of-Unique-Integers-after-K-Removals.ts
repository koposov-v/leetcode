/**
 * @link
 * https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/description/?envType=daily-question&envId=2024-02-16
 */

function findLeastNumOfUniqueInts(arr: number[], k: number): number {

  const map = new Map<number, number>()
  arr.sort((a, b) => a - b)

  for (let i = 0; i < arr.length; i++) {
    map.set(arr[i], (map.get(arr[i]) ?? 0) + 1)
  }


  const sortedAmounts = Array.from(map.values()).sort((a, b) => a - b)
  let unique = sortedAmounts.length;

  for (let i = 0; i < sortedAmounts.length; i++) {
    if (k >= sortedAmounts[i]) {
      k -= sortedAmounts[i];
      unique--;
    } else {
      break;
    }
  }

  return unique;

}


// const arr = [5, 5, 4], k = 1
 const arr = [4,4,4,3,1,1,3,3,2], k = 3
findLeastNumOfUniqueInts(arr, k)


export {}