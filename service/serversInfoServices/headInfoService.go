package serversInfoServices

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"regexp"
)

type HeadInfo struct{
	Title string
	Icon string
}

func getHeadInfo(url string)(headInfo HeadInfo){
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	title := getTitle(string(html))
	icon := getIcon(string(html))
	headInfo = HeadInfo{Title: title, Icon: icon}
	return
}

func getTitle(pageHtml string)(title string){
	var titleTag = "<title>"
	titleStartIndex := strings.Index(pageHtml, titleTag)
	if titleStartIndex == -1 {
        fmt.Println("No title element found")
	}
	titleEndIndex := strings.Index(pageHtml, "</title>")
	title = pageHtml[titleStartIndex + len(titleTag):titleEndIndex]
	fmt.Println(title)
	return
}

func getIcon(pageHtml string)(icon string){
	regularExp := regexp.MustCompile("<link(.|\n)*?>")
    linkTags := regularExp.FindAllString(string(pageHtml), -1)
    if linkTags == nil {
        fmt.Println("No matches.")
    } else {
        for _, linkTag := range linkTags {
			if strings.Index(linkTag, "image/x-icon") != -1{
				iconStartIndex := strings.Index(linkTag, `href="`)
				var iconSubString = linkTag[iconStartIndex + 6:]
				iconEndIndex := strings.Index(iconSubString, `"`)
				icon = iconSubString[0:iconEndIndex]
			}
        }
	}
	fmt.Println(icon)
	return
}