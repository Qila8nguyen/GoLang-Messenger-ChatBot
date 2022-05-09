package fb

type (
	CoinData struct {
		Name       string
		Upperbound int
		Lowerbound int
	}

	RequestData struct {
		FollowedCoinList []CoinData
		TimeInterval     int //mins
	}

	UserData struct {
		State string //full, pend
		Data  RequestData
	}
)
