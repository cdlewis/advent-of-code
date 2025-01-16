package Day11

import AOC.Reader
import AOC.Point
import AOC.TupleAsPoint
import AOC.Graph.dijkstra
import AOC.SeqHelpers

object Day11 {

    @main
    def solve(): BigInt = {
        val map = Reader.Read("day11_test").linesIterator.toSeq

        val emptyRows = map.zipWithIndex.filter { (row, _) => row.distinct.length == 1 }.map(_._2).toSet
        val emptyColumns = map.transpose.zipWithIndex.filter { (column, _) => column.distinct.length == 1 }.map(_._2).toSet

        def edgesWithCost(current: (Int, Int)): Map[(Int, Int), BigInt] = {
            Point.manhattanDirections
                .map(_ + current)
                .filter(point => map.isDefinedAt(point._1) && map(point._1).isDefinedAt(point._2))
                .map {
                    case (row, col) if emptyRows.contains(row) => (row, col) -> BigInt(1000000)
                    case (row, col) if emptyColumns.contains(col) => (row, col) -> BigInt(1000000)
                    case point => point -> BigInt(1)
                }
                .toMap
        }

        val galaxies = map
            .indices
            .filter(row => map(row).contains('#'))
            .flatMap(row => map(row).indices.filter(map(row)(_) == '#').map((row, _)))

        val shortestPaths = galaxies.foldLeft(Map[(Int, Int), Map[(Int, Int), BigInt]]()) { (results, nextGalaxy) =>
            results + (nextGalaxy -> dijkstra(edgesWithCost, nextGalaxy)._1)
        }

        galaxies.combinations(2).map(_.toTuple2).map((pointA, pointB) => shortestPaths(pointA)(pointB)).sum
    }
}
