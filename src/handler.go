package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
)

func startApiHandler(u *url)  {
    r := mux.NewRouter()

    r.HandleFunc( CREATE_PRODUCER , createProducer )
    /*
    r.HandleFunc( PRODUCER_INFORMATION , producerInformation )
	r.HandleFunc( CREATE_IDOL , producerInformation )
	r.HandleFunc( IDOL_CONTENTS , producerInformation )
	r.HandleFunc( PRODUCER_LIST , producerInformation )
	r.HandleFunc( IDOL_LIST , producerInformation )
	r.HandleFunc( START_PRODUCE , producerInformation )
	r.HandleFunc( MORNING_CHECK , producerInformation )
	r.HandleFunc( WEEK_CHECK , producerInformation )
	r.HandleFunc( TREND_CHECK , producerInformation )
	r.HandleFunc( INVITE_CHECK , producerInformation )
	r.HandleFunc( WEEK_HIGHLIGHT , producerInformation )
	r.HandleFunc( SCHEDULE , producerInformation )
	r.HandleFunc( IDOL_SCHEDULE , producerInformation )
	r.HandleFunc( PRODUCER_SCHEDULE , producerInformation )
	r.HandleFunc( ACTIVITY_LIST , producerInformation )
	r.HandleFunc( IDOL_ACTIVITY_LIST , producerInformation )
	r.HandleFunc( PRODUCER_ACTIVITY_LIST , producerInformation )
	r.HandleFunc( AUDITION_LIST , producerInformation )
	r.HandleFunc( LESSON_LIST , producerInformation )
	r.HandleFunc( EVENT_LIST , producerInformation )
	r.HandleFunc( NEGOTIATION_LIST , producerInformation )
	r.HandleFunc( VISIT_LIST , producerInformation )
	r.HandleFunc( COMMU_LIST , producerInformation )
	r.HandleFunc( LIVE_LIST , producerInformation )
	r.HandleFunc( START_WEEK , producerInformation )
	r.HandleFunc( ACTIVITY , producerInformation )
	r.HandleFunc( COMMU_RESULT , producerInformation )
	r.HandleFunc( NIGHT_REPORT , producerInformation )
	r.HandleFunc( IDOL_ACTIVITY_REPORT , producerInformation )
	r.HandleFunc( FAN_REPORT , producerInformation )
	r.HandleFunc( LIMIT_CHECK , producerInformation )
	r.HandleFunc( PLAN_CHECK , producerInformation )
	r.HandleFunc( REST_PRODUCE , producerInformation )
    */

	http.ListenAndServe(":"+strconv.Itoa(u.port) , r)
}