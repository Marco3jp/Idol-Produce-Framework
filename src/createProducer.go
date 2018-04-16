package main

import (
	"net/http"
	"github.com/satori/go.uuid"
	"os"
	"encoding/json"
	"io/ioutil"
)

type producerInfo struct{
	Name  string    `json:"name"`
	UUID  uuid.UUID `json:"uuid"`
	Idol  int       `json:"idol"`
	Level int       `json:"level"`
	Exp   int       `json:"exp"`
}

type producerList struct{
	//Num string `json:"num"`
	Producers []producerInfo `json:"producers"`
}

func createProducer(w http.ResponseWriter, r *http.Request){
	producerListCheck()

	readData, err := ioutil.ReadFile("saveData/producerList.json")
	errorCheck(&err)

	producerList := new(producerList)
	json.Unmarshal(readData, &producerList)

	producerData := new(producerInfo)
	producerData.setUUID()
	producerData.statusInit()

	producerList.Producers = append(producerList.Producers,*producerData)
	export,err := json.Marshal(&producerList)
	errorCheck(&err)

	err = ioutil.WriteFile("saveData/producerList.json", export,666)
	errorCheck(&err)

}

func (p *producerInfo)setUUID(){
	p.UUID = uuid.Must(uuid.NewV4())
}

func (p *producerInfo)statusInit(){
	p.Exp = 0
	p.Level = 0
	p.Idol = 0
}

func producerListCheck() {
	//ファイルの存在確認
	_, err := os.Stat("saveData/producerList.json")
	if err != nil {
		// ファイル作成
		fp, err := os.Create("saveData/producerList.json")
		errorCheck(&err)

		//jsonの雛形作成
		tmp := new(producerList)
		b, err := json.Marshal(tmp)
		errorCheck(&err)

		//書き込み
		fp.Write(b)
		fp.Close()
	}
}

