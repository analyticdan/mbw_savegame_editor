package main

const (
	Sargoth = 21 + iota
	Tihr
	Veluca
	Suno
	Jelkala
	Praven
	Uxkhal
	Reyvadin
	Khudan
	Tulga
	Curaw
	Wercheg
	Rivacheg
	Halmar
	Yalen
	Dhirim
	Ichamur
	Narra
	Shariz
	Durquba
	Ahmerrad
	Bariyye
)

var TownIds = []int{
	Sargoth,
	Tihr,
	Veluca,
	Suno,
	Jelkala,
	Praven,
	Uxkhal,
	Reyvadin,
	Khudan,
	Tulga,
	Curaw,
	Wercheg,
	Rivacheg,
	Halmar,
	Yalen,
	Dhirim,
	Ichamur,
	Narra,
	Shariz,
	Durquba,
	Ahmerrad,
	Bariyye,
}

const (
	CulmarrCastle = 43 + iota
	MalayurgCastle
	BulughaCastle
	RadoghirCastle
	TehlrogCastle
	TilbautCastle
	SungetcheCastle
	JeirbeCastle
	JamicheCastle
	AlburqCastle
	CurinCastle
	ChalbekCastle
	KelredanCastle
	MarasCastle
	ErgellonCastle
	AlmerraCastle
	DistarCastle
	IsmiralaCastle
	YrumaCastle
	DerchiosCastle
	IbdelesCastle
	UnuzdaqCastle
	TevarinCastle
	ReindiCastle
	RyibeletCastle
	SenuzgdaCastle
	RindyarCastle
	GrunwalderCastle
	NelagCastle
	AsuganCastle
	VyincourdCastle
	KnudarrCastle
	EtrosqCastle
	HrusCastle
	HaringothCastle
	JelbegiCastle
	DramugCastle
	TulbukCastle
	SlezkhCastle
	UhhunCastle
	JameyyedCastle
	TerammaCastle
	SharwaCastle
	DurrinCastle
	CarafCastle
	WeyyahCastle
	SamarraCastle
	BardaqCastle
)

var CastleIds = []int{
	CulmarrCastle,
	MalayurgCastle,
	BulughaCastle,
	RadoghirCastle,
	TehlrogCastle,
	TilbautCastle,
	SungetcheCastle,
	JeirbeCastle,
	JamicheCastle,
	AlburqCastle,
	CurinCastle,
	ChalbekCastle,
	KelredanCastle,
	MarasCastle,
	ErgellonCastle,
	AlmerraCastle,
	DistarCastle,
	IsmiralaCastle,
	YrumaCastle,
	DerchiosCastle,
	IbdelesCastle,
	UnuzdaqCastle,
	TevarinCastle,
	ReindiCastle,
	RyibeletCastle,
	SenuzgdaCastle,
	RindyarCastle,
	GrunwalderCastle,
	NelagCastle,
	AsuganCastle,
	VyincourdCastle,
	KnudarrCastle,
	EtrosqCastle,
	HrusCastle,
	HaringothCastle,
	JelbegiCastle,
	DramugCastle,
	TulbukCastle,
	SlezkhCastle,
	UhhunCastle,
	JameyyedCastle,
	TerammaCastle,
	SharwaCastle,
	DurrinCastle,
	CarafCastle,
	WeyyahCastle,
	SamarraCastle,
	BardaqCastle,
}

const (
	Yaragar = 91 + iota
	Burglen
	Azgad
	Nomar
	Kulum
	Emirin
	Amere
	Haen
	Buvran
	Mechin
	Dusturil
	Emer
	Nemeja
	Sumbuja
	Ryibelet
	Shapeshte
	Mazen
	Ulburban
	Hanun
	Uslum
	Bazeck
	Shulus
	Ilvia
	Ruldi
	Dashbigha
	Pagundur
	Glunmar
	TashKulun
	Buillin
	Ruvar
	Ambean
	Tosdhar
	Ruluns
	Ehlerdah
	Fearichen
	Jayek
	AdaKulun
	Ibiran
	Reveran
	Saren
	Dugan
	DirighAban
	Zagush
	Peshmi
	Bulugur
	Fedner
	Epeshe
	Veidar
	Tismirr
	Karindi
	Jelbegi
	Amashke
	Balanli
	Chide
	Tadsamesh
	Fenada
	Ushkuru
	Vezin
	Dumar
	Tahlberl
	Aldelen
	Rebache
	Rduna
	Serindiar
	Iyindah
	Fisdnar
	Tebandra
	Ibdeles
	Kwynn
	Dirigsene
	Tshibtin
	Elberl
	Chaeza
	Ayyike
	Bhulaban
	Kedelke
	Rizi
	Sarimish
	Istiniar
	Vayejeg
	Odasan
	Yalibe
	Gisim
	Chelez
	Ismirala
	Slezkh
	Udiniad
	Tulbuk
	Uhhun
	Jamiche
	AynAssuadi
	Dhibbain
	Qalyut
	Mazigh
	Tamnuh
	Habba
	Sekhtem
	Mawiti
	Fishara
	Iqbayl
	Uzgha
	ShibalZumr
	Mijayet
	Tazjunat
	Aab
	Hawaha
	Unriya
	MitNun
	Tilimsal
	Rushdigh
)

