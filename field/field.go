package field

import "github.com/karlek/sensou/temp"

// A single field of object editor values
type Field struct {
	Name  string
	Value interface{}
}

// Returns the field type, int and bool (0), fixed decimal (1), float(2), string(3)
// ###Add fixed
func (f *Field) FieldType() int {

	switch f.Value.(type) {
	case int:
		return 0x00
	case bool:
		return 0x01
	case float64:
		return 0x02
	case string:
		return 0x03
	}
	// panic("asdf")
	// Should never be able to come here because it will overflow byte
	// ###Make better solution
	return -1
}

// Converts the value of the field to a byte array which then can be written to a file
// Better name
func (f *Field) ValueToBuf() []byte {
	switch i := f.Value.(type) {
	case int:
		return temp.Binary(i).Little()

	case bool:
		if i == true {
			return []byte{10, 00, 00}
		}
		return []byte{00, 00, 00}

	case float64:
		return []byte{byte(i)}
	case string:
		return []byte(i)
	}

	// Should never be able to come here
	// Add err
	return nil
}

// ###Add fixed
// Retrive the field description with help from the field name and the type of the value
func (f *Field) GetDescriber() (describer string) {

	// Our map which contains field names as keys and descriptions as elements
	var fieldMap = &map[string]string{}

	// Depending on the type of the value, different maps will be used
	switch f.Value.(type) {
	case int:
		fieldMap = &IntValues
	case bool:
		fieldMap = &BoolValues
	case float64:
		fieldMap = &FloatValues
	case string:
		fieldMap = &StringValues
	}

	// Return the description if matching field key and raw code are found
	for rawCode, describer := range *fieldMap {
		if f.Name == rawCode && describer != "" {
			return describer
		}
	}

	// Otherwise simply return the old field name
	return f.Name
}

var BoolValues = map[string]string{
	"utcc": "Art - Allow Custom Team Color",
	"ushr": "Art - Has Water Shadow",
	"uscb": "Art - Scale Projectiles",
	"usew": "Art - Selection Circle On Water",
	"ulos": "Art - Use Extended Line of Sight",
	"umh1": "Combat - Attack 1 - Projectile Homing Enabled",
	"uwu1": "Combat - Attack 1 - Show Ul",
	"umh2": "Combat - Attack 2 - Projectile Homing Enabled",
	"uwu2": "Combat - Attack 2 - Show Ul",
}

var IntValues = map[string]string{
	"ubpx": "Art - Button Position (X)",
	"ubpy": "Art - Button Position (Y)",
	"uept": "Art - Elevation - Sample Points",
	"uver": "Art - Model File - Extra Versions",
	"uori": "Art - Orientation Interpolation",
	"utco": "Art - Team Color",
	"uclr": "Art - Tinting Color 1",
	"uclg": "Art - Tinting Color 2",
	"uclb": "Art - Tinting Color 3",
	"ua1f": "Combat - Attack 1 - Area of Effect (Full Damage)",
	"ua1h": "Combat - Attack 1 - Area of Effect (Medium Damage)",
	"ua1q": "Combat - Attack 1 - Area of Effect (Small Damage)",
	"ua1b": "Combat - Attack 1 - Damage Base",
	"ua1d": "Combat - Attack 1 - Damage Number of Dice",
	"ua1s": "Combat - Attack 1 - Damage Sides per Die",
	"udu1": "Combat - Attack 1 - Damage Upgrade Amount",
	"utc1": "Combat - Attack 1 - Maximum Number of Targets",
	"ua1z": "Combat - Attack 1 - Projectile Speed",
	"ua1r": "Combat - Attack 1 - Range",
	"ua2f": "Combat - Attack 2 - Area of Effect (Full Damage)",
	"ua2h": "Combat - Attack 2 - Area of Effect (Medium Damage)",
	"ua2q": "Combat - Attack 2 - Area of Effect (Small Damage)",
	"ua2b": "Combat - Attack 2 - Damage Base",
	"ua2d": "Combat - Attack 2 - Damage Number of Dice",
	"umh2": "",
	"ua2z": "",
	"ua2r": "",
	"uaen": "",
	"udea": "",
	"udef": "",
	"uamn": "",
	"udro": "",
	"ucam": "",
	"uhos": "",
	"uspe": "",
	"utss": "",
	"uine": "",
	"ubld": "",
	"ufle": "",
	"ufoo": "",
	"ufma": "",
	"ufor": "",
	"ubba": "",
	"ubdi": "",
	"ubsi": "",
	"ugol": "",
	"uhom": "",
	"uhpm": "",
	"ubdg": "",
	"ulev": "",
	"ulba": "",
	"ulbd": "",
	"ulbs": "",
	"ulum": "",
	"umpi": "",
	"umpm": "",
	"uibo": "",
	"ucbo": "",
	"unbm": "",
	"unbr": "",
	"ugor": "",
	"ulur": "",
	"urtm": "",
	"usid": "",
	"usin": "",
	"usle": "",
	"usma": "",
	"usrg": "",
	"usst": "",
}

