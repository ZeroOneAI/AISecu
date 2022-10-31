package viper

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

type Secret[T interface{}] struct {
	v     *viper.Viper
	mtx   sync.RWMutex
	value T
}

func New[T interface{}](filepath string) (*Secret[T], error) {
	viperConfig := viper.New()
	viperConfig.SetConfigFile(filepath)

	err := viperConfig.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var value T
	err = viperConfig.Unmarshal(&value)
	if err != nil {
		return nil, err
	}

	secret := &Secret[T]{
		v:     viperConfig,
		mtx:   sync.RWMutex{},
		value: value,
	}

	return secret, nil
}

func (s *Secret[T]) SyncOnChange() {
	s.v.OnConfigChange(func(e fsnotify.Event) {
		var value T
		if err := s.v.Unmarshal(&value); err != nil {
			return
		}
		s.mtx.Lock()
		defer s.mtx.Unlock()
		s.value = value
	})
}

func (s *Secret[T]) Get() T {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.value
}
