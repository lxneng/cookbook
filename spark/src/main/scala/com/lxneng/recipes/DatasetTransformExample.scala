package com.lxneng.recipes

import org.apache.spark.sql.catalyst.encoders.ExpressionEncoder
import org.apache.spark.sql.{Dataset, SparkSession}

case class FlightSchedule(flightNo: String, departure: Long, arrival: Long)
case class FlightInfo(flight: String, duration: String)

object DatasetTransformExample {

  def main(args: Array[String]): Unit = {
    val spark = SparkSession.builder().master("local[*]").appName("transform").getOrCreate()

    import spark.implicits._

    val schedules: Dataset[FlightSchedule] = Seq(
      FlightSchedule("AAA", 1569315038, 1569319183),
      FlightSchedule("ABB", 1569290498, 1569298178),
      FlightSchedule("ACC", 1567318178, 1567351838)
    ).toDS()

    schedules.show()

    val flightInfo: Dataset[FlightInfo] = schedules.transform(makeFlightInfo)

    flightInfo.show()
  }

  def makeFlightInfo(schedules: Dataset[FlightSchedule]): Dataset[FlightInfo] = {
    implicit val enc: ExpressionEncoder[FlightInfo] = ExpressionEncoder[FlightInfo]
    schedules.map(s => existingTransformationFunction(s))
  }

  def existingTransformationFunction(flightSchedule: FlightSchedule): FlightInfo = {
    val duration = (flightSchedule.arrival - flightSchedule.departure) / 60 / 60
    FlightInfo(flightSchedule.flightNo, s"$duration hrs")
  }
}
