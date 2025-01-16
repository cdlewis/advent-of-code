package Day5

import AOC.Reader

object Day5 {
    def solve(): BigInt = {
        val lines = Reader.Read("day5_test").split("\n") :+ "" // adding a cheeky newline makes parsing easier

        // Seed IDs
        val seeds = lines(0).split(' ').tail.grouped(2).map(x => BigRange(BigInt(x(0)), BigInt(x(0)) + BigInt(x(1)))).toSeq

        // All mappings in the order in which they are to be applied
        val allMappings = lines.slice(2, lines.length).foldLeft((Seq[Mapping](), Mapping())) { case ((result, partialMapping), current) =>
            if (current == "") (result :+ partialMapping, Mapping())
            else if (partialMapping.name == "") (result, partialMapping.copy(name = current))
            else {
                val tokens = current.split(' ').map(BigInt(_))
                val newRange = SourceDest(
                    sourceStart = tokens(1),
                    destinationStart = tokens(0),
                    step = tokens(2),
                )

                (result, partialMapping.copy(ranges = partialMapping.ranges :+ newRange))
            }
        }._1

        val result = seeds.map { seed =>
            allMappings.foldLeft(Seq(seed)) { case (seedRanges, nextMap) => // apply mapping to seed ranges
                seedRanges.flatMap(nextMap.get)
            }
        }

        result.flatten.map(_.start).min
    }
}

// Mapping is a collection of Source Range -> Dest Range mappings
case class Mapping(name: String = "", ranges: Seq[SourceDest] = Seq.empty) {
    def get(ids: BigRange): Seq[BigRange] = {
        val (allMapped, remainder) = this.ranges.foldLeft((Seq.empty[BigRange], Seq[BigRange](ids))) { case ((result, remainingSeeds), nextRange) =>
            if (remainingSeeds.isEmpty) (result, remainingSeeds)
            else {
                val (mappedRanges, unmappedRanges) = remainingSeeds.map(nextRange.intersect).unzip
                (result ++ mappedRanges.flatten, unmappedRanges.flatten.filter(_.valid()))
            }
        }

        allMapped ++ remainder
    }
}

// SourceDest holds a single mapping of Source Int -> Dest Int along with the logic for
// computing the intersection of these ranges with an input range.
case class SourceDest(sourceStart: BigInt, destinationStart: BigInt, step: BigInt) {
    val sourceEnd = sourceStart + step
    val offset = destinationStart - sourceStart
    val source = BigRange(sourceStart, sourceEnd)

    def intersect(ids: BigRange): (Option[BigRange], Seq[BigRange]) = {
        val (found, remainder) = if (ids.start <= this.sourceStart && ids.end >= this.sourceEnd) { // subsumes range
            (Some(this.source), Seq(BigRange(ids.start, this.sourceStart), BigRange(this.sourceEnd, ids.end)))
        } else if (this.source.contains(ids.start) && this.source.contains(ids.end)) { // fully contained in range
            (Some(ids), Seq.empty[BigRange])
        } else if (this.source.contains(ids.start)) { // overlaps (at start)
            (Some(BigRange(ids.start, this.sourceEnd)), Seq(BigRange(this.sourceEnd, ids.end)))
        } else if (this.source.contains(ids.end)) { // overlaps (at end)
            (Some(BigRange(this.sourceStart, ids.end)), Seq(BigRange(ids.start, this.sourceStart)))
        } else { // no overlap
            (None, Seq(ids))
        }

        (found.map(_.offset(offset)), remainder)
    }
}

// BigRange is a custom version of Range for handling very large numbers
case class BigRange(start: BigInt, end: BigInt) {
    def contains(id: BigInt): Boolean = id >= this.start && id < this.end

    def offset(other: BigInt): BigRange = this.copy(this.start + other, this.end + other)

    def valid(): Boolean = end > start
}

