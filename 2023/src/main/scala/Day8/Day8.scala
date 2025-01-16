package Day8

import AOC.Reader

object Day8 {

    @main
    def solve(): BigInt = {
        val lines = AOC.Reader.Read("day8_test").split('\n')
        val instructions = lines.head
        val graph = lines.collect {
            case s"$start = ($left, $right)" => start -> (left, right)
        }.toMap

        val cycleLengths = graph.keys.filter(_.endsWith("A")).map { start =>
            Iterator
                .iterate((start, 0)) { case (position, steps) =>
                    val direction = instructions(steps % instructions.length)
                    val newPosition = if ( direction == 'L') graph(position)._1 else graph(position)._2
                    (newPosition, steps + 1)
                }
                .find(_._1.endsWith("Z")).map(_._2).get
        }

        // Calculate least common multiple of cycle lengths
        cycleLengths.map(BigInt(_)).reduce((x, y) => (x * y) / x.gcd(y))
    }
}
