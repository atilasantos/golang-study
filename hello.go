package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	req "net/http"
	"os"
	"strings"
	"time"
)

const monitorCicle = 3
const waitCicleTime = 4

func main() {
	fmt.Println(returnSitesFromFile())
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
		checkWebsites(sitesToMonitor())
	case 2:
		fmt.Println("Oia os log")
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid option...\nType [1,2,0]")
	}

}

func checkWebsites(sites []string) {
	for i := 0; i < monitorCicle; i++ {
		for _, site := range sites {
			isSiteUp(site)
		}
		fmt.Println()
		time.Sleep(waitCicleTime * time.Second)
	}

}

func sitesToMonitor() []string {
	sitesToMonitor := []string{"http://www.google.com", "http://www.alura.com.br", "http://www.facebook.com"}
	return sitesToMonitor
}

func isSiteUp(site string) {
	response, error := req.Get(site)
	if error != nil {
		log.Fatalln("Error:", error)
	}
	if response.StatusCode == 200 {
		fmt.Printf("Site %s carregado com status %d de sucesso!\n", site, response.StatusCode)
	} else {
		fmt.Printf("Site %s carregado com status %d de erro!\n", site, response.StatusCode)
	}
}

func returnSitesFromFile() []string {

	var sites []string

	file, error := os.Open("./sites.txt")
	if error != nil {
		log.Fatalln("Error:", error)
	}

	reader := bufio.NewReader(file)
	for {
		readLine, err := reader.ReadString('\n')
		readLine = strings.Trim(readLine, "\n")
		if err == io.EOF {
			break
		}
		sites = append(sites, readLine)
	}
	file.Close()
	return sites
}
