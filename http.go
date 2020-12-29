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
const tnekoRecord = 7
const numberIDIndex = 0
const statusIndex = 1

type parsedInfoType struct {
	headers [tnekoRecord]string
	orders  [][][tnekoRecord]string
}

func getFromTneko(numbers []string) (*parsedInfoType, error) {
	content, err := postHTTP(numbers)
	if err != nil {
		return nil, err
	}

	doc := soup.HTMLParse(content)

	meisaiContents := doc.FindAll("table", "class", "meisai")

	parsedInfo := parsedInfoType{}
	headers := &parsedInfo.headers
	var fields [][][tnekoRecord]string

	for i, meisaiHeader := range meisaiContents[0].FindAll("tr")[0].FindAll("th") {
		headers[i+1] = meisaiHeader.Text()
	}
	headers[statusIndex] = "#"
	headers[numberIDIndex] = "注文番号"

	for n, order := range meisaiContents {
		var eachFields [][tnekoRecord]string
		for _, tr := range order.FindAll("tr")[1:] {
			morimori := tr.FindAll("td")
			var field [tnekoRecord]string
			for i, mm := range morimori {
				field[i+1] = mm.FullText()
			}
			field[statusIndex] = "↓"
			field[numberIDIndex] = numbers[n]
			eachFields = append(eachFields, field)
		}
		fields = append(fields, eachFields)
	}

	for i, eachFields := range fields {
		lastIndexMeisanContents := len(eachFields) - 1
		fields[i][lastIndexMeisanContents][statusIndex] = "□"
	}
	parsedInfo.orders = fields

	return &parsedInfo, nil
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
