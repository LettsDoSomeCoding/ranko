package persistence

type Player struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
}

func (p Player) GetIdentifier() string {
	return p.ID
}

func (p Player) Save() (Persistable, error) {
	return save(p)
}

func (p Player) Update() (Persistable, error) {
	return update(p)
}

func (p Player) SaveOrUpdate() (Persistable, error) {
	return saveOrUpdate(p)
}
