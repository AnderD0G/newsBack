package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	//"github.com/labstack/echo/v4/middleware"
)

const yqApi = `http://c.m.163.com/ug/api/wuhan/app/data/list-total?t=330415245809`

func main() {
	var t T
	e := echo.New()
	e.HTTPErrorHandler = CustomPErrorHandler
	conf := e.Group("3rd")
	conf.GET("/list", func(c echo.Context) error {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		response, err := client.Get(yqApi)
		if err != nil {
			return err
		}
		if response.StatusCode != http.StatusOK {
			return errors.New("kkk")
		}
		readAll, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(readAll, &t)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, t)
	})
	err := e.Start(":14901")
	if err != nil {
		panic(err)
	}
}

func CustomPErrorHandler(err error, c echo.Context) {
	fmt.Println(err.Error())
	c.JSON(http.StatusInternalServerError, nil)

}

type T struct {
	ReqId int64  `json:"reqId"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  struct {
		ChinaTotal struct {
			Today struct {
				Confirm      interface{} `json:"confirm"`
				Suspect      interface{} `json:"suspect"`
				Heal         interface{} `json:"heal"`
				Dead         interface{} `json:"dead"`
				Severe       interface{} `json:"severe"`
				StoreConfirm interface{} `json:"storeConfirm"`
			} `json:"today"`
			Total struct {
				Confirm int `json:"confirm"`
				Suspect int `json:"suspect"`
				Heal    int `json:"heal"`
				Dead    int `json:"dead"`
				Severe  int `json:"severe"`
				Input   int `json:"input"`
			} `json:"total"`
			ExtData struct {
				NoSymptom int `json:"noSymptom"`
			} `json:"extData"`
		} `json:"chinaTotal"`
		ChinaDayList []struct {
			Date  string `json:"date"`
			Today struct {
				Confirm      int         `json:"confirm"`
				Suspect      int         `json:"suspect"`
				Heal         int         `json:"heal"`
				Dead         int         `json:"dead"`
				Severe       interface{} `json:"severe"`
				StoreConfirm int         `json:"storeConfirm"`
				Input        int         `json:"input"`
			} `json:"today"`
			Total struct {
				Confirm      int         `json:"confirm"`
				Suspect      int         `json:"suspect"`
				Heal         int         `json:"heal"`
				Dead         int         `json:"dead"`
				Severe       interface{} `json:"severe"`
				Input        int         `json:"input"`
				StoreConfirm int         `json:"storeConfirm"`
			} `json:"total"`
			ExtData        interface{} `json:"extData"`
			LastUpdateTime interface{} `json:"lastUpdateTime"`
		} `json:"chinaDayList"`
		LastUpdateTime        string `json:"lastUpdateTime"`
		OverseaLastUpdateTime string `json:"overseaLastUpdateTime"`
		AreaTree              []struct {
			Today struct {
				Confirm      *int        `json:"confirm"`
				Suspect      *int        `json:"suspect"`
				Heal         *int        `json:"heal"`
				Dead         *int        `json:"dead"`
				Severe       *int        `json:"severe"`
				StoreConfirm interface{} `json:"storeConfirm"`
				Input        int         `json:"input,omitempty"`
			} `json:"today"`
			Total struct {
				Confirm int `json:"confirm"`
				Suspect int `json:"suspect"`
				Heal    int `json:"heal"`
				Dead    int `json:"dead"`
				Severe  int `json:"severe"`
				Input   int `json:"input,omitempty"`
			} `json:"total"`
			ExtData struct {
				NoSymptom int `json:"noSymptom,omitempty"`
			} `json:"extData"`
			Name           string `json:"name"`
			Id             string `json:"id"`
			LastUpdateTime string `json:"lastUpdateTime"`
			Children       []struct {
				Today struct {
					Confirm      int  `json:"confirm"`
					Suspect      *int `json:"suspect"`
					Heal         int  `json:"heal"`
					Dead         int  `json:"dead"`
					Severe       *int `json:"severe"`
					StoreConfirm *int `json:"storeConfirm"`
					Input        int  `json:"input,omitempty"`
				} `json:"today"`
				Total struct {
					Confirm    int `json:"confirm"`
					Suspect    int `json:"suspect"`
					Heal       int `json:"heal"`
					Dead       int `json:"dead"`
					Severe     int `json:"severe"`
					Input      int `json:"input"`
					NewConfirm int `json:"newConfirm,omitempty"`
					NewDead    int `json:"newDead,omitempty"`
					NewHeal    int `json:"newHeal,omitempty"`
				} `json:"total"`
				ExtData struct {
				} `json:"extData"`
				Name           string `json:"name"`
				Id             string `json:"id"`
				LastUpdateTime string `json:"lastUpdateTime"`
				Children       []struct {
					Today struct {
						Confirm      int         `json:"confirm"`
						Suspect      *int        `json:"suspect"`
						Heal         *int        `json:"heal"`
						Dead         *int        `json:"dead"`
						Severe       *int        `json:"severe"`
						StoreConfirm interface{} `json:"storeConfirm"`
					} `json:"today"`
					Total struct {
						Confirm    int `json:"confirm"`
						Suspect    int `json:"suspect"`
						Heal       int `json:"heal"`
						Dead       int `json:"dead"`
						Severe     int `json:"severe"`
						NewHeal    int `json:"newHeal,omitempty"`
						NewConfirm int `json:"newConfirm,omitempty"`
						NewDead    int `json:"newDead,omitempty"`
					} `json:"total"`
					ExtData struct {
					} `json:"extData"`
					Name           string        `json:"name"`
					Id             string        `json:"id"`
					LastUpdateTime string        `json:"lastUpdateTime"`
					Children       []interface{} `json:"children"`
				} `json:"children"`
			} `json:"children"`
		} `json:"areaTree"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}
