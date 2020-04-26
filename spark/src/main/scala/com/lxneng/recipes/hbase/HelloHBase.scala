package com.lxneng.recipes.hbase

import org.apache.hadoop.hbase.client._
import org.apache.hadoop.hbase.{CellUtil, HBaseConfiguration, TableName}
import org.apache.hadoop.hbase.util.Bytes
import scala.collection.JavaConverters._

object HelloHBase {

  def printRow(result : Result) = {
    val cells = result.rawCells();
    print( Bytes.toString(result.getRow) + " : " )
    cells.foreach { cell =>
      val col_name = Bytes.toString(CellUtil.cloneQualifier(cell))
      val col_value = Bytes.toString(CellUtil.cloneValue(cell))
      print("(%s,%s) ".format(col_name, col_value))
    }
    println()
  }

  def main(args: Array[String]): Unit = {
    val HbaseConf = HBaseConfiguration.create()
    HbaseConf.set("hbase.zookeeper.quorum","hbase-docker")
    HbaseConf.set("hbase.zookeeper.property.clientPort","2181")
    val connection = ConnectionFactory.createConnection(HbaseConf)
    val admin = connection.getAdmin
    admin.listTableNames().foreach(println)

    val tbl = connection.getTable(TableName.valueOf(Bytes.toBytes("users")))

    // put
    var put = new Put(Bytes.toBytes("10000021"))
    put.addColumn(Bytes.toBytes("info"), Bytes.toBytes("name"), Bytes.toBytes("lxneng"))
    put.addColumn(Bytes.toBytes("info"), Bytes.toBytes("nickname"), Bytes.toBytes("Eric"))
    tbl.put(put)

    // get
    val get = new Get(Bytes.toBytes("1000002"))
    val result = tbl.get(get)
    printRow(result)

    // scan
    var scan = tbl.getScanner(new Scan())
    scan.asScala.foreach { r =>
      printRow(r)
    }
    tbl.close()
    connection.close()

  }
}
