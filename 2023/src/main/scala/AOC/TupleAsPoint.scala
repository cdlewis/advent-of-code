package AOC

implicit class TupleAsPoint(a: (Int, Int)) {

    def +(another: (Int, Int)): (Int, Int) = (a._1 + another._1, a._2 + another._2)

}