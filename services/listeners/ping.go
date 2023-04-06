package listeners

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Invan2/invan_notification_service/config"
	"github.com/Invan2/invan_notification_service/genproto/notification_service"
	"github.com/Invan2/invan_notification_service/models"
	"github.com/Invan2/invan_notification_service/pkg/logger"
)

type smsService struct {
	log logger.Logger
	cfg config.Config
}

type SMSService interface {
	SendSms(ctx context.Context, message *notification_service.SendSmsRequest) (*notification_service.SendSmsRespone, error)
}

func NewSMSService(log logger.Logger, cfg config.Config) SMSService {
	return &smsService{
		log: log,
		cfg: cfg,
	}
}

func (p *smsService) SendSms(ctx context.Context, message *notification_service.SendSmsRequest) (*notification_service.SendSmsRespone, error) {

	reqBody := models.SendSmsRequestToPM{
		Messages: []*models.Message{
			{
				MessageID: "invan5",
				Recipient: "998944172774",
				Sms: &models.Sms{
					Originator: "3700",
					Content: &models.Content{
						Text: "Hello Great",
					},
				},
			},
		},
	}

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(reqBodyJSON)

	req, err := http.NewRequest(http.MethodPost, p.cfg.PlayMobileBaseURL+"/broker-api/send", body)
	if err != nil {
		p.log.Error("create requesst", logger.Error(err))
		return nil, err
	}

	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", p.cfg.PlayMobileUsername, p.cfg.PlayMobilePassword))))
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	resBodyByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(resBodyByte))

	if response.StatusCode >= 400 {
		return nil, errors.New("play mobile bad request")
	}

	return &notification_service.SendSmsRespone{}, nil
}
