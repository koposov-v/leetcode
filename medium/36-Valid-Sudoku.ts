/**
 * @link https://leetcode.com/problems/valid-sudoku/
 */
function isValidSudoku(board: string[][]): boolean {

  const cols = dict<string[]>(() => [])
  const rows = dict<string[]>(() => [])
  const squares = dict<string[]>(() => [])

  for (let r = 0; r < 9; r++) {
    for (let c = 0; c < 9; c++) {
      const v = board[r][c]
      if (v === '.') continue


      const squaresKey = Math.floor(r / 3) + ',' + Math.floor(c / 3)

      if (cols[c].includes(v) ||
        rows[r].includes(v) ||
        squares[squaresKey].includes(v)
      ) {
        return false
      }
      cols[c].push(v)
      rows[r].push(v)
      squares[squaresKey].push(v)
    }
  }


  return true
}

function dict<T>(defaultValue: () => T): { [key: string]: T } {
  const dict: { [key: string]: T } = {}
  return new Proxy(dict, {
    get(target, prop: string): any {
      if (!(prop in target)) {
        target[prop] = defaultValue()
      }
      return target[prop]
    }
  })
}


const board =
  [
    ["5", "3", ".", ".", "7", ".", ".", ".", "."]
    , ["6", ".", ".", "1", "9", "5", ".", ".", "."]
    , [".", "9", "8", ".", ".", ".", ".", "6", "."]
    , ["8", ".", ".", ".", "6", ".", ".", ".", "3"]
    , ["4", ".", ".", "8", ".", "3", ".", ".", "1"]
    , ["7", ".", ".", ".", "2", ".", ".", ".", "6"]
    , [".", "6", ".", ".", ".", ".", "2", "8", "."]
    , [".", ".", ".", "4", "1", "9", ".", ".", "5"]
    , [".", ".", ".", ".", "8", ".", ".", "7", "9"]
  ]


console.log(isValidSudoku(board))
