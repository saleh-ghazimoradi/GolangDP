package iterator

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
