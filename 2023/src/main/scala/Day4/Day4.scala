package Day4

import AOC.Reader
import AOC.MapHelpers

object Day4 {
  def solve(): Int = {
    val cards = Reader
      .Read("day4_test")
      .split("\n")
      .map(_.split('|').map(_.split(' ').flatMap(_.toIntOption)))
      .zipWithIndex

    cards
      .foldLeft(cards.indices.map(_ -> 1).toMap) { case (cardCounts, (line, index)) =>
        val copiesOfThisCard = cardCounts(index)
        val winningNumbers = line.reduce((x, y) => x.intersect(y)).length

        if (winningNumbers > 0) {
            val newCounts = (index + 1 to (index + winningNumbers)).map(_ -> copiesOfThisCard).toMap
            cardCounts.mergedWith(newCounts)
        } else {
            cardCounts
        }
      }
      .values
      .sum
  }
}
