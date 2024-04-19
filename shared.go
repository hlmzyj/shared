package shared

type Item struct {
	Priority        string
	Type_           string
	Message_type    string
	Timestamp       int
	Local_timestamp int

	// book_ticker
	Ask_price float64
	Bid_price float64

	// ob
	Ask_prices       []float64
	Bid_prices       []float64
	Ask_amounts      []float64
	Bid_amounts      []float64
	Ask_agg_amounts  []float64
	Bid_agg_amounts  []float64
	Ask_flow_amounts []float64
	Bid_flow_amounts []float64
	Is_snapshot      bool

	// trade&ob_inc
	Side         string
	Price        float64
	Amount       float64
	Max_price    float64
	Min_price    float64
	First_price  float64
	Last_price   float64
	Change_price float64
	Change_rate  float64

	Index int
}

// 因子存在这个结构体里
// trade数据的那部分因子还没写（）
type Feature struct {
	Data_status int //0-未准备好， 1-准备好， 2-数据异常

	Priority     int
	Message_type string
	Timestamp    int

	//book_ticker&trade
	Ask_price float64
	Bid_price float64
	Max_buy   float64
	Min_sell  float64

	//ob 这里的参数（depth与agg的bp）可以调 或者直接改成指针来存不固定长度的数组（在func main里修改）
	//顺序为[50s,60s,30s,10s,5s,3s,2s,1s,0.5s,0.3s,0.2s,0.1s]
	Bid_ask_spread_mean      []float64
	Ob_amount_imb_mean       [][]float64
	Ob_price_amount_imb_mean [][]float64
	Ob_amount_agg_imb_mean   [][]float64
	Of_amount_imb_mean       [][]float64

	//book_ticker
	Mid_price_change_mean     []float64
	Mid_price_change_abs_mean []float64
	Mid_price_change_std      []float64
	Mid_price_change_abs_std  []float64
	Mid_price_tv              []float64
	Last_label_5s             float64

	//tr
	Tr_avg_price                    []float64
	Tr_weighted_price               []float64
	Tr_avg_price_shift              []float64
	Tr_weighted_price_shift         []float64
	Tr_amount_buy_mean              []float64
	Tr_amount_sell_mean             []float64
	Tr_amount_buy_sell_mean         []float64
	Tr_amount_buy_p50               []float64
	Tr_amount_sell_p50              []float64
	Tr_amount_buy_sell_p50          []float64
	Tr_amount_buy_p80               []float64
	Tr_amount_sell_p80              []float64
	Tr_amount_buy_sell_p80          []float64
	Tr_amount_buy_p95               []float64
	Tr_amount_sell_p95              []float64
	Tr_amount_buy_sell_p95          []float64
	Tr_amount_buy_shift             []float64
	Tr_amount_sell_shift            []float64
	Tr_amount_buy_mean_change_rate  [][]float64
	Tr_amount_sell_mean_change_rate [][]float64
	Tr_volume_buy_mean              []float64
	Tr_volume_sell_mean             []float64
	Tr_volume_buy_sell_mean         []float64
	Tr_volume_buy_p50               []float64
	Tr_volume_sell_p50              []float64
	Tr_volume_buy_sell_p50          []float64
	Tr_volume_buy_p80               []float64
	Tr_volume_sell_p80              []float64
	Tr_volume_buy_sell_p80          []float64
	Tr_volume_buy_p95               []float64
	Tr_volume_sell_p95              []float64
	Tr_volume_buy_sell_p95          []float64
	Tr_volume_buy_shift             []float64
	Tr_volume_sell_shift            []float64
	Tr_volume_buy_mean_change_rate  []float64
	Tr_volume_sell_mean_change_rate []float64
	Tr_cnt_buy                      []int
	Tr_cnt_sell                     []int
	Vol_ret_corr                    []int
	Mean_taker_eat_spread           []float64
	Std_taker_eat_spread            []float64
	Jumpback_rate                   []float64
}

type Order struct {
	Side   string
	Price  float64
	Amount float64
}

type Level struct {
	Side   string
	Price  float64
	Amount float64
}
