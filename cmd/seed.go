package main

import (
	"github.com/test_kompas/news_app/pkg/entity"
)

func seedAuthors() []int {
	db.AutoMigrate(&entity.Authors{})

	authors := []entity.Authors{
		{
			Name:     "Penulis 1",
			Username: "penulis_1",
			Password: "aksesmasuk1",
		},
		{
			Name:     "Penulis 2",
			Username: "penulis_2",
			Password: "aksesmasuk2",
		},
		{
			Name:     "Penulis 3",
			Username: "penulis_3",
			Password: "aksesmasuk3",
		},
	}

	authorIds := make([]int, 0)
	for _, v := range authors {
		authorId, err := authorService.AddAuthor(&v)
		if err != nil {
			panic(err)
		}
		authorIds = append(authorIds, authorId)
	}

	return authorIds
}

func seedArticles(authorIds []int) {
	db.AutoMigrate(&entity.Article{})

	articles := []entity.Article{
		{
			Title:        "China reports 37 new COVID-19 cases among Olympic personnel",
			Body:         `<p>BEIJING (Reuters) - China detected 37 new cases of COVID-19 among Olympic Games related personnel on Jan 30, up from 34 a day earlier, the organising committee of the Beijing 2022 Winter Games said on Monday.</p><p>Eight of the total were athletes or team officials who tested positive after arriving at the airport on Sunday.</p><p>Of the total infections, 28 were among new airport arrivals, with the remaining nine already in the "closed loop" bubble that separates event personnel from the public, according to a notice on the Games' official website.</p>`,
			ReleasedDate: "2022-01-30 11:11:00",
			Status:       "PUBLISHED",
			AuthorID:     authorIds[0],
		},
		{
			Title:        "New Zealand leader Ardern tests negative for coronavirus",
			Body:         `<p>WELLINGTON, New Zealand -- New Zealand Prime Minister Jacinda Ardern said Monday she has tested negative for the coronavirus after coming into close contact with an infected person on a commercial flight.</p><p>Ardern had been isolating since late Saturday after the Jan. 22 exposure first came to light. She intends to continuing isolating through Tuesday to complete a 10-day health requirement. She has had no symptoms.</p><p>The exposure occurred on a flight from the town of Kerikeri to New Zealand's largest city, Auckland. Health officials listed a dozen flights as exposure events late Saturday, possibly indicating infections among flight crews.</p><p>Ardern and Governor-General Cindy Kiro, who is also isolating while awaiting a second test, were in the Northland region to do some filming ahead of New Zealand’s national day, Waitangi Day, on Feb. 6.</p><p>New Zealand has managed to stamp out or contain the virus for much of the pandemic, and has reported just 52 virus deaths among its population of 5 million. But an outbreak of the omicron variant is starting to take hold and is expected to rapidly grow over the coming weeks.</p><p>About 77% of New Zealanders are fully vaccinated, according to Our World in Data. That figure rises to 93% of those aged 12 and over, according to New Zealand officials.</p><p>Ardern's isolation comes as Parliament is on summer break. She will chair a Cabinet meeting remotely on Tuesday.</p>`,
			ReleasedDate: "2022-01-30 07:30:00",
			Status:       "PUBLISHED",
			AuthorID:     authorIds[1],
		},
		{
			Title:        `Envoy applauds Asean’s strategic vision`,
			Body:         `<p>ASEAN carries a strategic vision using a pluralistic and flexible approach to problem-solving with always a staunch reference to the rule of law.</p><p>For France, a country in search of a better balanced multi-polar world, Asean is an essential organisation, central in Asia as well as central across the Indo-Pacific.</p><p>This was said by Ambassador-Designate of France to Brunei Bernard Regnauld-Fabre during his talk on “The French and European Union visions for Asean and Indo-Pacific” during the Business Talk “The Road Towards Post 2025 Asean Community”.</p><p>He said Asean is a very important organisation to the European Union.</p><p>“EU has invested a lot in its bilateral relations with Asean from 2014 to 2020, €800mil (RM3.7bil) have been mobilised within the framework of Team Europe with a special focus in a programme to fight Covid-19,” he added.</p>`,
			ReleasedDate: "2022-01-30 23:15:00",
			Status:       "UNPUBLISHED",
			AuthorID:     authorIds[2],
		},
	}

	err := db.Create(&articles).Error
	if err != nil {
		panic(err)
	}
}
