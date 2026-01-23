package tronclient

type Event struct {
	BlockNumber           int    `json:"block_number"`
	BlockTimestamp        int64  `json:"block_timestamp"`
	CallerContractAddress string `json:"caller_contract_address"`
	ContractAddress       string `json:"contract_address"`
	EventIndex            int    `json:"event_index"`
	EventName             string `json:"event_name"`
	Result                struct {
		Field1 string `json:"0"`
		Field2 string `json:"1"`
		Field3 string `json:"2"`
		From   string `json:"from"`
		To     string `json:"to"`
		Value  string `json:"value"`
	} `json:"result"`
	ResultType struct {
		From  string `json:"from"`
		To    string `json:"to"`
		Value string `json:"value"`
	} `json:"result_type"`
	Event         string `json:"event"`
	TransactionId string `json:"transaction_id"`
}

type Events []Event

type GridResult[T any] struct {
	Data    T    `json:"data"`
	Success bool `json:"success"`
	Meta    struct {
		At          int64  `json:"at"`
		Fingerprint string `json:"fingerprint"`
		Links       struct {
			Next string `json:"next"`
		} `json:"links"`
		PageSize int `json:"page_size"`
	} `json:"meta"`
}
