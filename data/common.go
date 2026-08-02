package data

/* For debugging purposes, as NaN cannot be compared or turned into JSON. */
var DisableNaN bool

const (
	/*TODO: Maybe move this to its own file with the other tf_flags. */
	heroFlag = UInt64(0x00000010)
	/*TODO: This should be obtained from module.ini. */
	loadRegularTroopInventory = false
)

func hasAdditionalInfo(playerParty Party, i int) bool {
	troopId := playerParty.Stacks[i].TroopId
	/*TODO: isHero should be calculated based on troop.txt module data.
	Below uses a hard-coded value based on Native.*/
	isHero := troopId == 0 || (troopId >= 194 && troopId < 463)
	return !isHero
}
