package wikipedia

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const errStr = "Something going wrong, try to change your question"

//SearchResults structure with query and all Result data
type SearchResults struct {
	ready   bool
	Query   string
	Results []Result
}

//Result structure which contains parsed results
type Result struct {
	Name, Description, URL string
}

//UnmarshalJSON set data to strcuture
func (sr *SearchResults) UnmarshalJSON(bs []byte) error {
	array := []interface{}{}
	if err := json.Unmarshal(bs, &array); err != nil {
		return err
	}
	sr.Query = array[0].(string)
	for i := range array[1].([]interface{}) {
		sr.Results = append(sr.Results, Result{
			array[1].([]interface{})[i].(string),
			array[2].([]interface{})[i].(string),
			array[3].([]interface{})[i].(string),
		})
	}
	sr.ready = true
	return nil
}

//WikipediaAPI communicate with wikipedia api
func WikipediaAPI(request string) (answer []string) {
	s := make([]string, 3)

	//Sending request
	if response, err := http.Get(request); err != nil {
		s[0] = "Wikipedia is not respond"
	} else {
		defer response.Body.Close()

		//Reading answer
		contents, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(contents))

		//Unmarshal answer and set it to SearchResults struct
		sr := &SearchResults{}
		if err = json.Unmarshal(contents, sr); err != nil {
			s[0] = errStr
		}

		//Check if struct is not empty
		if !sr.ready {
			s[0] = errStr
		} else {
			for i := range sr.Results {
				s[i] = sr.Results[i].URL
			}
		}

	}
	return s
}
