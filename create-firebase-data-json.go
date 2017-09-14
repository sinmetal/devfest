package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"log"
	"strings"
	"encoding/json"
)

type SpeakerRow struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Company string `json:"company"`
	CompanyLogoURL string `json:"companyLogoUrl"`
	Title string `json:"title"`
	Tags []string `json:"tags"`
	PhotoURL string `json:"photoUrl"`
	Bio string `json:"bio"`
	ShortBio string `json:"shortBio"`
	Country string `json:"country"`
	Socials []SpeakerSocials `json:"socials"`
	Featured bool `json:"featured"`
}

type SpeakerSocials struct {
	Icon string `json:"icon"`
	Link string `json:"link"`
	Name string `json:"name"`
}

type SessionRow struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Tags []string `json:"tags"`
	SlideURL string `json:"slideUrl"`
	Complexity string `json:"complecity"`
}

type Speakers struct {
	Speakers []SpeakerRow `json:"speakers"`
}

func main() {
	file, err := os.Open("/Users/sinmetal/workspace/devfest_create_json/speaker.csv");
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()

	r := csv.NewReader(file)
	data, err := r.ReadAll()
	if err != nil {
		log.Fatalf(err.Error())
	}

	speakers := Speakers{}
	for ir, row := range data {
		if ir == 0 {
			continue
		}

		speakerRow := SpeakerRow{
			Id : ir,
			Featured: false,
		}
		// sessionRow := SessionRow{}
		for ic, col := range row {
			switch ic {
			case 1:
				speakerRow.Name = col
			case 2:
				speakerRow.Company = col
			case 3:
				speakerRow.CompanyLogoURL = col
			case 4:
				speakerRow.Title = col
			case 5:
				v := strings.Replace(col, `"`, "", -1)
				speakerRow.Tags = strings.Split(v,",")
			case 6:
				speakerRow.PhotoURL = col
			case 7:
				speakerRow.Bio = col
			case 8:
				speakerRow.ShortBio = col
			case 9:
				if (col == "日本") {
					speakerRow.Country = "Japan"
				} else {
					speakerRow.Country = col
				}
			case 10:
				if (len((col)) < 1) {
					break;
				}
				speakerRow.Socials = append(speakerRow.Socials, SpeakerSocials{
					Icon : "twitter",
					//Link : fmt.Sprintf("https://twitter.com/%s/", col),
					Link : col,
					Name : "Twitter",
				})
			case 11:
				if (len((col)) < 1) {
					break;
				}
				speakerRow.Socials = append(speakerRow.Socials, SpeakerSocials{
					Icon : "facebook",
					//Link : fmt.Sprintf("https://www.facebook.com/%s", col),
					Link : col,
					Name : "Facebook",
				})
			case 12:
				if (len((col)) < 1) {
					break;
				}
				speakerRow.Socials = append(speakerRow.Socials, SpeakerSocials{
					Icon : "linkedin",
					//Link : fmt.Sprintf("https://www.linkedin.com/in/%s/", col),
					Link : col,
					Name : "LinkedIn",
				})
			case 13:
				if (len((col)) < 1) {
					break;
				}
				speakerRow.Socials = append(speakerRow.Socials, SpeakerSocials{
					Icon : "gplus",
					//Link : fmt.Sprintf("https://plus.google.com/%s", col),
					Link : col,
					Name : "Google",
				})
			case 14:
				if (len((col)) < 1) {
					break;
				}
				speakerRow.Socials = append(speakerRow.Socials, SpeakerSocials{
					Icon : "github",
					//Link : fmt.Sprintf("https://github.com/%s", col),
					Link : col,
					Name : "GitHub",
				})
			}
		}
		speakers.Speakers = append(speakers.Speakers, speakerRow)
	}

	b, err := json.Marshal(speakers)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(b))
}

func replaceTag(tag string) string {
	switch tag {
	case "GCP":
		return "Google Cloud Platform"
	case "GKE":
		return "Google Container Engine"
	case "golang":
		return "Go"
	default:
		return tag
	}
}