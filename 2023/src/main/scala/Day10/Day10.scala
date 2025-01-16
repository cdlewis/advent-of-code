package Day10

import AOC.Reader

object Day10 {

    def getAdjacent(graph: Seq[String], node: (Int, Int)): Set[(Int, Int)] = {
        val (y, x) = node

        graph(y)(x) match
            case '|' => Set((y - 1, x), (y + 1, x))
            case '-' => Set((y, x - 1), (y, x + 1))
            case 'L' => Set((y - 1, x), (y, x + 1))
            case 'J' => Set((y - 1, x), (y, x - 1))
            case '7' => Set((y + 1, x), (y, x - 1))
            case 'F' => Set((y + 1, x), (y, x + 1))
            case '.' => Set()
            case 'S' => Set((y + 1, x), (y - 1, x), (y, x + 1), (y, x - 1))
                .filter((y, j) => graph.isDefinedAt(y) && graph(y).isDefinedAt(j))
                .filter(getAdjacent(graph, _).contains(y, x))
    }

    @main
    def solve(): Int = {
        val graph = Reader.Read("day10_test").linesIterator.toSeq

        val start = {
            val y = graph.indexWhere(_.contains('S'))
            (y, graph(y).indexOf('S'))
        }

        val loop = LazyList.iterate((start, getAdjacent(graph, start).head)) { (prev, current) =>
            val next = getAdjacent(graph, current) - prev
            (current, next.head)
        }

        val tilePath = start +: loop.map(_._2).takeWhile(_ != start)

        graph.indices.map { y =>
            val (_, count) = graph(y).indices.foldLeft((false, 0)) {
                case ((enclosed, count), x) if tilePath.contains(y, x) =>
                    (enclosed ^ getAdjacent(graph, (y, x)).contains(y - 1, x), count)
                case ((true, count), _) => (true, count + 1)
                case ((false, count), _) => (false, count)
            }
            count
        }.sum
    }

}
