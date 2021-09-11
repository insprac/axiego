package axiego

type Error struct {
	Message string `json:"message"`
}

type Axie struct {
	Id              string          `json:"id"`
	Image           string          `json:"image"`
	Class           string          `json:"class"`
	Chain           string          `json:"chain"`
	Name            string          `json:"name"`
	Genes           string          `json:"genes"`
	Owner           string          `json:"owner"`
	BirthDate       int             `json:"birthDate"`
	BodyShape       string          `json:"bodyShape"`
	SireId          int             `json:"sireId"`
	SireClass       string          `json:"sireClass"`
	MatronId        int             `json:"matronId"`
	MatronClass     string          `json:"matronClass"`
	Stage           int             `json:"stage"`
	Title           string          `json:"title"`
	BreedCount      int             `json:"breedCount"`
	Level           int             `json:"level"`
	Stats           AxieStats       `json:"stats"`
	Figure          AxieFigure      `json:"figure"`
	Parts           []AxiePart      `json:"parts"`
	Auction         Auction         `json:"auction"`
	OwnerProfile    PublicProfile   `json:"ownerProfile"`
	BattleInfo      AxieBattleInfo  `json:"battleInfo"`
	Children        []Axie          `json:"children"`
	TransferHistory TransferRecords `json:"transferHistory"`
}

type Auction struct {
	StartingPrice     string `json:"startingPrice"`
	EndingPrice       string `json:"endingPrice"`
	StartingTimestamp string `json:"startingTimestamp"`
	EndingTimestamp   string `json:"endingTimestamp"`
	Duration          string `json:"duration"`
	TimeLeft          string `json:"timeLeft"`
	CurrentPrice      string `json:"currentPrice"`
	CurrentPriceUSD   string `json:"currentPriceUSD"`
	SuggestedPrice    string `json:"suggestedPrice"`
	Seller            string `json:"seller"`
	ListingIndex      int    `json:"listingIndex"`
	State             string `json:"state"`
}

type AxieBattleInfo struct {
	Banned   bool   `json:"banned"`
	BanUntil string `json:"banUntil"`
	Level    int    `json:"level"`
}

type AxieCardAbility struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Attack        int    `json:"attack"`
	Defense       int    `json:"defense"`
	Energy        int    `json:"energy"`
	Description   string `json:"description"`
	BackgroundUrl string `json:"backgroundUrl"`
	EffectIconUrl string `json:"effectIconUrl"`
}

type AxieFigure struct {
	Atlas string `json:"atlas"`
	Model string `json:"model"`
	Image string `json:"image"`
}

type AxiePart struct {
	Id        string            `json:"id"`
	Name      string            `json:"name"`
	Class     string            `json:"class"`
	Type      string            `json:"type"`
	Stage     int               `json:"stage"`
	Abilities []AxieCardAbility `json:"abilities"`
}

type Axies struct {
	Results []Axie `json:"results"`
	Total   int    `json:"total"`
}

type AxieStats struct {
	Hp     int `json:"hp"`
	Speed  int `json:"speed"`
	Skill  int `json:"skill"`
	Morale int `json:"morale"`
}

type PublicProfile struct {
	Name string `json:"name"`
}

type SettledAuctions struct {
	Axies Axies `json:"axies"`
}

type TransferRecord struct {
	from        string
	fromProfile PublicProfile
	to          string
	toProfile   PublicProfile
	timestamp   int
	txHash      string
	withPrice   string
	withIceUsd  string
}

type TransferRecords struct {
	Results []TransferRecord `json:"results"`
	total   int              `json:"total"`
}
