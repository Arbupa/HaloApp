package http2

import "net/http"

var HeaderContent = "Content-Type"
var HeaderKey = "Ocp-Apim-Subscription-Key"
var HeaderLanguage = "Accept-Language"
var HeaderContentValue = "application/json"
var HeaderKeyValue = "befa6b6ab5f4416fbd991b02f5543160"
var HeaderLanguageValue = "en"

func GetCampaignMissions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(HeaderContent, HeaderContentValue)
	w.Header().Set(HeaderKey, HeaderKeyValue)
	w.Header().Set(HeaderLanguage, HeaderLanguageValue)

}
