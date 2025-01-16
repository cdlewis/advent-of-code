package Day7

import AOC.Reader
import AOC.SeqHelpers

object Day7 {

    private val cards = "AKQT98765432J".toSeq

    private val cardOrdering = cards.reverse.zipWithIndex.toMap

    def solve(): Int = {
        val handsBids = Reader.Read("day7_test").linesIterator.map(_.split(' ').toSeq.toTuple2)

        val handOrder = handsBids.toSeq.sortWith { case ((firstHand, _), (secondHand, _)) =>
            val firstHandType = handScore(firstHand)
            val secondHandType = handScore(secondHand)

            if (firstHandType == secondHandType) {
                val (_, index) = firstHand.zipWithIndex.find((char, index) => char != secondHand(index)).get
                cardOrdering(firstHand(index)) < cardOrdering(secondHand(index))
            } else firstHandType < secondHandType
        }

        handOrder.map(_._2.toInt).zipWithIndex.map((score, index) => score * (index + 1)).sum
    }

    private def handScore(hand: String): Int = cards.map { card =>
        val groups = hand.replace('J', card).toSeq.groupBy(x => x).values.map(_.length).toSeq.sorted.reverse
        10 * groups.head + groups.tail.headOption.getOrElse(0)
    }.max
}