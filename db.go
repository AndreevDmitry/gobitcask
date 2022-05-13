package gobitcask

type Db struct {
	dir string
	keyDir map[string]uint64 //Key -> fileOffset)
}

func New(dir string) *Db {
	return &Db{
		dir: dir}
}

func (db *Db) Put(key string, value string) {

}

func (db *Db) Get(key string) string {
	return ""
}
