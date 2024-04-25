package customerror

type OptionalParameter struct {
	Request  interface{} `json:"request"`  // request when error occured
	Response interface{} `json:"response"` // response when error occured
}
