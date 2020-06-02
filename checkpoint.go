package main

import (
	"strings"
	//"fmt"

	"github.com/ivan-bogach/utils"
)

// Allowed - ...
type Allowed struct {
	members []*Member
}

func (a Allowed) IDExist(id string) bool {
	ids := GetIDs()
	if utils.StringInSlice(id, ids) {
		return true
	} else {
		return false
	}
}

func GetIDs() []string {
	var a []string
	sl, _ := utils.ReadFileLines("collegues.csv")
	for _, s := range(sl) {
		sl := strings.Split(s,",")
		a = append(a, sl[1])
	}
	return a
}

func (a Allowed) AppendMember(id, name, image string) {
	m := NewMember(id, name, image)
	a.members = append(a.members, m)
}

// GetAllowed
func GetAllowed () Allowed {
	var a Allowed
	sl, _ := utils.ReadFileLines("collegues.csv")
	for _, s := range(sl) {
		sl := strings.Split(s,",")
		a.AppendMember(sl[0], sl[1], sl[2])
	}
	return a	
}

func (a Allowed) GetName(id string) string {
	var name string
	sl, _ := utils.ReadFileLines("collegues.csv")
	for _, s := range(sl) {
		sl := strings.Split(s,",")
		if sl[1] == id {
			name = sl[0]
		}
	}
	return name
}

// GetMembers ...
func GetMembers() map[string]string {
	s, _ := utils.ReadFileLines("collegues.csv")
	m := utils.StringToMap(strings.Join(s, "\n"), "\n", ",")
	return m
}

// Member - ...
type Member struct {
	id   string
	name string
	image string
}

func NewMember(id, name, image string) *Member {
	return &Member{id: id, name: name, image: image}
}





