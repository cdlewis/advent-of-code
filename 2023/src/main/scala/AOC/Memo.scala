package AOC

object Memo {

    def memoizeFunc[K, V](makeKey: K => Any)(func: K => V): K => V = {
        val cache = collection.mutable.Map[Any, V]()

        { (args: K) =>
            val key = makeKey(args)

            if (cache.contains(key)) cache(key)
            else {
                val result = func(args)
                cache.put(key, result)
                result
            }
        }
    }

}