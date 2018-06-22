package out

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []Server `xml:"server"`
	Description string `xml:",innerxml"`
}

type Server struct {

	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func Out() {
	v := &Servers{Version: "1.0"}
	fmt.Printf("struct data:\n %v \n", *v)

	v.Svs = append(v.Svs,Server{ServerName:"Shanghai_VPN",ServerIP:"127.0.0.1"})
	v.Svs = append(v.Svs,Server{ServerName:"Beijing_VPN",ServerIP:"127.0.0.2"})
	v.Description=`
 	<serverOther>
 		<serverName>Shanghai_VPN</serverName>
 		<serverIP>127.0.0.1</serverIP>
 	</serverOther>
 	<serverOther>
 		<serverName>Beijing_VPN</serverName>
 		<serverIP>127.0.0.2</serverIP>
 	</serverOther>`

	out,err := xml.MarshalIndent(v,"","	")
	if err != nil{
		fmt.Printf("marshal error:%v",err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(out)
}
