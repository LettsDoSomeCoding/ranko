package persistence

// Player represents a player, who can create and join competitions and take part in games
type Player struct {
	ID    int    `json:"Id"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

var playerDatabaseTable = databaseTable{nextID: 1, contents: make(map[int]Persistable)}

// Save inserts a new Player with a new ID if the ID is 0, or attempts an update if ID is populated. Returns the inserted/updated Player
func (p *Player) Save() (Persistable, error) {
	return playerDatabaseTable.save(p)
}

func (p *Player) getIdentifier() (id int, exists bool) {
	if p.ID == 0 {
		return 0, false
	}

	return p.ID, true
}

func (p *Player) setIdentifier(id int) {
	p.ID = id
}

// GetAllPlayers returns all players in the database
func GetAllPlayers() []*Player {
	v := make([]*Player, 0, len(playerDatabaseTable.contents))
	for _, val := range playerDatabaseTable.contents {
		v = append(v, val.(*Player))
	}
	return v
}
