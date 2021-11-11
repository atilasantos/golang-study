package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	req "net/http"
	"os"
	"strings"
	"time"
)

const monitorCicle = 4
const waitCicleTime = 4

func main() {
	showIntroduction()
	for {
		showMenu(getUserOption())

	}

}

func showIntroduction() {
	fmt.Println("+++++++++++++++++++++++++++++")
	fmt.Println("Olá Sr.", strings.Title(os.Getenv("NAME")))
}

func getUserOption() int {
	var option int
	fmt.Println("1- Monitorar os sites")
	fmt.Println("2- Visualizar log dos sites")
	fmt.Println("0- Sair")
	fmt.Println("+++++++++++++++++++++++++++++")

	fmt.Println("Digite a opção desejada:")
	fmt.Scan(&option)
	return option
}

func showMenu(option int) {
	switch option {
	case 1:
		fmt.Println("Tá tudo funfando")
		checkWebsites(returnSitesFromFile())
	case 2:
		fmt.Println("Oia os log")
		displayLogs()
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid option...\nType [1,2,0]")
	}

}

func checkWebsites(sites []string) {
	fmt.Println(sites)
	for i := 0; i < monitorCicle; i++ {
		for _, site := range sites {
			isSiteUp(site)
		}
		fmt.Println()
		time.Sleep(waitCicleTime * time.Second)
	}

}

func isSiteUp(site string) {
	response, error := req.Get(site)
	if error != nil {
		fmt.Println("An error ocurred: ", error)
	}
	if response.StatusCode == 200 {
		fmt.Printf("Site %s carregado com status %d de sucesso!\n", site, response.StatusCode)
		logSiteStatus(site, true)
	} else {
		fmt.Printf("Site %s carregado com status %d de erro!\n", site, response.StatusCode)
		logSiteStatus(site, false)
	}
}

func returnSitesFromFile() []string {

	var sites []string

	file, error := os.Open("./sites.txt")
	if error != nil {
		fmt.Println("An error ocurred: ", error)
	}

	reader := bufio.NewReader(file)
	for {
		readLine, err := reader.ReadString('\n')
		readLine = strings.TrimSpace(readLine)
		sites = append(sites, readLine)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func logSiteStatus(site string, isUp bool) {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro.", err)
	}

	if isUp {
		file.WriteString(time.Now().Format("02/01/2006 - 15:04:05") + " - [OK] - Site - " + site + " - available.")
		file.WriteString("\n")
	} else {
		file.WriteString(time.Now().Format("02/01/2006 - 15:04:05") + " - [NOK] - Site - " + site + " - Not available")
		file.WriteString("\n")
	}
	file.Close()
}

func displayLogs() {
	file, error := ioutil.ReadFile("./logs.txt")
	if error != nil {
		fmt.Println("An error ocurred: ", error)
	}

	fmt.Println(string(file))
}
