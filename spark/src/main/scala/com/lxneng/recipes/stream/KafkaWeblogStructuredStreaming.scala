package com.lxneng.recipes.stream

import org.apache.spark.SparkConf
import org.apache.spark.sql.SparkSession
import org.apache.spark.sql.types._
import org.apache.spark.sql.functions._

object KafkaWeblogStructuredStreaming {

  def main(args: Array[String]): Unit = {
    val sparkConf = new SparkConf()
      .setMaster("local[*]")
      .setAppName("KafkaStructuredStreaming")
      .set("hive.metastore.uris", "thrift://localhost:9083")

    val spark: SparkSession = SparkSession
      .builder()
      .config(sparkConf)
      .enableHiveSupport()
      .getOrCreate()

    import spark.implicits._

    spark.sparkContext.setLogLevel("WARN")

    try {
      val df = spark
        .readStream
        .format("kafka")
        .option("kafka.bootstrap.servers", "localhost:9092")
        .option("subscribe", "lxneng_json")
        .option("startingOffsets", "latest")
        .load()

      val schema = StructType(Seq(
        StructField("user_id", StringType, true),
        StructField("item_id", StringType, true),
        StructField("action", StringType, true),
        StructField("time", IntegerType, true)
      ))

      val df1 = df.selectExpr("CAST(value as STRING) as json", "CAST(timestamp AS TIMESTAMP) as timestamp")
        .select(from_json($"json", schema = schema).as("data"), $"timestamp")
        .select($"data.*", $"timestamp".as("log_time"))

      // ensure hive tables
      spark.sql("CREATE TABLE IF NOT EXISTS action_logs_v1 (user_id STRING, item_id STRING, action STRING, time BIGINT, log_time STRING)")
      spark.sql("select * from action_logs_v1").show(5)

//      df1.writeStream
//        .outputMode("append")
//        .queryName("table")
//        .format("console")
//        .start()
//        .awaitTermination()
      df1.writeStream.foreachBatch {
        (bdf, _) =>
          bdf.write
            .format("hive")
            .mode("append")
            .saveAsTable("action_logs_v1")
          println(bdf.count())
      }
        .start()
        .awaitTermination()

    } finally {
      spark.stop()
    }
  }
}

/*
```py
import time
import json
from datetime import datetime
from kafka import KafkaProducer

producer = KafkaProducer(
    bootstrap_servers=['localhost:9092'],
    value_serializer=lambda m: json.dumps(m).encode())
topic = 'lxneng_json'

while 1:
    for i in range(10000):
        for action in ("click", "like", "share"):
            producer.send(topic, {
                "user_id": "1000000101012312321",
                "item_id": "285ef76f4fc1212a188",
                "action": action,
                "time": int(time.mktime(datetime.now().timetuple()))
            })
    time.sleep(10)
```
 */