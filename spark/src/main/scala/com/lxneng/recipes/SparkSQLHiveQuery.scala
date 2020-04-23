package com.lxneng.recipes

import org.apache.spark.sql.SparkSession

object SparkSQLHiveQuery {
  def main(args: Array[String]): Unit = {
    val spark = SparkSession
      .builder()
      .master("local[*]")
      .appName("Spark Hive Example")
      .config("hive.metastore.uris", "thrift://localhost:9083")
      .enableHiveSupport()
      .getOrCreate()

    spark.sql("select * from pokes").show(5)

    spark.stop()
  }
}
