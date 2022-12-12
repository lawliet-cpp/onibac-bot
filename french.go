package main

import (
	"math/rand"
	"strings"

	"github.com/gocolly/colly"
)

func GetRandomFrenchExam() string {

	c := colly.NewCollector()
	var exams []string
	exam_link := ""
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if strings.Contains(link, "uploads") {
			exam_link = "https://3as.ency-education.com" + link
			exams = append(exams, exam_link)
		}
	})
	c.Visit("https://3as.ency-education.com/math-m-exams.html")
	randomIndex := rand.Intn(len(exams))
	return exams[randomIndex]
}