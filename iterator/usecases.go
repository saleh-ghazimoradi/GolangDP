package iterator

// 1. Array based iterator

// Song represents an item in the playlist
type Song struct {
	Name   string
	Artist string
}

// PlayList represents the collection as a slice
type PlayList struct {
	Songs []Song
}

// SongIterator defines the iterator interface
type SongIterator interface {
	HasNext() bool
	Next() *Song
}

// playListIterator implements the iterator for the array-based playlist
type playListIterator struct {
	Songs []Song
	index int
}

// HasNext checks if there are more songs to iterate
func (p *playListIterator) HasNext() bool {
	return p.index < len(p.Songs)
}

// Next returns the next song and advances the iterator
func (p *playListIterator) Next() *Song {
	if p.HasNext() {
		song := p.Songs[p.index]
		p.index++
		return &song
	}
	return nil
}

// AddSong adds a song to the array-based playlist
func (p *PlayList) AddSong(song Song) {
	p.Songs = append(p.Songs, song)
}

func (p *PlayList) CreateIterator() SongIterator {
	return NewPlayListIterator(p)
}

// NewPlayListIterator creates a new iterator for the array-based playlist
func NewPlayListIterator(playlists *PlayList) SongIterator {
	return &playListIterator{
		Songs: playlists.Songs,
		index: 0,
	}
}

// 2. Linked list based iterator

type SongNode struct {
	Song *Song
	Next *SongNode
}

type PlayListLinkedList struct {
	Head *SongNode
}

type playListLinkedListIterator struct {
	current *SongNode
}

func (i *playListLinkedListIterator) HasNext() bool {
	return i.current != nil
}

func (i *playListLinkedListIterator) Next() *Song {
	if !i.HasNext() {
		return nil
	}
	song := i.current.Song
	i.current = i.current.Next
	return song
}

func (p *PlayListLinkedList) AddSong(song Song) {
	newNode := &SongNode{Song: &song}
	newNode.Next = p.Head
	p.Head = newNode
}

func (p *PlayListLinkedList) CreateIterator() SongIterator {
	return NewPlayListLinkedListIterator(p)
}

func NewPlayListLinkedListIterator(playlist *PlayListLinkedList) SongIterator {
	return &playListLinkedListIterator{
		current: playlist.Head,
	}
}
