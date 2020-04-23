package com.lxneng.recipes

import org.apache.spark.sql.SparkSession
import org.elasticsearch.spark.sql._

case class Album(artist:String, yearOfRelease:Int, albumName: String)

object WriteDataToElasticsearch {

  def main(args: Array[String]): Unit = {
    writeToIndex()
  }

  def writeToIndex(): Unit = {
    val spark = SparkSession
      .builder()
      .appName("WriteToES")
      .master("local[*]")
      .config("spark.es.nodes", "localhost")
      .config("spark.es.port", "9200")
      .config("spark.es.nodes.wan.only","true")
      .getOrCreate()

    import spark.implicits._

    val indexDocuments = Seq(
      Album("五月天",2005,"知足"),
      Album("Brandi Carlile",2010,"The Story"),
      Album("周杰伦", 2001,"范特西"),
      Album("孙燕姿", 2007,"经典全纪录")
    ).toDF

    indexDocuments.saveToEs("albums")

    spark.stop()

  }

}
/*
check es indexes data
~ http localhost:9200/albums/_search
*/
