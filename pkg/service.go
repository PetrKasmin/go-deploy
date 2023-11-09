package pkg

import "log"

type Service struct {
	Name string
}

func (s *Service) Hello() {
	log.Printf("Hello %s!\n", s.Name)
}
