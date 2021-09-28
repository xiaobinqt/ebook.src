package lib

import (
	"fmt"
	"time"
)

type MusicEntry struct {
	ID     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{
		musics: make([]MusicEntry, 0),
	}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, fmt.Errorf("Index out of range")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, val := range m.musics {
		if val.Name == name {
			return &val
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removeMusic := &m.musics[index]
	m.musics = append(m.musics[:index], m.musics[index+1:]...)
	return removeMusic
}

func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for i, v := range m.musics {
		if v.Name == name {
			return m.Remove(i)
		}
	}
	return nil
}

type Player interface {
	Play(source string)
}

func Play(souce, mtype string) {
	var p Player
	switch mtype {
	case "mp3":
		p = &MP3Player{}
	case "wav":
	default:
		fmt.Println("未知的播放器 ", mtype)
		return
	}

	p.Play(souce)
}

type MP3Player struct {
	stat     int
	progress int
}

func (p MP3Player) Play(source string) {
	fmt.Println("Playing Mp3 music", source)

	p.progress = 0

	for p.progress < 100 {
		time.Sleep(100 * time.Microsecond)
		fmt.Printf(".")
		p.progress += 10
	}

	fmt.Println("\n 播放结束 ", source)
}
