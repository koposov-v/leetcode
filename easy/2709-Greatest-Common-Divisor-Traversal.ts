/**
 * @link https://leetcode.com/problems/greatest-common-divisor-traversal
 */


function canTraverseAllPairs(nums: number[]): boolean {
  const n = nums.length;
  const graph = new Map<number, number[]>();

  // Function to add an edge in the graph
  function addEdge(u: number, v: number) {
    if (!graph.has(u)) graph.set(u, []);
    graph.get(u)!.push(v);
  }

  // Function to get all prime factors
  function getPrimeFactors(num: number): number[] {
    const factors = new Set<number>();
    for (let i = 2; i * i <= num; i++) {
      while (num % i === 0) {
        factors.add(i);
        num /= i;
      }
    }
    if (num > 1) factors.add(num);
    return Array.from(factors);
  }

  // Creating connections based on prime factors
  const primeToIndex = new Map<number, number>();
  for (let i = 0; i < n; i++) {
    const factors = getPrimeFactors(nums[i]);
    for (const factor of factors) {
      if (primeToIndex.has(factor)) {
        addEdge(primeToIndex.get(factor)!, i);
        addEdge(i, primeToIndex.get(factor)!);
      }
      primeToIndex.set(factor, i);
    }
  }

  // Depth-First Search to check connectivity
  const visited = new Array(n).fill(false);
  function dfs(node: number) {
    visited[node] = true;
    for (const neighbor of graph.get(node) || []) {
      if (!visited[neighbor]) {
        dfs(neighbor);
      }
    }
  }

  // Start DFS from the first index
  dfs(0);

  // Check if all indices are visited
  return visited.every(v => v);
}


function gcd(a: number, b: number) {
  while (b !== 0) {
    [a, b] = [b, a % b]
  }
  return a
}


const nums = [2, 3, 6]
// Выход: true
// Объяснение: В этом примере возможны 3 пары индексов: (0, 1), (0, 2) и (1, 2).
//   Чтобы перейти от индекса 0 к индексу 1, мы можем использовать последовательность обходов 0 -> 2 -> 1, где мы переходим от индекса 0 к индексу 2, потому что gcd(nums[0], nums[2]) = gcd(2 , 6) = 2 > 1, а затем перейти от индекса 2 к индексу 1, поскольку НОД(nums[2], nums[1]) = НОД(6, 3) = 3 > 1.
// Чтобы перейти от индекса 0 к индексу 2, мы можем просто перейти напрямую, потому что gcd(nums[0], nums[2]) = gcd(2, 6) = 2 > 1. Аналогично, чтобы перейти от индекса 1 к индексу 2, мы можем просто пойти напрямую, потому что gcd(nums[1], nums[2]) = gcd(3, 6) = 3 > 1.

const nums2 = [3, 9, 5]
// Вывод: false
// Объяснение: В этом примере никакая последовательность обходов не может привести нас от индекса 0 к индексу 2. Итак, мы возвращаем false.

const nums3 = [4, 3, 12, 8]
// Вывод: true
// Объяснение: Существует 6 возможных пар индексов для перемещения между: (0, 1), (0, 2), (0, 3), (1, 2), (1, 3) и (2, 3). Для каждой пары существует допустимая последовательность обходов, поэтому мы возвращаем true.

canTraverseAllPairs(nums)

export {}