var FixedDecimalValues = map[string]string{
	"uble": "Art - Animation - Blend Time",
	"urun": "Art - Animation - Run Speed",
	"uwal": "Art - Animation - Walk Speed",
	"uerd": "Art - Elevation - Sample Radius",
	"ufrd": "Art - Fog of War - Sample Radius",
	"umxp": "Art - Maximum Pitch Angle (degrees)",
	"umxr": "Art - Maximum Roll Angle (degrees)",
	"usca": "Art - Scaling Value",
	"uslz": "Art - Selection Circle - Height",
	"ussc": "Art - Selection Scale",
	"ushx": "Art - Shadow Image - CenterX",
	"ushy": "Art- Shadow Image - CenterY",
	"ushh": "Art- Shadow Image - Height",
	"ushw": "Art- Shadow Image - Width",
}

var FloatValues = map[string]string{
	"ucbs": "Art - Animation - Cast Backswing",
	"ucpt": "Art - Animation - Cast Point",
	"udtm": "Art - Death Time (seconds)",
	"uocc": "Art - Occluder Height",
	"uisz": "Art - Projectile Impact - Z (Swimming)",
	"uimz": "Art - Projectile Impact- Z",
	"ulpx": "Art - Projectile Launch - X",
	"ulpy": "Art - Projectile Launch - Y",
	"ulsz": "Art - Projectile Launch - Z (Swimming)",
	"ulpz": "Art - Projectile Launch - Z",
	"uprw": "Art - Propulsion Window (degrees)",
	"uacq": "Combat - Acquisition Range",
	"ubs1": "Combat - Attack 1 - Animation Backswing Point",
	"udp1": "Combat - Attack 1 - Animation Damage Point",
	"ua1c": "Combat - Attack 1 - Cooldown Time",
	"uhd1": "Combat - Attack 1 - Damage Factor - Medium",
	"uqd1": "Combat - Attack 1 - Damage Factor - Small",
	"udl1": "Combat - Attack 1 - Damage Loss Factor",
	"usd1": "Combat - Attack 1 - Damage Spill Distance",
	"usr1": "Combat - Attack 1 - Damage Spill Radius",
	"uma1": "Combat - Attack 1 - Projectile Arc",
	"urb1": "Combat - Attack 1 - Range Motion Buffer",
	"ubs2": "Combat - Attack 2 - Animation Backswing Point",
	"udp2": "Combat - Attack 2 - Animation Damage Point",
	"ua2c": "Combat - Attack 2 - Cooldown Time",
	"uhd2": "Combat - Attack 2 - Damage Factor - Medium",
	"uqd2": "Combat - Attack 2 - Damage Factor - Small",
	"udl2": "Combat - Attack 2 - Damage Loss Factor",
	"usd2": "",
	"usr2": "",
	"uma2": "",
	"urb2": "",
	"umvh": "",
	"umvf": "",
	"umvr": "",
	"uabr": "",
	"ucol": "",
	"uhpr": "",
	"umpr": "",
	"upaw": "",
}

var StringValues = map[string]string{
	"unam": "Text - Name",
	"udaa": "Abilities - Default Active Ability",
	"uabi": "Abilities - Normal",
	"ucua": "Art - Caster Upgrade Art",
	"uico": "Art - Icon - Game Interface",
	"ussi": "Art - Icon - Score Screen",
	"umdl": "Art - Model File",
	"uani": "Art - Required Animation Names",
	"uaap": "Art - Required Animation Names - Attachments",
	"ualp": "Art - Required Attachment Link Names",
	"ubpr": "Art - Required Bone Names",
	"ushu": "Art - Shadow Image (Unit)",
	"ushb": "Art- Shadow Texture (Building)",
	"uspa": "Art- Special",
	"utaa": "Art- Target",
	"uarm": "Combat - Armor Type",
	"ua1p": "Combat - Attack 1 - Area of Effect Targets",
	"ua1t": "Combat - Attack 1 - Attack Type",
	"ua1m": "Combat - Attack 1 - Projectile Art",
	"ua1g": "Combat - Attack 1 - Targets Allowed",
	"ucs1": "Combat - Attack 1 - Weapon Sound",
	"ua1w": "Combat - Attack 1 - Weapon Type",
	"ua2p": "Combat - Attack 2 - Area of Effect Targets",
	"ua2t": "Combat - Attack 2 - Attack Type",
	"ua2m": "",
	"ua2g": "",
	"ucs2": "",
	"ua2w": "",
	"udty": "",
	"utar": "",
	"util": "",
	"umvt": "",
	"uabt": "",
	"umsl": "",
	"ursl": "",
	"usnd": "",
	"uhrt": "",
	"urac": "",
	"uubs": "",
	"upat": "",
	"upar": "",
	"upap": "",
	"ubsl": "",
	"utyp": "",
	"udep": "",
	"umki": "",
	"usei": "",
	"ureq": "",
	"urqa": "",
	"ures": "",
	"useu": "",
	"utra": "",
	"uupt": "",
	"upgr": "",
	"ides": "",
	"uhot": "",
	"unsf": "",
	"utip": "",
	"utub": "",
}
