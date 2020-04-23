name := "recipes"

version := "0.1"

scalaVersion := "2.11.12"

val sparkVersion = "2.4.5"

libraryDependencies ++= Seq(
  "com.amazon.redshift" % "redshift-jdbc42" % "1.2.8.1005",
  "org.apache.spark" %% "spark-core" % sparkVersion,
  "org.apache.spark" %% "spark-sql" % sparkVersion,
  "org.apache.spark" %% "spark-hive" % sparkVersion,
  "org.apache.spark" %% "spark-streaming" % sparkVersion,
  "org.apache.spark" %% "spark-streaming-kafka-0-10" % sparkVersion,
  "mysql" % "mysql-connector-java" % "5.1.48",
  "org.elasticsearch" %% "elasticsearch-spark-20" % "7.6.2"
)
resolvers ++= Seq(
  "huawei"    at "https://mirrors.huaweicloud.com/repository/maven/",
  "redshift"  at "https://s3.amazonaws.com/redshift-maven-repository/release",
  "central"   at "https://repo1.maven.org/maven2/"
)