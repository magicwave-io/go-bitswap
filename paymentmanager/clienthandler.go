package paymentmanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OutgoingCommandModel struct {
	SessionId 	string
	CommandId	string
	CommandType int32
	CommandBody []byte
	NodeId		string
	CallbackUrl	string
}

type OutgoingCommandResponseModel struct {
	ResponseBody	[]byte
	CommandId		string
	NodeId			string
	SessionId 		string
}

type ValidationResponseModel struct {
	Quantity	uint32
}

type ProcessPaymentRequestModel struct {
	CallbackUrl		string
	PaymentRequest	string
	NodeId			string
	Route			[]string
}

type ProcessPaymentResponseModel struct {
	SessionId 		string
}

type PPClient struct {
	channelUrl			string
	commandListenPort	int
	sessionHandler 		*SessionHandler
}

type PaymentTransaction struct {
	TransactionSourceAddress  string
	ReferenceAmountIn         uint32
	AmountOut                 uint32
	XDR                       string
	PaymentSourceAddress	  string
	PaymentDestinationAddress string
	StellarNetworkToken       string
	ServiceSessionId		  string
}

func NewClient(channelUrl string, commandListenPort int, sessionHandler *SessionHandler) ClientHandler {
	return &PPClient{
		channelUrl:			channelUrl,
		commandListenPort:	commandListenPort,
		sessionHandler: 	sessionHandler,
	}
}

func (pm *PPClient) ProcessResponse(commandId string, responseBody []byte, nodeId string, sessionId string) {
	values := &OutgoingCommandResponseModel{
		ResponseBody: responseBody,
		CommandId:    commandId,
		NodeId:       nodeId,
		SessionId:    sessionId,
	}

	jsonValue, err := json.Marshal(values)

	if err != nil {
		log.Error("failed to serialize process command response: %s", err.Error())

		return
	}

	reply, err := http.Post(fmt.Sprintf("%s/api/gateway/processResponse", pm.channelUrl), "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		log.Error("failed to call process command response: %s", err.Error())

		return
	}

	log.Info("process command response call status: %d", reply.StatusCode)
}

func (pm *PPClient) ProcessCommand(commandId string, commandType int32, commandBody []byte, nodeId string, sessionId string) error {
	values := &OutgoingCommandModel{
		CommandId:   commandId,
		CommandType: commandType,
		CommandBody: commandBody,
		NodeId:      nodeId,
		SessionId: 	 sessionId,
		CallbackUrl: fmt.Sprintf("http://localhost:%d/api/commandResponse", pm.commandListenPort),
	}

	jsonValue, err := json.Marshal(values)

	if err != nil {
		log.Error("failed to serialize process command request: %s", err.Error())

		return err
	}

	reply, err := http.Post(fmt.Sprintf("%s/api/utility/processCommand", pm.channelUrl), "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		log.Error("failed to call process command: %s", err.Error())

		return err
	}

	defer reply.Body.Close()

	log.Info("process command call status: %d", reply.StatusCode)

	return nil
}

func (pm *PPClient) ProcessPayment(paymentRequest string, nodeId string) {
	request := &ProcessPaymentRequestModel {
		CallbackUrl:      fmt.Sprintf("http://localhost:%d/api/command", pm.commandListenPort),
		PaymentRequest:   paymentRequest,
		NodeId: 			nodeId,
		Route:			make([]string, 0), // TODO: remove to start chain payment
	}

	jsonValue, err := json.Marshal(request)

	if err != nil {
		log.Error("failed to serialize process payment request: %s", err.Error())

		return
	}

	reply, err := http.Post(fmt.Sprintf("%s/api/gateway/processPayment", pm.channelUrl), "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		log.Error("failed to call process payment: %s", err.Error())

		return
	}

	log.Info("process payment call status: %d", reply.StatusCode)

	defer reply.Body.Close()

	response := &ProcessPaymentResponseModel{}

	err = json.NewDecoder(reply.Body).Decode(response)

	if err != nil {
		log.Error("failed to read process payment response: %s", err.Error())

		return
	}

	pm.sessionHandler.Open(response.SessionId, nodeId)
}

func (pm *PPClient) ValidatePayment(paymentRequest string) (uint32, error) {
	values := map[string]string {
		"PaymentRequest":   paymentRequest,
		"ServiceType": "ipfs",
		"CommodityType": "data",
	}

	jsonValue, err := json.Marshal(values)

	if err != nil {
		log.Error("failed to serialize validate payment request: %s", err.Error())

		return 0, err
	}

	reply, err := http.Post(fmt.Sprintf("%s/api/utility/validatePayment", pm.channelUrl), "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		log.Error("failed to call validate payment: %s", err.Error())

		return 0, err
	}

	log.Info("create payment call status: %d", reply.StatusCode)

	response := &ValidationResponseModel{}

	err = json.NewDecoder(reply.Body).Decode(response)

	if err != nil {
		log.Error("failed to deserialize validate payment response: %s", err.Error())

		return 0, err
	}

	return response.Quantity, nil
}

func (pm *PPClient) CreatePaymentInfo(amount uint32) (string, error) {
	values := map[string]interface{}{"ServiceType": "ipfs", "CommodityType": "data", "Amount": amount}

	jsonValue, err := json.Marshal(values)

	if err != nil {
		log.Error("failed to serialize create payment request: %s", err.Error())

		return "", err
	}

	reply, err := http.Post(fmt.Sprintf("%s/api/utility/createPaymentInfo", pm.channelUrl), "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		log.Error("failed to call create payment: %s", err.Error())

		return "", err
	}

	defer reply.Body.Close()

	log.Info("create payment call status: %d", reply.StatusCode)

	bodyBytes, err := ioutil.ReadAll(reply.Body)

	if err != nil {
		log.Error("failed to read create payment response body: %s", err.Error())

		return "", err
	}

	return string(bodyBytes), nil
}

func (pm *PPClient) GetTransaction(sessionId string) (*PaymentTransaction, error) {
	reply, err := http.Get(fmt.Sprintf("%s/api/utility/transaction/%s", pm.channelUrl, sessionId))

	if err != nil {
		log.Error("failed to get transaction: %s", err.Error())

		return nil, err
	}

	defer reply.Body.Close()

	log.Info("get transaction call status: %d", reply.StatusCode)

	trx := &PaymentTransaction{}
	err = json.NewDecoder(reply.Body).Decode(trx)

	if err != nil {
		log.Error("failed to deserialize transaction: %s", err.Error())

		return nil, err
	}

	return trx, nil
}


