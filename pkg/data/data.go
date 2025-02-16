package data

import (
	"strings"

	"github.com/Elanoran/d2go/pkg/data/area"
	"github.com/Elanoran/d2go/pkg/data/skill"
	"github.com/Elanoran/d2go/pkg/data/stat"
	"github.com/Elanoran/d2go/pkg/data/state"
)

// since stat.MaxLife is returning max life without stats, we are setting the max life value that we read from the
// game memory, overwriting this value each time it increases. It's not a good solution but it will provide
// more accurate values for the life %. This value is checked for each memory iteration.
type PointCounter struct {
	MaxPoint   int
	MaxPointBo int
}

func (pc *PointCounter) Percent(point int, maxPoint int, hasBo bool) int {
	if pc.MaxPoint == 0 && pc.MaxPointBo == 0 {
		pc.MaxPoint = maxPoint
		pc.MaxPointBo = maxPoint
	}
	if hasBo {
		if pc.MaxPointBo < point {
			pc.MaxPointBo = point
		}
		return int((float64(point) / float64(pc.MaxPointBo)) * 100)
	}
	if pc.MaxPoint < point {
		pc.MaxPoint = point
	}
	return int((float64(point) / float64(pc.MaxPoint)) * 100)
}

var maxHp PointCounter
var maxMp PointCounter

const (
	goldPerLevel = 10000

	// Monster Types
	MonsterTypeNone        MonsterType = "None"
	MonsterTypeChampion    MonsterType = "Champion"
	MonsterTypeMinion      MonsterType = "Minion"
	MonsterTypeUnique      MonsterType = "Unique"
	MonsterTypeSuperUnique MonsterType = "SuperUnique"
)

type Data struct {
	AreaOrigin Position
	Corpse     Corpse
	Monsters   Monsters
	// First slice represents X and second Y
	CollisionGrid  [][]bool
	PlayerUnit     PlayerUnit
	NPCs           NPCs
	Items          Items
	Objects        Objects
	AdjacentLevels []Level
	Rooms          []Room
	OpenMenus      OpenMenus
	Roster         Roster
	HoverData      HoverData
	TerrorZones    []area.Area
}

type Room struct {
	Position
	Width  int
	Height int
}

type HoverData struct {
	IsHovered bool
	UnitID
	UnitType int
}

func (r Room) GetCenter() Position {
	return Position{
		X: r.Position.X + r.Width/2,
		Y: r.Position.Y + r.Height/2,
	}
}

func (r Room) IsInside(p Position) bool {
	if p.X >= r.X && p.X <= r.X+r.Width {
		return p.Y >= r.Y && p.Y <= r.Y+r.Height
	}

	return false
}

func (d Data) MercHPPercent() int {
	for _, m := range d.Monsters {
		if m.IsMerc() {
			// Hacky thing to read merc life properly
			maxLife := m.Stats[stat.MaxLife] >> 8
			life := float64(m.Stats[stat.Life] >> 8)
			if m.Stats[stat.Life] <= 32768 {
				life = float64(m.Stats[stat.Life]) / 32768.0 * float64(maxLife)
			}

			return int(life / float64(maxLife) * 100)
		}
	}

	return 0
}

type RosterMember struct {
	Name     string
	Area     area.Area
	Position Position
}
type Roster []RosterMember

func (r Roster) FindByName(name string) (RosterMember, bool) {
	for _, rm := range r {
		if strings.EqualFold(rm.Name, name) {
			return rm, true
		}
	}

	return RosterMember{}, false
}

type Level struct {
	Area       area.Area
	Position   Position
	IsEntrance bool // This means the area can not be accessed just walking through it, needs to be clicked
}

type Class uint

const (
	Amazon Class = iota
	Sorceress
	Necromancer
	Paladin
	Barbarian
	Druid
	Assassin
)

type Corpse struct {
	Found     bool
	IsHovered bool
	Position  Position
}

type Position struct {
	X int
	Y int
}

type PlayerUnit struct {
	Name               string
	ID                 UnitID
	Area               area.Area
	Position           Position
	Stats              map[stat.ID]int
	Skills             map[skill.Skill]int
	States             state.States
	Class              Class
	LeftSkill          skill.Skill
	RightSkill         skill.Skill
	AvailableWaypoints []area.Area // Is only filled when WP menu is open and only for the specific selected tab
}

func (pu PlayerUnit) MaxGold() int {
	return goldPerLevel * pu.Stats[stat.Level]
}

// TotalGold returns the amount of gold, including inventory and stash
func (pu PlayerUnit) TotalGold() int {
	return pu.Stats[stat.Gold] + pu.Stats[stat.StashGold]
}

func (pu PlayerUnit) HPPercent() int {
	_, found := pu.Stats[stat.MaxLife]
	if !found {
		return 100
	}
	return maxHp.Percent(pu.Stats[stat.Life], pu.Stats[stat.MaxLife], pu.States.HasState(state.Battleorders))
}

func (pu PlayerUnit) MPPercent() int {
	_, found := pu.Stats[stat.MaxMana]
	if !found {
		return 100
	}
	return maxMp.Percent(pu.Stats[stat.Mana], pu.Stats[stat.MaxMana], pu.States.HasState(state.Battleorders))
}

func (pu PlayerUnit) HasDebuff() bool {
	debuffs := []state.State{
		state.Amplifydamage,
		state.Attract,
		state.Confuse,
		state.Conversion,
		state.Decrepify,
		state.Dimvision,
		state.Ironmaiden,
		state.Lifetap,
		state.Lowerresist,
		state.Terror,
		state.Weaken,
		state.Convicted,
		state.Poison,
		state.Cold,
		state.Slowed,
		state.BloodMana,
		state.DefenseCurse,
	}

	for _, s := range pu.States {
		for _, d := range debuffs {
			if s == d {
				return true
			}
		}
	}

	return false
}

type PointOfInterest struct {
	Name     string
	Position Position
}

type OpenMenus struct {
	Inventory     bool
	LoadingScreen bool
	NPCInteract   bool
	NPCShop       bool
	Stash         bool
	Waypoint      bool
	MapShown      bool
	SkillTree     bool
	Character     bool
	QuitMenu      bool
	Cube          bool
	SkillSelect   bool
	Anvil         bool
}
