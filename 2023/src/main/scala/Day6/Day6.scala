package Day6

import AOC.Reader
import AOC.SeqHelpers

object Day6 {
    def solve(): BigInt = {
        val (raceTime, recordDistance) = Reader
            .Read("day6_test")
            .split("\n")
            .map(_.split(' ').filterNot(_.isEmpty).tail.mkString)
            .map(BigInt(_))
            .toSeq
            .toTuple2

        Range.BigInt(0, raceTime, 1)
            .map(holdTime => (raceTime - holdTime) * holdTime > recordDistance)
            .count(_ == true)
    }
}

