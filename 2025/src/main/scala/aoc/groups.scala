package aoc

import scala.util.matching.Regex

trait FromString[A]:
  def parse(s: String): Option[A]

object FromString:
  given FromString[String] with
    def parse(s: String): Option[String] = Some(s)

  given FromString[Int] with
    def parse(s: String): Option[Int] =
      s.toIntOption

object Groups:
  def unapply[A: FromString, B: FromString](m: Regex.Match): Option[(A, B)] =
    m.subgroups match
      case a :: b :: _ =>
        val fa = summon[FromString[A]]
        val fb = summon[FromString[B]]
        for
          av <- fa.parse(a)
          bv <- fb.parse(b)
        yield (av, bv)
      case _ => None
