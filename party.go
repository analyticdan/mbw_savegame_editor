package main

import "os"

type Party struct {
	Id                          String
	Name                        String
	Flags                       UInt64
	MenuId                      Int32
	PartyTemplateId             Int32
	FactionId                   Int32
	Personality                 Int32
	DefaultBehavior             Int32
	CurrentBehavior             Int32
	DefaultBehaviorObjectId     Int32
	CurrentBehaviorObjectId     Int32
	InitialPositionX            Float
	InitialPositionY            Float
	TargetPositionX             Float
	TargetPositionY             Float
	PositionX                   Float
	PositionY                   Float
	PositionZ                   Float
	NumStacks                   Int32
	Stacks                      []PartyStack
	Bearing                     Float
	Renamed                     Bool
	ExtraText                   String
	Morale                      Float
	Hunger                      Float
	_unused1                    Float
	PatrolRadius                Float
	Initiative                  Float
	Helpfulness                 Float
	LabelVisible                Int32
	BanditAttraction            Float
	Marshall                    Int32
	IgnorePlayerTimer           Int64
	BannerMapIconId             Int32
	ExtraMapIconId              Int32
	ExtraMapIconUpDownDistance  Float
	ExtraMapIconUpDownFrequency Float
	ExtraMapIconRotateFrequency Float
	ExtraMapIconFadeFrequency   Float
	AttachedToPartyId           Int32
	_unused2                    Int32
	IsAttached                  Bool
	NumAttachedPartyIds         Int32
	AttachedPartyIds            []Int32
	NumParticleSystemIds        Int32
	ParticleSystemIds           []Int32
	Notes                       [16]Note
	NumSlots                    Int32
	Slots                       []Int64
}

func (party *Party) Read(file *os.File, gameVersion Int32) {
	party.Id.Read(file)
	party.Name.Read(file)
	party.Flags.Read(file)
	party.MenuId.Read(file)
	party.PartyTemplateId.Read(file)
	party.FactionId.Read(file)
	party.Personality.Read(file)
	party.DefaultBehavior.Read(file)
	party.CurrentBehavior.Read(file)
	party.DefaultBehaviorObjectId.Read(file)
	party.CurrentBehaviorObjectId.Read(file)
	party.InitialPositionX.Read(file)
	party.InitialPositionY.Read(file)
	party.TargetPositionX.Read(file)
	party.TargetPositionY.Read(file)
	party.PositionX.Read(file)
	party.PositionY.Read(file)
	party.PositionZ.Read(file)

	party.NumStacks.Read(file)
	party.Stacks = make([]PartyStack, party.NumStacks)
	for i := 0; i < len(party.Stacks); i++ {
		party.Stacks[i].Read(file)
	}

	party.Bearing.Read(file)
	party.Renamed.Read(file)
	party.ExtraText.Read(file)
	party.Morale.Read(file)
	party.Hunger.Read(file)
	party._unused1.Read(file)
	party.PatrolRadius.Read(file)
	party.Initiative.Read(file)
	party.Helpfulness.Read(file)
	party.LabelVisible.Read(file)
	party.BanditAttraction.Read(file)
	if (gameVersion >= 900 && gameVersion < 1000) || gameVersion >= 1020 {
		party.Marshall.Read(file)
	}
	party.IgnorePlayerTimer.Read(file)
	party.BannerMapIconId.Read(file)
	if gameVersion >= 1137 {
		party.ExtraMapIconId.Read(file)
		party.ExtraMapIconUpDownDistance.Read(file)
		party.ExtraMapIconUpDownFrequency.Read(file)
		party.ExtraMapIconRotateFrequency.Read(file)
		party.ExtraMapIconFadeFrequency.Read(file)
	}

	party.AttachedToPartyId.Read(file)
	if gameVersion >= 1162 {
		party._unused2.Read(file)
	}
	party.IsAttached.Read(file)
	party.NumAttachedPartyIds.Read(file)
	party.AttachedPartyIds = make([]Int32, party.NumAttachedPartyIds)
	for i := 0; i < len(party.AttachedPartyIds); i++ {
		party.AttachedPartyIds[i].Read(file)
	}

	party.NumParticleSystemIds.Read(file)
	party.ParticleSystemIds = make([]Int32, party.NumParticleSystemIds)
	for i := 0; i < len(party.ParticleSystemIds); i++ {
		party.ParticleSystemIds[i].Read(file)
	}

	for i := 0; i < len(party.Notes); i++ {
		party.Notes[i].Read(file)
	}

	party.NumSlots.Read(file)
	party.Slots = make([]Int64, party.NumSlots)
	for i := 0; i < len(party.Slots); i++ {
		party.Slots[i].Read(file)
	}
}
