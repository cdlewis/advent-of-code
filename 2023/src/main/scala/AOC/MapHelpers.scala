package AOC

implicit class MapHelpers[A, B](a: Map[A, B]) {
    def mergedWith(b: Map[A, B])(implicit num: Numeric[B]): Map[A, B] = {
        b ++ a.map { case (key, value) => key -> num.plus(value, b.getOrElse(key, num.zero)) }
    }

}
