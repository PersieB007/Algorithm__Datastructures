package dataStructure

import "fmt"

type Song struct {
	name   string
	artist string
	next   *Song
}

type Playlist struct {
	name       string
	head       *Song
	nowPlaying *Song
}

func CreatePlaylist(name string) *Playlist {
	return &Playlist{
		name: name,
	}
}

func (p *Playlist) AddSong(name, artist string) error {
	s := &Song{
		name:   name,
		artist: artist,
	}
	if p.head == nil {
		p.head = s
	} else {
		currentNode := p.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = s
	}
	return nil
}

func (p *Playlist) showAllSongs() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Playlist is empty.")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}

	return nil
}
