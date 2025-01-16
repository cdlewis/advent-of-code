package AOC

implicit class SeqHelpers[A](elements: Seq[A]) {

    def toTuple2 = elements match {case Seq(a, b) => (a, b) }

    def toTuple3 = elements match {case Seq(a, b, c) => (a, b, c) }

}
