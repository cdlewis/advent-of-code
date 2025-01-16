package AOC

import scala.io.Source
import scala.util.Using

val BASE = "src/main/resources/"

object Reader {
    def Read(file: String): String = Using(Source.fromFile(BASE + file))(_.mkString).get
}
