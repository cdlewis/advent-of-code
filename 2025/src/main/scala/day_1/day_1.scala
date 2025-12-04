import aoc.*


@main def run() =
  val pattern = "([LR])([0-9]+)".r
  val actualInput = Input.load(2025, 1)

  val result =
    pattern.findAllIn(actualInput).matchData.scanLeft((50, 0)) {
      case ((acc, _), Groups[String, Int]("L", amount)) =>
        val newVal = (acc - amount) %% 100
        val hits =
          if acc == 0 then amount / 100
          else (amount + 100 - acc) / 100
        (newVal, hits)

      case ((acc, _), Groups[String, Int]("R", amount)) =>
        val newVal = (acc + amount) %% 100
        val hits = (acc + amount) / 100
        (newVal, hits)
    }

  print(result.foldLeft(0) {
    case (acc, (_, count)) => acc + count
  })
