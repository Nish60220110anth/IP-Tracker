package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/Nish60220110anth/ip-tracker/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type IpInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
}

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  `Trace the IP.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				pip := net.ParseIP(ip)
				if pip == nil {
					log.Fatalln("Invalid IP address")
				}

				if pip.IsPrivate() {
					log.Printf("%s is a private IP address\n", color.New(color.FgRed).Add(color.Underline).Add(color.Italic).Sprintf("%v", pip))
					log.Printf("Please provide a public IP address, %s\n\n", color.New(color.FgYellow).Add(color.Italic).Sprint("else you will see null values"))
				}

				if pip.IsLoopback() {
					log.Printf("%s is a loopback IP address\n", color.New(color.FgRed).Add(color.Underline).Add(color.Italic).Sprintf("%v", pip))
					log.Printf("Please provide a public IP address, %s\n\n", color.New(color.FgYellow).Add(color.Italic).Sprint("else you will see null values"))
				}

				showData(pip)
			}
		} else {
			fmt.Println("Please provide IP to trace.")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

func showData(ip net.IP) {
	url := "http://ipinfo.io/" + ip.String() + "/geo"
	responseByte := getData(url)

	data := IpInfo{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Fatalln("Unable to unmarshal the response")
	}

	c := color.New(color.FgBlue, color.BgWhite).Add(color.Underline).Add(color.Bold)
	c.Println("Data found :")

	fmt.Printf("%s : %s\n%s : %s\n%s : %s\n%s : %s\n%s : %s\n%s : %s\n%s : %s\n\n", util.GenHeader("IP"), data.IP, util.GenHeader("CITY"),
		data.City, util.GenHeader("REGION"), data.Region, util.GenHeader("COUNTRY"), data.Country,
		util.GenHeader("LOCATION"), data.Loc, util.GenHeader("TIMEZONE"),
		data.Timezone, util.GenHeader("POSTAL"), data.Postal)

}

func getData(url string) []byte {

	response, err := http.Get(url)

	if err != nil {
		log.Println("Unable to get the response")
		log.Fatalf("Error: %s\n", color.New(color.FgRed).Add(color.Underline).Add(color.Italic).Sprintf("%v", err))
	}

	responseByte, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Unable to read the response")
	}

	return responseByte
}
