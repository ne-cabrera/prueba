package serversInfoServices

import(
	"os/exec"
	"bytes"
	"os"
	"strings"
)

type WhoisInfo struct{
	Organization string
	Country string
}

func getWhois(domain string)(info WhoisInfo){
	cmd := exec.Command("whois", domain)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
  		os.Stderr.WriteString(err.Error())
	}
	const organizationWord = "Registrant Organization: "
	const countryWord = "Registrant Country: "
	var res = string(cmdOutput.Bytes())
	var organization = splitAfter(res, organizationWord)
	var country = splitAfter(res, countryWord)
	info = WhoisInfo{organization, country}
	return
}

func splitAfter(text string, word string)(splitedWord string){
	var firstIndex = strings.Index(text, word)
	splitedWord = "Not found"
	if firstIndex != -1{
		var split = text[firstIndex:]
		var lastIndex = strings.Index(split, "\n")
		splitedWord = split[len(word):lastIndex - 1]
	}
	return
}