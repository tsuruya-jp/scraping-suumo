package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"	
)

const URL = "https://suumo.jp/jj/bukken/ichiran/JJ012FC001/?ar=060&bs=021&cn=9999999&cnb=0&et=9999999&fw2=&hb=0&ht=9999999&kb=1&kt=500&sc=26101&sc=26102&sc=26103&sc=26104&sc=26105&sc=26106&sc=26107&sc=26108&sc=26109&sc=26110&sc=26111&sc=26204&sc=26206&sc=26208&sc=26209&scTemp=26101&scTemp=26102&scTemp=26103&scTemp=26104&scTemp=26105&scTemp=26106&scTemp=26107&scTemp=26108&scTemp=26109&scTemp=26110&scTemp=26111&scTemp=26204&scTemp=26206&scTemp=26208&scTemp=26209&ta=26&tb=0&tj=0&tt=9999999&pc=30&po=5&pj=1"

func main() {
	c := colly.NewCollector()

	c.OnHTML("div[class=property_unit_group]", func(e *colly.HTMLElement) {
		e.ForEach(".property_unit", func(_ int, s *colly.HTMLElement) {
			item := s.ChildText("h2")
			fmt.Println(item)
		})
	})

	c.Visit(URL)
}
