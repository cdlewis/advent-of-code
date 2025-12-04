

ThisBuild / version := "0.1.0-SNAPSHOT"

ThisBuild / scalaVersion := "3.3.0"

lazy val root = (project in file("."))
  .settings(
    name := "advent-of-code"
  )

libraryDependencies ++= Seq(
    "org.scalatest" %% "scalatest" % "3.2.15" % Test
)

outputStrategy := Some(StdoutOutput)