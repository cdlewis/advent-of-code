package Day9

import AOC.Reader

object Day9 {

    def solve(): Int = {
        val sensorReadings = Reader.Read("day9_test").linesIterator.map(_.split(' ').map(_.toInt))

        val extrapolatedSensorValues = sensorReadings.map { row =>
            // continually find the difference between adjacent elements until the role is all zeroes
            val adjacentNumberDifferences = Iterator
                .iterate(row) { _
                    .sliding(2)
                    .toArray
                    .map(_.reduce((x, y) => y - x))
                }
                .takeWhile(_.exists(_ != 0))

            // work backwards up the arrays of differences calculating the extrapolated head element as we go
            adjacentNumberDifferences
                .toArray
                .reverseIterator
                .foldLeft(0) { case (result, currentRow) => currentRow.head - result }
        }

        // Sum the extrapolated readings to find our solution
        extrapolatedSensorValues.sum
    }

}
