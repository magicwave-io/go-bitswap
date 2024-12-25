package paymentmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/libp2p/go-libp2p-core/peer"
	"io"
	"net/http"
)

type IncomingCommandModel struct {
	SessionId	string
	CommandId	string
	CommandType int32
	CommandBody []byte
	NodeId		string
}

type IncomingCommandResponseModel struct {
	CommandResponse	[]byte
	CommandId		string
	NodeId			string
	SessionId		string
}

type IncomingPaymentStatusResponseModel struct {
	SessionId	string
	Status		int
}

type PPCallbackServer struct {
	server			*http.Server
	peerHandler		PeerHandler
	sessionHandler 	*SessionHandler
}

func (p *PPCallbackServer) Start() {
	err := p.server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func (p *PPCallbackServer) Shutdown(ctx context.Context) {
	err := p.server.Shutdown(ctx)

	if err != nil {
		log.Error("pp channel callback server shutdown failed %s", err.Error())
	}
}

func NewServer(commandListenPort int, peerHandler PeerHandler, sessionHandler *SessionHandler) CallbackHandler {
	router := mux.NewRouter()

	callbackServer := &PPCallbackServer{
		peerHandler: peerHandler,
		sessionHandler: sessionHandler,
	}

	router.HandleFunc("/api/command", callbackServer.ProcessCommand).Methods("POST")
	router.HandleFunc("/api/commandResponse", callbackServer.ProcessCommandResponse).Methods("POST")
	router.HandleFunc("/api/paymentResponse", callbackServer.ProcessPaymentResponse).Methods("POST")

	callbackServer.server = &http.Server{
		Addr: fmt.Sprintf(":%d", commandListenPort),
		Handler: router,
	}

	return callbackServer
}

func (p *PPCallbackServer) ProcessPaymentResponse(w http.ResponseWriter, r *http.Request)  {
	request := &IncomingPaymentStatusResponseModel{}

	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())

		return
	}

	session, err := p.sessionHandler.Close(request.SessionId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, err.Error())

		return
	}

	targetId, err := peer.IDHexDecode(session.OriginNodeId)

	p.peerHandler.PaymentStatusResponse(targetId, request.SessionId, request.Status == 1)

	w.WriteHeader(http.StatusOK)
}

func (p *PPCallbackServer) ProcessCommandResponse(w http.ResponseWriter, r *http.Request) {
	// Extract command response from the request and forward it to peer
	request := &IncomingCommandResponseModel{}

	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())

		return
	}

	targetId, err := peer.IDHexDecode(request.NodeId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())

		return
	}

	p.peerHandler.PaymentResponse(targetId, request.CommandId, request.CommandResponse, request.SessionId)

	w.WriteHeader(http.StatusOK)
}

func (p *PPCallbackServer) ProcessCommand(w http.ResponseWriter, r *http.Request) {
	// Extract command request from the request and forward it to peer
	request := &IncomingCommandModel{}

	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())

		return
	}

	targetId, err := peer.IDHexDecode(request.NodeId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())

		return
	}

	p.peerHandler.PaymentCommand(targetId, request.CommandId, request.CommandBody, request.CommandType, request.SessionId)

	w.WriteHeader(http.StatusOK)
}
