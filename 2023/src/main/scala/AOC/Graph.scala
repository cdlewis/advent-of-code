package AOC

object Graph {

    def dijkstra(g: ((Int, Int)) => Map[(Int, Int), BigInt], source: (Int, Int)): (Map[(Int, Int), BigInt], Map[(Int, Int), (Int, Int)]) = {
        def go(active: Set[(Int, Int)], res: Map[(Int, Int), BigInt], pred: Map[(Int, Int), (Int, Int)]):
        (Map[(Int, Int), BigInt], Map[(Int, Int), (Int, Int)]) =
            if (active.isEmpty) (res, pred)
            else {
                val node = active.minBy(res)
                val cost = res.getOrElse(node, BigInt(1))
                val neighbours = for {
                    (n, c) <- g(node) if !res.contains(n) || cost + c < res(n)
                } yield n -> (cost + c)
                val active1 = active - node ++ neighbours.keys
                val preds = neighbours.view.mapValues(_ => node)
                go(active1, res ++ neighbours, pred ++ preds)
            }

        go(Set(source), Map(source -> 0), Map.empty)
    }

}
