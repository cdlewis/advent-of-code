package Day12

import AOC.Reader
import AOC.SeqHelpers
import AOC.Memo

object Day12 {

    val permutations: ((Array[Char], Array[Int])) => BigInt = Memo.memoizeFunc[(Array[Char], Array[Int]), BigInt](
        (record, constraints) => (record.mkString(""), constraints.mkString(","))
    )((record: Array[Char], constraints: Array[Int]) =>
        val indexToReplace = record.indexWhere(_ == '?')

        if (indexToReplace == -1) {
            if (record.mkString("").split('.').map(_.length).filterNot(_ == 0) sameElements constraints) BigInt(1)
            else BigInt(0)
        } else {
            // do some pruning
            val recordCompletionIndex = record.slice(0, indexToReplace).lastIndexWhere(_ == '.')
            val islands = record.slice(0, recordCompletionIndex + 1).mkString("").split('.').map(_.length).filterNot(_ == 0)
            val couldBeValid = islands.length <= constraints.length && (islands sameElements constraints.slice(0, islands.length))

            if (couldBeValid) {
                val trimmedRecord = record.slice(recordCompletionIndex, record.length)
                val trimmedConstraints = constraints.slice(islands.length, constraints.length)
                val updatedIndexToReplace = trimmedRecord.indexWhere(_ == '?')

                val brokenSpringPerms = {
                    if (constraints.isEmpty) BigInt(0)
                    else permutations(trimmedRecord.patch(updatedIndexToReplace, Seq('#'), 1), trimmedConstraints)
                }
                val workingSpringPerms = permutations(trimmedRecord.patch(updatedIndexToReplace, Seq('.'), 1), trimmedConstraints)

                brokenSpringPerms + workingSpringPerms
            } else BigInt(0)
        }
    )

    @main
    def solve(): BigInt = {
        val rows = Reader.Read("day12_test")
            .linesIterator
            .map(_.split(' ').toSeq.toTuple2)
            .map((record, constraints) => {
                val parsedConstraints = constraints.split(',').map(_.toInt)
                (
                    (1 to 5).map(_ => record).mkString("?").toArray,
                    (1 to 5).flatMap(_ => parsedConstraints).toArray,
                )
            })

        rows.map((record, constraints) => permutations(record, constraints)).sum
    }

}
