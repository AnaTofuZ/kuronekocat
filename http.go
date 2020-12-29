package kuronekocat

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/anaskhan96/soup"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const tnekoURL = "https://toi.kuronekoyamato.co.jp/cgi-bin/tneko"
const tnekoRecord = 6

func getFromTneko(numbers []string) ([][tnekoRecord]string, error) {
	content, err := postHTTP(numbers)
	if err != nil {
		return nil, err
	}

	doc := soup.HTMLParse(content)

	meisaiContents := doc.FindAll("table", "class", "meisai")

	var fields [][tnekoRecord]string
	var headers [tnekoRecord]string
	for i, meisaiHeader := range meisaiContents[0].FindAll("tr")[0].FindAll("th") {
		headers[i] = meisaiHeader.Text()
	}
	headers[0] = "#"
	fields = append(fields, headers)

	for _, tr := range meisaiContents[0].FindAll("tr")[1:] {
		morimori := tr.FindAll("td")
		var field [tnekoRecord]string
		for i, mm := range morimori {
			field[i] = mm.FullText()
		}
		field[0] = "↓"
		fields = append(fields, field)
	}

	lastIndexMeisanContents := len(fields) - 1
	fields[lastIndexMeisanContents][0] = "□"
	return fields, nil
}

func postHTTP(numbers []string) (string, error) {
	queryParams := url.Values{}
	for i, number := range numbers {
		queryParams.Add(fmt.Sprintf("number%02d", i+1), number)
	}
	queryParams.Add("number00", "1") // must

	resp, err := http.PostForm(tnekoURL, queryParams)
	if err != nil {
		return "", err
	}
	reader := transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder())
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
