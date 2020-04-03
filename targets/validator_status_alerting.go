package targets

import (
	"chainflow-vitwit/config"
	"encoding/json"
	"fmt"
	"log"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

// JailedAlerting to send transaction alert to telegram and mail
func ValidatorStatusAlert(ops HTTPOptions, cfg *config.Config, c client.Client) {
	log.Println("Coming inside validator status alerting")

	ops.Endpoint = ops.Endpoint + cfg.ValOperatorAddress

	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	var validatorResp ValidatorResp
	err = json.Unmarshal(resp.Body, &validatorResp)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	alertTime1 := cfg.AlertTime1
	alertTime2 := cfg.AlertTime2

	t1, _ := time.Parse(time.Kitchen, alertTime1)
	t2, _ := time.Parse(time.Kitchen, alertTime2)

	now := time.Now().UTC()
	t := now.Format(time.Kitchen)

	a1 := t1.Format(time.Kitchen)
	a2 := t2.Format(time.Kitchen)

	log.Println("a1, a2 and present time : ", a1, a2, t)

	if t == a1 || t == a2 {
		validatorStatus := validatorResp.Result.Jailed
		if !validatorStatus {
			_ = SendTelegramAlert(fmt.Sprintf("Your validator is currently voting"), cfg)
			_ = SendEmailAlert(fmt.Sprintf("Your validator is currently voting"), cfg)
			log.Println("Sent validator status alert")
		} else {
			_ = SendTelegramAlert(fmt.Sprintf("Your validator is in jailed status"), cfg)
			_ = SendEmailAlert(fmt.Sprintf("Your validator is in jailed status"), cfg)
			log.Println("Sent validator status alert")
		}
	}
	return
}
