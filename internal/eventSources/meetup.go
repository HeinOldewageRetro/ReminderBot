package eventSources

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/boltdb/bolt"

	"bitbucket.org/Neoin/ReminderBot/internal/interfaces"
)

type Meetup struct {
	Key     string
	handler func(interfaces.Message)
}

type meetupEvent struct {
	Created       int64  `json:"created"`
	Duration      int    `json:"duration"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	RsvpLimit     int    `json:"rsvp_limit"`
	Status        string `json:"status"`
	Time          int64  `json:"time"`
	LocalDate     string `json:"local_date"`
	LocalTime     string `json:"local_time"`
	Updated       int64  `json:"updated"`
	UtcOffset     int    `json:"utc_offset"`
	WaitlistCount int    `json:"waitlist_count"`
	YesRsvpCount  int    `json:"yes_rsvp_count"`
	Venue         struct {
		ID                   int     `json:"id"`
		Name                 string  `json:"name"`
		Lat                  float64 `json:"lat"`
		Lon                  float64 `json:"lon"`
		Repinned             bool    `json:"repinned"`
		Address1             string  `json:"address_1"`
		City                 string  `json:"city"`
		Country              string  `json:"country"`
		LocalizedCountryName string  `json:"localized_country_name"`
	} `json:"venue"`
	Group struct {
		Created           int64   `json:"created"`
		Name              string  `json:"name"`
		ID                int     `json:"id"`
		JoinMode          string  `json:"join_mode"`
		Lat               float64 `json:"lat"`
		Lon               float64 `json:"lon"`
		Urlname           string  `json:"urlname"`
		Who               string  `json:"who"`
		LocalizedLocation string  `json:"localized_location"`
		Region            string  `json:"region"`
	} `json:"group"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Visibility  string `json:"visibility"`
}

func (me *meetupEvent) String() string {
	return me.Name + " by " + me.Group.Name +
		"\n " +
		"at " + me.LocalDate + " " + me.LocalTime +
		"\n " +
		me.Link
}
func (me *meetupEvent) Target() []string {
	return []string{"Programmer"}
}
func (me *meetupEvent) Country() string {
	return me.Venue.Country
}

func (m *Meetup) Start() {
	// Get topic catagories
	// https://api.meetup.com/find/topic_categories?query=tech&country=ZA&key=4035c2764b2e701d4f735e5a4f7f1f

	db, err := bolt.Open("meetup.db", os.ModePerm, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	for {

		endTime := time.Now().Add(time.Hour * 24 * 30).Format("2006-01-02T15:04:05")

		meetupURL, err := url.Parse("https://api.meetup.com/find/upcoming_events")
		if err != nil {
			panic(fmt.Errorf("Cannot parse url %v", err))
		}
		q := meetupURL.Query()
		q.Set("q", "golang")

		q.Set("topic_category", "292")
		q.Set("country", "ZA")
		q.Set("key", m.Key)
		q.Set("end_date_range", endTime)
		meetupURL.RawQuery = q.Encode()

		resp, err := http.Get(meetupURL.String())
		if err != nil {
			panic(fmt.Errorf("Cannot get meetup url %v", err))
		}
		defer resp.Body.Close()

		events := &allEvents{}
		json.NewDecoder(resp.Body).Decode(events)

		for _, event := range events.Events {
			if m.handler == nil {
				continue
			}

			hasKey := false
			err := db.Batch(func(tx *bolt.Tx) error {
				bucket, err := tx.CreateBucketIfNotExists([]byte("meetup"))
				if err != nil {
					return err
				}
				key := bucket.Get([]byte(event.ID))
				if key != nil {
					hasKey = true
				} else {
					bucket.Put([]byte(event.ID), make([]byte, 0))
				}
				return nil
			})
			if err != nil {
				fmt.Println(err)
				continue
			}
			if hasKey {
				continue
			}
			m.handler(event)

		}
		time.Sleep(time.Minute)
	}
}

func (m *Meetup) Handler(handler func(interfaces.Message)) {
	m.handler = handler
}

type allEvents struct {
	Events []*meetupEvent `json:"events"`
}
