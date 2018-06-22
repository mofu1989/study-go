package in

import (
	"encoding/xml"
	"os"
	"fmt"
	"io/ioutil"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []Server `xml:"server"`
	Description string `xml:",innerxml"`
}

type Server struct {
	XMLName xml.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}

func In(filePath string){
	file,err := os.Open(filePath)
	if err !=nil {
		fmt.Printf("open file error: %v\n",err)
		return
	}
	defer file.Close()
	fmt.Printf("file: %v \n",file.Name())

	data,err := ioutil.ReadAll(file)
	if err != nil{
		fmt.Printf("read data error: %v\n",err)
		return
	}
	fmt.Printf("file data:\n %v \n", string(data))
	v := Servers{}

	err = xml.Unmarshal(data,&v)
	if err!= nil{
		fmt.Printf("unmarshal error: %v\n",err)
		return
	}
	fmt.Printf("struct info:\n %v\n",v)
}
