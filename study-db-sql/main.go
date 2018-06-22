package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"time"
	"fmt"
)

func main()  {
	mySql()
}


func mySql() {
	db, err := sql.Open("mysql","root:nopass.123@tcp(192.168.145.5:3306)/yixin_platform?charset=utf8&parseTime=true")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO ipblock(source,ip,createdate,level,state) values(?,?,?,?,?)")
	checkErr(err)

	res,err := stmt.Exec(2,"118.118.118.118",time.Now(),0,0)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Printf("affected rows: %d \n",affect)

	lastId, err:= res.LastInsertId()
	fmt.Printf("last inserted id: %d \n" ,lastId)

	res ,err = db.Exec("INSERT INTO ipblock(source,ip,createdate,level,state) values(?,?,?,?,?)",2,"118.118.118.118",time.Now(),0,0)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Printf("affected rows: %d \n",affect)

	lastId, err = res.LastInsertId()
	fmt.Printf("last inserted id: %d \n" ,lastId)


	ipBlocks := make([]IpBlock ,0,42)
	rows, err := db.Query("SELECT id,source,createdate,ip,`level`,state FROM	ipblock	")
	checkErr(err)

	for rows.Next(){
		ipBlock := IpBlock{}

		rows.Scan(&ipBlock.Id, &ipBlock.Source,&ipBlock.CreateDate,&ipBlock.Ip,&ipBlock.Level,&ipBlock.State)

		ipBlocks = append(ipBlocks,ipBlock)
	}
	fmt.Println("id\tsource\tcreatedate\tip\tlevel\tstate")
	for _, ipBlock := range ipBlocks{
		//.Format("2006-01-02 15:04:05")
		fmt.Printf("%d\t%d\t%s\t%s\t%d\t%d\n",ipBlock.Id,ipBlock.Source,ipBlock.CreateDate.Format("2006-01-02 15:04:05"),ipBlock.Ip,ipBlock.Level,ipBlock.State)
	}



}

type  IpBlock struct {
	Id int64
	Source int32
	CreateDate  time.Time
	Ip string
	Level int32
	State int32
}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}
