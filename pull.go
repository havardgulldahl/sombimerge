package main

import (
	"encoding/json"
	"fmt"
	"net/http"
  "io/ioutil"
)

/*

{
metadata: {
  count: 8040
},
results: [
    {
      _id: "55367c82feab0995be27d4e7",
      id: "8436297ddf3f39b5aa43900aad8c684e",
      src: "uploadFeed",
      title: "XXX (14) fra Fredrikstad i Østfold",
      project_id: "552e77b7fd3cf1c636eda6d7",
      project_name: "superunderskrifter",
      query: "http://superunderskrifter.mesos.nrk.no/",
      url: "http://superunderskrifter.mesos.nrk.no/8436297ddf3f39b5aa43900aad8c684e",
      tags: [
        "#underskrift"
      ],
      location: [
        10.93787,
        59.21759
      ],
      place: "Fredrikstad",
      image: {
        low: "http://gfx.nrk.no/XXX",
        thumb: "http://gfx.nrk.no/XXX",
        standard: "http://gfx.nrk.no/XXX"
      },
      video: {
        low: null,
        standard: null
      },
      user: {
        id: null,
        username: null,
        fullname: "XXX",
        avatar: null
      },
      project_metadata: [
        {
          id: "552e77b7fd3cf1c636eda6d7",
          name: "superunderskrifter",
          query: "",
          starred: false,
          impressions: 0,
          categories: [ ],
          moderation: 1
        }
      ],
      moderation: 1,
      datetime: {
        sec: 1429634178,
        usec: 594000
      },
      dates: {
        saved: {
          sec: 1429634178,
          usec: 594000
        },
        pubDate: {
          sec: 1429634178,
          usec: 0
        },
        updated: null
      }
    }
  ]
}


*/

type SombiSlice struct {
  Metadata   struct {
    Count    int
  }
  Results    []SombiRecord
}

type SombiRecord struct {
	Id         string
	Title      string
	Latitude   float64 `json:"location"`
	Longitude  float64 `json:"location"`
	Place      string
  Image      struct  {
    Thumb    string
    Full     string `json:"standard"`
  }
	Approved   bool   `json:"moderation,int"`
}

func main() {
	jobUrl := "http://sombi.nrk.no/api/1.2/data/?limit=3&skip=4000&moderation=1&metadataQuery=true&project_id=552e77b7fd3cf1c636eda6d7"
	resp, err := http.Get(jobUrl)
	if err != nil {
		panic(err)
	}
  defer resp.Body.Close()
	var slice SombiSlice
  body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &slice)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", slice)
}
