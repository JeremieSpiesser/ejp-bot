package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
    "os"
    "strconv"
)

/*
   EJP-Go-Bot
   Simple EJP bot which tells you when the electricity bill is going to go up
   Usage : ./ejp-bot <region : nord, ouest, paca or sud> <update interval>
*/

type EJPResponse struct {
	JourJ1 struct {
		EjpNord  string `json:"EjpNord"`
		EjpOuest string `json:"EjpOuest"`
		EjpPaca  string `json:"EjpPaca"`
		EjpSud   string `json:"EjpSud"`
	} `json:"JourJ1"`
	JourJ struct {
		EjpNord  string `json:"EjpNord"`
		EjpOuest string `json:"EjpOuest"`
		EjpPaca  string `json:"EjpPaca"`
		EjpSud   string `json:"EjpSud"`
	} `json:"JourJ"`
}

func getEJPResponse() EJPResponse{
    var strDate = time.Now().Format("2006-01-02")
    var url = fmt.Sprintf("https://particulier.edf.fr/bin/edf_rc/servlets/ejptemponew?Date_a_remonter=%s&TypeAlerte=EJP",strDate)
    client := &http.Client{}
    req, err := http.NewRequest("GET",url,nil)
    if (err != nil){
        log.Fatal(err)
    }
    req.Header.Set("User-Agent","PostmanRuntime/7.26.5")

    resp, err := client.Do(req)
    if (err != nil){
        panic(err)
    }

    var p EJPResponse
    body, _ := ioutil.ReadAll(resp.Body)
    err = json.Unmarshal(body,&p)
    if err != nil{
        panic(err)
    }
    return p
}

func ExtractCorrectRegion(p EJPResponse, region string) (string,string){
    if strings.Contains(strings.ToUpper(region),"PACA"){
        return p.JourJ.EjpPaca, p.JourJ1.EjpPaca
    }else if strings.Contains(strings.ToUpper(region),"SUD"){
        return p.JourJ.EjpSud, p.JourJ1.EjpSud
    }else if strings.Contains(strings.ToUpper(region),"NORD"){
        return p.JourJ.EjpNord, p.JourJ1.EjpNord
    }else {
        return p.JourJ.EjpOuest, p.JourJ1.EjpOuest
    }
}

func main(){
    var refresh int
    var err error
    var reg string
    if len(os.Args) >= 2{
        reg = os.Args[1]
    }else{
        reg = "nord"
    }
    if len(os.Args) >= 3{
        refresh, err = strconv.Atoi(os.Args[2])
        if err != nil{
            refresh = 10
            panic(err)
        }
    }else{
        refresh = 10
    }
    fmt.Println("Region : ", reg )
    fmt.Println("Refreshing every ", refresh, "seconds")
    for {
        var p EJPResponse = getEJPResponse()
        fmt.Println(ExtractCorrectRegion(p,reg))
        time.Sleep(10*time.Second)
    }

}
