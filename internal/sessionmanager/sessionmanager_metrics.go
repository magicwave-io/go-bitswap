package sessionmanager

import "github.com/ipfs/go-bitswap/internal/session"

func (sm *SessionManager) FetchQueueCount() int {
	var val int = 0
	sm.sessLk.Lock()
	if sm.sessions != nil { // check if SessionManager was shutdown
		for _, sess := range sm.sessions {
			if v, ok := sess.(*session.Session); ok {
				val += v.FetchQueueCount()
			}
		}
	}
	sm.sessLk.Unlock()
	return val
}

func (sm *SessionManager) LiveWantsCount() int {
	var val int = 0
	sm.sessLk.Lock()
	if sm.sessions != nil { // check if SessionManager was shutdown
		for _, sess := range sm.sessions {
			if v, ok := sess.(*session.Session); ok {
				val += v.LiveWantsCount()
			}
		}
	}
	sm.sessLk.Unlock()
	return val
}
func (sm *SessionManager) LiveWantsOldestAge() int {
	var val int = 0
	sm.sessLk.Lock()
	if sm.sessions != nil { // check if SessionManager was shutdown
		for _, sess := range sm.sessions {
			if v, ok := sess.(*session.Session); ok {
				v := v.LiveWantsOldestAge()
				if val < v {
					val = v
				}

			}
		}
	}
	sm.sessLk.Unlock()
	return val
}
func (sm *SessionManager) LiveWantsFirstAge() int {
	var val int = 0
	sm.sessLk.Lock()
	if sm.sessions != nil { // check if SessionManager was shutdown
		for _, sess := range sm.sessions {
			if v, ok := sess.(*session.Session); ok {
				v := v.LiveWantsOldestAge()
				if val > v {
					val = v
				}

			}
		}
	}
	sm.sessLk.Unlock()
	return val
}
