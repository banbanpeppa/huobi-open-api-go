package rest

const (
	HTTP_OK    = "ok"
	HTTP_ERROR = "error"
)

type Error struct {
	Ts      int    `json:"ts"`
	Status  string `json:"status"`
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

type Handler struct {
	Params   *ApiParameter
	listener chan interface{}
}

func NewDefaultRestHandler() *Handler {
	apiParams := CreateDefaultApiParameter()
	return &Handler{Params: apiParams, listener: make(chan interface{})}
}

func NewRestHandler(apiParams *ApiParameter) *Handler {
	return &Handler{Params: apiParams, listener: make(chan interface{})}
}

func (handler *Handler) Listen() <-chan interface{} {
	return handler.listener
}