var VillageIds = []int{
	Yaragar,
	Burglen,
	Azgad,
	Nomar,
	Kulum,
	Emirin,
	Amere,
	Haen,
	Buvran,
	Mechin,
	Dusturil,
	Emer,
	Nemeja,
	Sumbuja,
	Ryibelet,
	Shapeshte,
	Mazen,
	Ulburban,
	Hanun,
	Uslum,
	Bazeck,
	Shulus,
	Ilvia,
	Ruldi,
	Dashbigha,
	Pagundur,
	Glunmar,
	TashKulun,
	Buillin,
	Ruvar,
	Ambean,
	Tosdhar,
	Ruluns,
	Ehlerdah,
	Fearichen,
	Jayek,
	AdaKulun,
	Ibiran,
	Reveran,
	Saren,
	Dugan,
	DirighAban,
	Zagush,
	Peshmi,
	Bulugur,
	Fedner,
	Epeshe,
	Veidar,
	Tismirr,
	Karindi,
	Jelbegi,
	Amashke,
	Balanli,
	Chide,
	Tadsamesh,
	Fenada,
	Ushkuru,
	Vezin,
	Dumar,
	Tahlberl,
	Aldelen,
	Rebache,
	Rduna,
	Serindiar,
	Iyindah,
	Fisdnar,
	Tebandra,
	Ibdeles,
	Kwynn,
	Dirigsene,
	Tshibtin,
	Elberl,
	Chaeza,
	Ayyike,
	Bhulaban,
	Kedelke,
	Rizi,
	Sarimish,
	Istiniar,
	Vayejeg,
	Odasan,
	Yalibe,
	Gisim,
	Chelez,
	Ismirala,
	Slezkh,
	Udiniad,
	Tulbuk,
	Uhhun,
	Jamiche,
	AynAssuadi,
	Dhibbain,
	Qalyut,
	Mazigh,
	Tamnuh,
	Habba,
	Sekhtem,
	Mawiti,
	Fishara,
	Iqbayl,
	Uzgha,
	ShibalZumr,
	Mijayet,
	Tazjunat,
	Aab,
	Hawaha,
	Unriya,
	MitNun,
	Tilimsal,
	Rushdigh,
}

const (
	Borcha = 194 + iota
	Marnid
	Ymira
	Rolf
	Baheshtur
	Firentis
	Deshavi
	Matheld
	Alayen
	Bunduk
	Katrin
	Jeremus
	Nizar
	Lezalit
	Artimenner
	Klethi
)

var CompanionIds = []int{
	Borcha,
	Marnid,
	Ymira,
	Rolf,
	Baheshtur,
	Firentis,
	Deshavi,
	Matheld,
	Alayen,
	Bunduk,
	Katrin,
	Jeremus,
	Nizar,
	Lezalit,
	Artimenner,
	Klethi,
}

var CompanionsNameMap = map[int]string{
	Borcha:     "Borcha",
	Marnid:     "Marnid",
	Ymira:      "Ymira",
	Rolf:       "Rolf",
	Baheshtur:  "Baheshtur",
	Firentis:   "Firentis",
	Deshavi:    "Deshavi",
	Matheld:    "Matheld",
	Alayen:     "Alayen",
	Bunduk:     "Bunduk",
	Katrin:     "Katrin",
	Jeremus:    "Jeremus",
	Nizar:      "Nizar",
	Lezalit:    "Lezalit",
	Artimenner: "Artimenner",
	Klethi:     "Klethi",
}

const (
	// See ek_.* in header_items.py
	EquippedItem0 = iota
	EquippedItem1
	EquippedItem2
	EquippedItem3
	EquippedHead
	EquippedBody
	EquippedFoot
	EquippedGloves
	EquippedHorse
)

const (
	// See fac_.* in ID_factions.py
	NoFaction = iota
	Commoners
	Outlaws
	Neutral
	Innocents
	Merchants
	_DarkKnights
	_Culture1
	_Culture2
	_Culture3
	_Culture4
	_Culture5
	_Culture6
	PlayerFaction
	PlayersSupporters
	KingdomOfSwadia
	KingdomOfVaegirs
	KhergitKhanate
	KingdomOfNords
	KingdomOfRhodoks
	SarranidSultanate
	_KingdomsEnd
	_RobberKnights
	_Khergits
	_BlackKhergits
	Manhunters
	Deserters
	MountainBandits
	ForestBandits
	_Undeads
	_Slavers
	_PeasantRebels
	_NobleRefugees
)

const (
	// See slot_village_state, svs_.* in module_constants.py
	VillageNormal = iota
	VillageBeingRaided
	VillageLooted
)

const (
	// See lrep_.* in module_constants.py
	None = iota
	Martial
	Quarrelsome
	SelfRighteous
	Cunning
	Debauched
	GoodNatured
	Upstanding
)

const (
	BookSeller1 = 186 + iota
	BookSeller2
)
