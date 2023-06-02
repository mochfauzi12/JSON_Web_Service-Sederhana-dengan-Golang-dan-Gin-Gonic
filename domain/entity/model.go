package entity

type PersonResponse struct {
	Gender   string `json:"gender"`
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
	Picture  string `json:"picture"`
}

type RandomUserResponse struct {
	Results []struct {
		Gender string `json:"gender"`
		Name   struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Location struct {
			Street struct {
				Number int    `json:"number"`
				Name   string `json:"name"`
			} `json:"street"`
			City     string `json:"city"`
			State    string `json:"state"`
			Postcode int    `json:"postcode"`
		} `json:"location"`
		Picture struct {
			Large     string `json:"large"`
			Medium    string `json:"medium"`
			Thumbnail string `json:"thumbnail"`
		} `json:"picture"`
	} `json:"results"`
	Info struct {
		Seed    string `json:"seed"`
		Results int    `json:"results"`
		Page    int    `json:"page"`
		Version string `json:"version"`
	} `json:"info"`
}

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Response struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	Elapsed_time string `json:"elapsed_time"`
}
