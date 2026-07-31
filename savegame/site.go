package savegame

import "os"

type Site struct {
	NumSlots Int32
	Slots    []Int64
}

func (site *Site) Read(file *os.File) {
	site.NumSlots.Read(file)
	site.Slots = make([]Int64, site.NumSlots)
	for i := 0; i < len(site.Slots); i++ {
		site.Slots[i].Read(file)
	}
}

func (site *Site) Append(buf []byte) []byte {
	buf = site.NumSlots.Append(buf)
	for i := 0; i < len(site.Slots); i++ {
		buf = site.Slots[i].Append(buf)
	}
	return buf
}
