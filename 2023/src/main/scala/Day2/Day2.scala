package Day2

import AOC.Reader

object Day2 {
    def solve(): Int = {
        val re = """(([0-9]+) (red|green|blue))""".r
        val games = Reader
            .Read("day2_test")
            .split("\n")
            .map(re.findAllIn) // find all pairs
            .map(_.map { token => // parse the pairs into (int, colour) tuples
                val pair = token.split(" ")
                (pair(0).toInt, pair(1))
            })
            .map(_.toSeq)

        games
            .map(
                _.groupBy(_._2) // group each pair by its colour
                    .map(_._2.map(_._1)) // convert (Count, Colour) to just Count
                    .map(_.max).product) // find the max of each colour group and multiply them
            .sum
    }
}
