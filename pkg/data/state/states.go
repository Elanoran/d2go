package state

type States []State

func (s States) HasState(state State) bool {
	for _, st := range s {
		if st == state {
			return true
		}
	}

	return false
}

type State uint

const (
	None State = iota
	Freeze
	Poison
	Resistfire
	Resistcold
	Resistlightning
	Resistmagic
	Playerbody
	Resistall
	Amplifydamage
	Frozenarmor
	Cold
	Inferno
	Blaze
	Bonearmor
	Concentrate
	Enchant
	Innersight
	SkillMove
	Weaken
	Chillingarmor
	Stunned
	Spiderlay
	Dimvision
	Slowed
	Fetishaura
	Shout
	Taunt
	Conviction
	Convicted
	Energyshield
	Venomclaws
	Battleorders
	Might
	Prayer
	Holyfire
	Thorns
	Defiance
	Thunderstorm
	Lightningbolt
	Blessedaim
	Stamina
	Concentration
	Holywind
	Holywindcold
	Cleansing
	Holyshock
	Sanctuary
	Meditation
	Fanaticism
	Redemption
	Battlecommand
	Preventheal
	Conversion
	Uninterruptable
	Ironmaiden
	Terror
	Attract
	Lifetap
	Confuse
	Decrepify
	Lowerresist
	Openwounds
	Dopplezon
	Criticalstrike
	Dodge
	Avoid
	Penetrate
	Evade
	Pierce
	Warmth
	Firemastery
	Lightningmastery
	Coldmastery
	Blademastery
	Axemastery
	Macemastery
	Polearmmastery
	Throwingmastery
	Spearmastery
	Increasedstamina
	Ironskin
	Increasedspeed
	Naturalresistance
	Fingermagecurse
	Nomanaregen
	Justhit
	Slowmissiles
	Shiverarmor
	Battlecry
	Blue
	Red
	DeathDelay
	Valkyrie
	Frenzy
	Berserk
	Revive
	Itemfullset
	Sourceunit
	Redeemed
	Healthpot
	Holyshield
	JustPortaled
	Monfrenzy
	CorpseNodraw
	Alignment
	Manapot
	Shatter
	SyncWarped
	ConversionSave
	Pregnant
	State111
	Rabies
	DefenseCurse
	BloodMana
	Burning
	Dragonflight
	Maul
	CorpseNoselect
	Shadowwarrior
	Feralrage
	Skilldelay
	Tigerstrike
	Cobrastrike
	Phoenixstrike
	Fistsoffire
	Bladesofice
	Clawsofthunder
	ShrineArmor
	ShrineCombat
	ShrineResistLightning
	ShrineResistFire
	ShrineResistCold
	ShrineResistPoison
	ShrineSkill
	ShrineManaRegen
	ShrineStamina
	ShrineExperience
	FenrisRage
	Wolf
	Bear
	Bloodlust
	Changeclass
	Attached
	Hurricane
	Armageddon
	Invis
	Barbs
	Wolverine
	Oaksage
	VineBeast
	Cyclonearmor
	Clawmastery
	CloakOfShadows
	Recycled
	Weaponblock
	Cloaked
	Quickness
	Bladeshield
	Fade
	Summonresist
	Oaksagecontrol
	Wolverinecontrol
	Barbscontrol
	Debugcontrol
	Itemset1
	Itemset2
	Itemset3
	Itemset4
	Itemset5
	Itemset6
	Runeword
	Restinpeace
	Corpseexp
	Whirlwind
	Fullsetgeneric
	Monsterset
	Delerium
	Antidote
	Thawing
	Staminapot
	PassiveResistfire
	PassiveResistcold
	PassiveResistltng
	Uberminion
	Cooldown
	Sharedstash
	Hidedead
)
