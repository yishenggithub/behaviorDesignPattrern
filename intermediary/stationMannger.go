package main

type stationManager struct {
	isPlatFormFree bool
	trainQueue     []train // 放interface的列表
}

func newStationManger() *stationManager {
	return &stationManager{
		isPlatFormFree: true,
	}
}

func (s *stationManager) canArrive(t train) bool { // 没有接口, 声明是stationManager的方法
	if s.isPlatFormFree {
		s.isPlatFormFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false

}

func (s *stationManager) notifyAboutDeparture() {
	if !s.isPlatFormFree {
		s.isPlatFormFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}
