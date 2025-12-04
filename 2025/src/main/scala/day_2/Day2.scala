package Day2

import aoc.*
import scala.math.BigInt
import aoc.FromString.given_FromString_Int.given_FromString_BigInt

val test = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

@main def run() = {
  val pattern = "([0-9]+)-([0-9]+)".r
  val actualInput = Input.load(2025, 2)

  val result = pattern.findAllIn(actualInput).matchData.flatMap {
      case Groups[BigInt, BigInt](start, end) => {
        (start to end).map { x =>
          val xString = x.toString
          val repeatFound = (1 to xString.length/2).filter(xString.length % _ == 0).exists { y =>
            val target = xString.slice(0, y)
            xString.slice(y, xString.length).replaceAll(target, "").isEmpty
          }
          if (repeatFound) x else BigInt("0")
        }
      }
    }
    .reduce(_ + _)

  print(result)
}
