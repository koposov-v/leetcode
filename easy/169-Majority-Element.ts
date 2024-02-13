function majorityElement(nums: number[]): number {
  let candidate: number = 0;
  let count = 0;

  for (const num of nums) {
    if (count === 0) {
      candidate = num;
    }

    count += (num === candidate) ? 1 : -1
  }

  return candidate
}

export {}