package regex_util

import (
	"encoding/json"
	"regexp"
)

type Data struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Tag   string `json:"tag"`
}

func ExtractDataFromString(str string) (string, error) {
	regexTitle := regexp.MustCompile(`"title": "(.*?)"`)
	regexText := regexp.MustCompile(`"text": "(.*?)"`)
	regexTag := regexp.MustCompile(`"tag": "(.*?)"`)

	titles := regexTitle.FindAllStringSubmatch(str, -1)
	texts := regexText.FindAllStringSubmatch(str, -1)
	tags := regexTag.FindAllStringSubmatch(str, -1)

	result := []Data{}
	for i := 0; i < len(titles); i++ {
		obj := Data{
			Title: titles[i][1],
			Text:  "",
			Tag:   "",
		}
		if i < len(texts) {
			obj.Text = texts[i][1]
		}
		if i < len(tags) {
			obj.Tag = tags[i][1]
		}
		result = append(result, obj)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
