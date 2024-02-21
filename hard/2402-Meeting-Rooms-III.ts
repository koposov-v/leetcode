/**
 * @link https://leetcode.com/problems/meeting-rooms-iii/
 * @param n
 * @param meetings
 */


interface IMeeting {
  count: number
  lastInterval?: {
    index: number
    from: number
    to: number
  }
}

function mostBooked(n: number, meetings: number[][]): number {
  const hash = new Map<number, IMeeting>

  for (let i = 1; i <= n; i++) {
    hash.set(i, {
      count: 0
    })
  }

  let hour = 0
  let currentIndex = 0
  let counter = meetings.length
  while (counter) {
    for (const [, meeting] of hash.entries()) {
      if (currentIndex > meetings.length - 1) {
        break
      }
      const value = meetings[currentIndex]
      const [from, to] = value
      const lastInterval = meeting.lastInterval

      if (!lastInterval) {
        meeting.lastInterval = {
          index: currentIndex,
          from,
          to,
        }
        meeting.count++
        currentIndex++
        --counter
      } else if (lastInterval.to <= hour) {
        lastInterval.index = currentIndex
        lastInterval.from = hour
        lastInterval.to += to - from

        currentIndex++
        meeting.count++
        --counter
      }
    }
    hour++

  }

  let minRoom = 0
  let prevCount: number | undefined = undefined
  let isEqual = true
  for (const [room, meeting] of hash.entries()) {
    if (isEqual && meeting.count == prevCount) {
      minRoom = 0
      continue
    }

    if (prevCount == undefined) {
      prevCount = meeting.count
      continue
    }

    isEqual = false

    if (meeting.count < prevCount) {
      minRoom = room
      prevCount = meeting.count
    }
  }

  console.log(minRoom)
  return minRoom
}


// const n = 2, meetings = [[0, 10], [1, 5], [2, 7], [3, 4]]
const n = 3, meetings = [[1,20],[2,10],[3,5],[4,9],[6,8]]
// const n = 3, meetings = [[1,20],[2,10],[3,5],[4,9],[6,8]]
mostBooked(n, meetings)













