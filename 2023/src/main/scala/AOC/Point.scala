package AOC

case class Point(x: Int, y: Int) {
    def +(anotherPoint: Point): Point = this.copy(x = this.x + anotherPoint.x, y = this.y + anotherPoint.y)
}

object Point {
    val directions: Seq[Point] = Seq((1, 0), (-1, 0), (0, 1), (0, -1), (1, 1), (-1, -1), (-1, 1), (1, -1)).map(Point.apply)

    val manhattanDirections: Seq[(Int, Int)] = Seq((1, 0), (-1, 0), (0, 1), (0, -1))
}
