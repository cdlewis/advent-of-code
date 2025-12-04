package aoc

extension (x: Int)
    infix def %%(mod: Int): Int =
        val r = x % mod
        if r < 0 then r + (if mod < 0 then -mod else mod)
        else r
