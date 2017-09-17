package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
  "strconv"
)

type SpeakerRow struct {
	Id             int              `json:"id"`
	Name           string           `json:"name"`
	Company        string           `json:"company"`
	CompanyLogoURL string           `json:"companyLogoUrl"`
	Title          string           `json:"title"`
	Tags           []string         `json:"tags"`
	PhotoURL       string           `json:"photoUrl"`
	Bio            string           `json:"bio"`
	ShortBio       string           `json:"shortBio"`
	Country        string           `json:"country"`
	Socials        []SpeakerSocials `json:"socials"`
	Featured       bool             `json:"featured"`
}

type SpeakerSocials struct {
	Icon string `json:"icon"`
	Link string `json:"link"`
	Name string `json:"name"`
}

type Speakers struct {
	Speakers []SpeakerRow `json:"speakers"`
}

type SessionRow struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	SlideURL    string   `json:"slideUrl"`
  Language string `json:"language"`
  Presentation string`json:"presentation"`
  Speakers []int `json:"speakers"`
	Complexity  string   `json:"complecity"`
}

func main() {
	file, err := os.Open("/Users/sinmetal/workspace/devfest_create_json/speaker.csv")
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
  sessions := make(map[string]SessionRow)
	for ir, row := range data {
		if ir == 0 {
			continue
		}

		id := ir + 100
		speakerRow := SpeakerRow{
			Id:       id,
			Featured: false,
		}
		sessionRow := SessionRow{
			Id: id,
			Speakers: []int{id},
			Language: "Japanese",
		}
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
				speakerRow.Tags = createTags(col)
			case 6:
				speakerRow.PhotoURL = col
			case 7:
				speakerRow.Bio = col
			case 8:
				speakerRow.ShortBio = col
			case 9:
				if col == "日本" {
					speakerRow.Country = "Japan"
				} else {
					speakerRow.Country = col
				}
			case 10:
				if len((col)) < 1 {
					break
				}
				speakerRow.Socials = append(speakerRow.Socials, SpeakerSocials{
					Icon: "twitter",
					//Link : fmt.Sprintf("https://twitter.com/%s/", col),
					Link: col,
					Name: "Twitter",
				})
			case 11:
				if len((col)) < 1 {
					break
				}
				speakerRow.Socials = append(speakerRow.Socials, SpeakerSocials{
					Icon: "facebook",
					//Link : fmt.Sprintf("https://www.facebook.com/%s", col),
					Link: col,
					Name: "Facebook",
				})
			case 12:
				if len((col)) < 1 {
					break
				}
				speakerRow.Socials = append(speakerRow.Socials, SpeakerSocials{
					Icon: "linkedin",
					//Link : fmt.Sprintf("https://www.linkedin.com/in/%s/", col),
					Link: col,
					Name: "LinkedIn",
				})
			case 13:
				if len((col)) < 1 {
					break
				}
				speakerRow.Socials = append(speakerRow.Socials, SpeakerSocials{
					Icon: "gplus",
					//Link : fmt.Sprintf("https://plus.google.com/%s", col),
					Link: col,
					Name: "Google",
				})
			case 14:
				if len((col)) < 1 {
					break
				}
				speakerRow.Socials = append(speakerRow.Socials, SpeakerSocials{
					Icon: "github",
					//Link : fmt.Sprintf("https://github.com/%s", col),
					Link: col,
					Name: "GitHub",
				})
			case 15:
				sessionRow.Title = col
			case 16:
				sessionRow.Description = col
			case 17:
				sessionRow.Tags = createTags(col)
			case 18:
				sessionRow.SlideURL = col
			case 19:
				sessionRow.Complexity = col
			}
		}
		speakers.Speakers = append(speakers.Speakers, speakerRow)
    sessions[strconv.Itoa(id)] = sessionRow
	}

	outputJson(speakers)
	outputJson(sessions)
}

func outputJson(value interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(b))
}

func createTags(csv string) []string {
	v := strings.Replace(csv, `"`, "", -1)
	return strings.Split(v, ",")
}
