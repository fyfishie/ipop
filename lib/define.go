package lib

type Node struct {
	ID       string
	IPList   []string
	Location Location
}
type Location struct {
	Province string
	City     string
	District string
}
type MongoCollectionAddr struct {
	IP             string `json:"ip"`
	Port           string `json:"port"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	DBName         string `json:"dbname"`
	CollectionName string `json:"collection_name"`
}

type SmarkTrace struct {
	Ip      string `json:"ip"`
	Results []Hop  `json:"results"`
	Circle  bool   `json:"circle"`
}

type Hop struct {
	TTL    int    `json:"ttl"`
	Status int    `json:"status"`
	Ip     string `json:"ip"`
	Rtt    int    `json:"rtt"`
}

type SpingTrace struct {
	IP  string `json:"ip"`
	RTT int    `json:"rtt"`
	TTL int    `json:"ttl"`
}

type IPWithNode struct {
	IP   string `json:"ip"`
	Node string `json:"node"`
}

type IPUUQuest struct {
	LocaleType string   `json:"locale_type"`
	IPList     []string `json:"ips"`
}
