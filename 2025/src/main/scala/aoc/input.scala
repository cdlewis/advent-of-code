import java.net.URI
import java.net.http.{HttpClient, HttpRequest, HttpResponse}
import java.nio.file.{Files, Paths, Path}
import java.nio.charset.StandardCharsets

object Input {

  private val client: HttpClient = HttpClient.newHttpClient()

  def load(year: Int, day: Int, sessionOpt: Option[String] = sys.env.get("AOC_SESSION")): String = {
    val session = sessionOpt.getOrElse {
      throw new IllegalStateException(
        "Missing AoC session token. Set AOC_SESSION env var or pass it explicitly."
      )
    }

    val cacheFile = cachePath(year, day)
    if (Files.exists(cacheFile)) {
      return Files.readString(cacheFile, StandardCharsets.UTF_8)
    }

    val uri = URI.create(s"https://adventofcode.com/$year/day/$day/input")

    val request = HttpRequest.newBuilder(uri)
      .header("Cookie", s"session=$session")
      .header("User-Agent", "github.com/cdlewis/advent-of-code (Advent of Code helper)")
      .GET()
      .build()

    val response = client.send(request, HttpResponse.BodyHandlers.ofString())

    if (response.statusCode() != 200) {
      throw new RuntimeException(
        s"Failed to fetch AoC input: HTTP ${response.statusCode()} (${response.body().take(200)})"
      )
    }

    val body = response.body()
    Files.createDirectories(cacheFile.getParent)
    Files.writeString(cacheFile, body, StandardCharsets.UTF_8)

    body
  }

  private def cachePath(year: Int, day: Int): Path = {
    Paths.get("cache", "aoc", year.toString, f"day$day%02d.txt")
  }
}
