package com.lxneng.recipes.stream

import org.apache.spark.{SparkConf, TaskContext}
import org.apache.kafka.common.serialization.StringDeserializer
import org.apache.spark.streaming.StreamingContext
import org.apache.spark.streaming.Seconds
import org.apache.spark.streaming.kafka010._
import org.apache.spark.streaming.kafka010.LocationStrategies.PreferConsistent
import org.apache.spark.streaming.kafka010.ConsumerStrategies.Subscribe
import org.apache.spark.sql.{DataFrame, Row}
import org.apache.spark.sql.types._
import org.apache.spark.sql.functions.split
import org.apache.spark.sql.SparkSession
import org.apache.log4j.{Level, Logger}

import java.text.SimpleDateFormat
import java.util.Date


object KafkaWebLogStreaming {
  def formatTime(mills: Long): String = {
    val dateTimePattern = "yyyy-MM-dd HH:mm:ss"
    val sdf = new SimpleDateFormat(dateTimePattern)
    sdf.format(new Date(mills))
  }

  def row(line: List[String]): Row = Row(line(0), line(1), line(2), line(3))

  def main(args: Array[String]): Unit = {
    Logger.getLogger("org.apache.spark").setLevel(Level.WARN)
    val sparkConf = new SparkConf()
      .setMaster("local[*]")
      .setAppName("action-log-etl")
      .set("hive.metastore.uris", "thrift://localhost:9083")

    val spark: SparkSession = SparkSession
      .builder()
      .appName("WeblogStreaming")
      .config(sparkConf)
      .enableHiveSupport()
      .getOrCreate()
    val streamingContext = new StreamingContext(spark.sparkContext, Seconds(10))

    spark.sql("show tables").show()

    try {
      val topics = Array("lxneng")
      val brokers = "localhost:9092"
      val groupId = "action-log-etl"
      val kafkaParams = Map[String, Object](
        "bootstrap.servers" -> brokers,
        "key.deserializer" -> classOf[StringDeserializer],
        "value.deserializer" -> classOf[StringDeserializer],
        "group.id" -> groupId,
        "auto.offset.reset" -> "latest", // earliest
        "enable.auto.commit" -> (false: java.lang.Boolean)
      )

      val stream = KafkaUtils.createDirectStream[String, String](
        streamingContext,
        PreferConsistent,
        Subscribe[String, String](topics, kafkaParams)
      )
      stream.foreachRDD { (rdd, batchTime) =>
        val time = formatTime(batchTime.milliseconds)
        val offsetRanges = rdd.asInstanceOf[HasOffsetRanges].offsetRanges
        println(time + " --------------------- start")
        rdd.foreachPartition { _ =>
          //          val partitionId = TaskContext.getPartitionId()
          val o: OffsetRange = offsetRanges(TaskContext.get.partitionId)
          println(s"${o.topic} ${o.partition} ${o.fromOffset} ${o.untilOffset}")
          println(s"\ttopic:${o.topic} partition:${o.partition} fromOffset:${o.fromOffset} untilOffset:${o.untilOffset}")
        }
        val schema = StructType(
          Seq(
            StructField(name = "ip", dataType = StringType, nullable = false),
            StructField(name = "dateTime", dataType = StringType, nullable = false),
            StructField(name = "request", dataType = StringType, nullable = false),
            StructField(name = "statusCode", dataType = StringType, nullable = false)
          )
        )

        if (!rdd.isEmpty) {
          import spark.implicits._

          val values = rdd.map(e => e.value.split(",").to[List]).map(row)
          println(values.collect().size)
          val df = spark.createDataFrame(values, schema)
            .withColumn("_tmp", split($"request", " "))
            .select($"ip", $"dateTime",
              $"_tmp".getItem(0).as("requestType"),
              $"_tmp".getItem(1).as("requestPage"),
              $"statusCode")
          saveLogToHive(spark, df)
          println(df.show(5))
          stream.asInstanceOf[CanCommitOffsets].commitAsync(offsetRanges)
        }
      }

      streamingContext.start()
      streamingContext.awaitTermination()

    } finally {
      streamingContext.stop(true, true)
    }

  }

  def saveLogToHive(spark: SparkSession, df: DataFrame) {
    spark.sql("CREATE TABLE IF NOT EXISTS webaclogs (ip STRING, dateTime STRING, requestType STRING, requestPage STRING, statusCode STRING)")
    df.write.format("hive").mode("append").saveAsTable("webaclogs")
  }

}

/*
import time

while 1:
    for i in range(1000000):
        producer.send(topic, f'1.1.1.1,23/Apr/2020:14:27:26,GET /posts/123 HTTP/1.1,200'.encode())
    time.sleep(10)
 */