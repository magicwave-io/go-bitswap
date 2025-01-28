package speedcontrol

import (
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
)

type MultiSpeedDetector struct {
	locker sync.Mutex
	speeds map[peer.ID]int
	states map[peer.ID]int
}

func NewSpeedDetector() *MultiSpeedDetector {
	return &MultiSpeedDetector{
		locker: sync.Mutex{},
		speeds: make(map[peer.ID]int),
		states: make(map[peer.ID]int),
	}
}

func (s *MultiSpeedDetector) WrapStream(p peer.ID, stream network.Stream) network.Stream {
	return &streamWriapper{
		pid:         p,
		Stream:      stream,
		h:           s,
		startDate:   time.Time{},
		durationSum: 0,
		bytesSum:    0,
		currentLen:  0,
	}
}

func (s *MultiSpeedDetector) PeerConnected(id peer.ID) {
	s.locker.Lock()
	s.speeds[id] = 0
	s.locker.Unlock()
}
func (s *MultiSpeedDetector) PeerDisconnected(id peer.ID) {
	s.locker.Lock()
	delete(s.speeds, id)
	s.locker.Unlock()
}

func (s *MultiSpeedDetector) DetectedSpeed(pid peer.ID, speed int) {
	s.locker.Lock()
	s.speeds[pid] = (speed + s.speeds[pid]) / 2
	s.locker.Unlock()
}

type writeHandler interface {
	DetectedSpeed(p peer.ID, speed int)
}

type streamWriapper struct {
	network.Stream
	h           writeHandler
	startDate   time.Time
	durationSum time.Duration
	bytesSum    int
	currentLen  int
	pid         peer.ID
}

func (s *streamWriapper) WriteStart(p []byte) error {
	s.startDate = time.Now()
	s.currentLen = len(p)
	return nil
}
func (s *streamWriapper) WriteStop() {
	d := time.Since(s.startDate)
	s.durationSum += d
	s.bytesSum += s.currentLen

}
func (s *streamWriapper) Write(p []byte) (n int, err error) {
	err = s.WriteStart(p)
	if err != nil {
		return
	}
	n, err = s.Stream.Write(p)

	return
}
func (s *streamWriapper) Speed() int {
	seconds := s.durationSum.Seconds()
	return int(float64(s.bytesSum) / seconds)

}
func (s *streamWriapper) Close() error {
	s.h.DetectedSpeed(s.pid, s.Speed())
	return s.Stream.Close()
}
