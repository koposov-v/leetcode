function isAnagram(s: string, t: string): boolean {
  /**
   * Первое решение
   */

  // if (s.length !== t.length) return false
  // s = s.split('').sort().join()
  // t = t.split('').sort().join()
  // return s === t


  /**
   * Второе решение
   */


  const mapS = new Map<string, number>()
  const mapT = new Map<string, number>()

  for (let i = 0; i < s.length; i++) {
    const countS = (mapS.get(s[i]) ?? 0) + 1
    const countT = (mapT.get(t[i]) ?? 0) + 1
    mapS.set(s[i], countS)
    mapT.set(t[i], countT)
  }

  for (const [key, value] of mapS.entries()) {
    if (mapT.get(key) !== value) {
      return false
    }
  }

  return true
}


const s = "rat", t = "car"
console.log(isAnagram(s, t))