package Day3

import AOC.Reader
import AOC.Point

object Day3 {

    def solve(): Int = {
        val matrix = Reader
            .Read("day3_test")
            .split("\n")
            .map(_.toSeq)
            .zipWithIndex
            .map((line, y) => (line.zipWithIndex, y))

        // Create a map of Point -> (Part Number, Number Start Point)
        val enginePartMap = matrix
            .map { case (line, y) =>
                line.foldLeft((Seq[Seq[(Char, Point)]](), Seq[(Char, Point)]())) { case ((results, partialResult), (current, x)) =>
                    if (current.isDigit) (results, partialResult :+ (current, Point(x, y)))
                    else if (partialResult.nonEmpty) (results :+ partialResult, Seq[(Char, Point)]())
                    else (results, partialResult)
                }
            }
            .flatMap { (results, remainder) =>
                if (remainder.nonEmpty) results :+ remainder
                else results
            }
            .map { partialResult =>
                val number = partialResult.map(_._1).mkString.toInt
                partialResult.map((_, point) => point -> (partialResult.head._2, number))
            }
            .flatten
            .toMap

        // Create a list of all gear positions
        val gears = matrix
            .map { (line, y) =>
                line.collect {
                    case (char, x) if char == '*' => Point(x, y)
                }
            }
            .flatten

        // Check which gear positions are adjacent to exactly two parts, compute their gear ratios
        val gearRatios = gears
            .map { gearPosition =>
                Point.directions
                    .map(direction => enginePartMap.get(gearPosition + direction))
                    .flatten
                    .toSet // dedupe references to the same value (leveraging our tracking of the start coordinates)
            }
            .map { nearbyGears =>
                if (nearbyGears.size == 2) nearbyGears.map(_._2).product
                else 0
            }

        // Sum the gear ratios to compute the final result
        gearRatios.sum
    }
}
