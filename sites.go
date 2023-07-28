package main

// Site object
type Site struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// helper function to return site objects
func sites() []Site {
	var out []Site
	site := Site{
		Name: "SiteA",
		URL:  "http://aaa.com",
	}
	out = append(out, site)
	site = Site{
		Name: "SiteB",
		URL:  "http://bbb.com",
	}
	out = append(out, site)
	site = Site{
		Name: "SiteC",
		URL:  "http://ccc.com",
	}
	out = append(out, site)
	return out
}
