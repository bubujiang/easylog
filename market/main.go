package market

type actionType string
const (
	StartLog      actionType = "start-log"
	StartHttp     actionType = "start-http"
	StartAll      actionType = "start"
	ReStartLog    actionType = "restart-log"
	ReStartHttp   actionType = "restart-http"
	ReStartAll    actionType = "restart"
	StopLog       actionType = "stop-log"
	StopHttp      actionType = "stop-http"
	StopAll       actionType = "stop"
)
func (t actionType) Verify() bool {
	switch t {
	case StartLog: return true
	case StartHttp: return true
	case StartAll: return true
	case ReStartHttp: return true
	case ReStartLog: return true
	case ReStartAll: return true
	case StopLog: return true
	case StopHttp: return true
	case StopAll: return true
	}
	return false
}

type Arguments struct {
	ConfigFile string
	Action     actionType
}
//////////////
type LogFormat struct {
	Module string `json:"module"`
	Tags []string `json:"tags"`
	Time uint32 `json:"time"`
	Content struct{
		Url string `json:"url"`
		Query map[string]interface{} `json:"query"`
		Post map[string]interface{} `json:"post"`
		Header map[string]interface{} `json:"header"`
		Response interface{} `json:"resp"`
	} `json:"content"`
	//Content map[string]interface{} `json:"content"`
	//client db.DB
}