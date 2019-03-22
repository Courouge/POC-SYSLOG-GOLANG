name := "demoapp"
organization := "com.courouge"
version := "0.1"
scalaVersion := "2.11.12"

resolvers += "Sonatype OSS Snapshots" at "https://oss.sonatype.org/content/repositories/snapshots"

libraryDependencies += "org.apache.kafka" %% "kafka" % "2.1.0"

libraryDependencies += "net.liftweb" %% "lift-json" % "3.3.0"

libraryDependencies += "org.apache.spark" %% "spark-sql" % "2.4.0"
libraryDependencies += "org.apache.spark" %% "spark-streaming" % "2.4.0"
libraryDependencies += "org.apache.spark" %% "spark-streaming-kafka-0-10" % "2.4.0" excludeAll(
ExclusionRule(organization = "org.spark-project.spark", name = "unused"),
ExclusionRule(organization = "org.apache.spark", name = "spark-streaming"),
ExclusionRule(organization = "org.apache.hadoop"))

assemblyMergeStrategy in assembly := {
  case PathList("META-INF", xs @ _*) => MergeStrategy.discard
  case x => MergeStrategy.first
}

target in assembly := file("build")
assemblyJarName in assembly := s"${name.value}-${version.value}.jar"
