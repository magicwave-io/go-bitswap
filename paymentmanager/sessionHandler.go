package paymentmanager

import (
	"errors"
	"sync"
)

type Session struct {
	Retry	int
	OriginNodeId	string
}

type SessionHandler struct {
	openSessions	map[string]*Session
	mutex			*sync.Mutex
}

func NewSessionHandler() *SessionHandler {
	return &SessionHandler{
		openSessions: make(map[string]*Session),
		mutex: &sync.Mutex{},
	}
}

func (h *SessionHandler) Open(sessionId string, originNodeId string)  {
	h.mutex.Lock()

	defer h.mutex.Unlock()

	h.openSessions[sessionId] = &Session{
		Retry:  0,
		OriginNodeId: originNodeId,
	}
}

func (h *SessionHandler) Close(sessionId string) (*Session, error) {
	h.mutex.Lock()

	defer h.mutex.Unlock()

	session, success := h.openSessions[sessionId]

	if !success {
		return nil, errors.New("unknown session")
	}

	return session, nil
}
