package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/ini.v1"
)

type SliceType interface {
	~string | ~int | ~float64 // add more *comparable* types as needed
}

func removeDuplicates[T SliceType](s []T) []T {
	if len(s) < 1 {
		return s
	}

	// sort
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	prev := 1
	for curr := 1; curr < len(s); curr++ {
		if s[curr-1] != s[curr] {
			s[prev] = s[curr]
			prev++
		}
	}

	return s[:prev]
}

func main() {

	mixFiles := make(map[string]string)

	files, err := ioutil.ReadDir("MERGE_MIX")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		b, err := os.ReadFile("MERGE_MIX/" + file.Name()) // just pass the file name
		if err != nil {
			fmt.Print(err)
		}
		mixFiles[file.Name()] = string(b)
	}

	err = os.RemoveAll("./temp")
	if err != nil {
		fmt.Printf("Fail to remove file: %v", err)
		os.Exit(1)
	}

	err = os.Mkdir("temp", 0755) // 0755 sets permissions for the directory
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rules, err := ini.LoadSources(ini.LoadOptions{
		UnparseableSections:     []string{},
		IgnoreInlineComment:     true,
		IgnoreContinuation:      true,
		SkipUnrecognizableLines: true,
	}, "MERGE_INI/rules.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	rulesCount, err := ini.LoadSources(ini.LoadOptions{
		UnparseableSections:     []string{},
		IgnoreInlineComment:     true,
		IgnoreContinuation:      true,
		SkipUnrecognizableLines: true,
	}, "INI/rules.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	ParticlesCurrent, _ := rulesCount.GetSection("Particles")
	ParticleSystemsCurrent, _ := rulesCount.GetSection("ParticleSystems")
	WarheadsCurrent, _ := rulesCount.GetSection("Warheads")
	AnimationsCurrent, _ := rulesCount.GetSection("Animations")
	BuildingTypesCurrent, _ := rulesCount.GetSection("BuildingTypes")
	AircraftTypesCurrent, _ := rulesCount.GetSection("AircraftTypes")
	VehicleTypesCurrent, _ := rulesCount.GetSection("VehicleTypes")
	InfantryTypesCurrent, _ := rulesCount.GetSection("InfantryTypes")

	var maxParticleKey int
	for _, key := range ParticlesCurrent.Keys() {
		i, _ := strconv.Atoi(key.Name())
		if i > maxParticleKey && i != 999 {
			maxParticleKey = i
		}
	}

	var maxParticleSystemKey int
	for _, key := range ParticleSystemsCurrent.Keys() {
		i, _ := strconv.Atoi(key.Name())
		if i > maxParticleSystemKey && i != 999 {
			maxParticleSystemKey = i
		}
	}

	var maxWarheadKey int
	for _, key := range WarheadsCurrent.Keys() {
		i, _ := strconv.Atoi(key.Name())
		if i > maxWarheadKey && i != 999 {
			maxWarheadKey = i
		}
	}

	var maxInfantryKey int
	for _, key := range InfantryTypesCurrent.Keys() {
		i, _ := strconv.Atoi(key.Name())
		if i > maxInfantryKey && i != 999 {
			maxInfantryKey = i
		}
	}

	var maxAnimationKey int
	for _, key := range AnimationsCurrent.Keys() {
		i, _ := strconv.Atoi(key.Name())
		if i > maxAnimationKey && i != 999 {
			maxAnimationKey = i
		}
	}

	var maxBuildingKey int
	for _, key := range BuildingTypesCurrent.Keys() {
		i, _ := strconv.Atoi(key.Name())
		if i > maxBuildingKey && i != 999 {
			maxBuildingKey = i
		}
	}

	var maxAircraftKey int
	for _, key := range AircraftTypesCurrent.Keys() {
		i, _ := strconv.Atoi(key.Name())
		if i > maxAircraftKey && i != 999 {
			maxAircraftKey = i
		}
	}

	var maxVehicleKey int
	for _, key := range VehicleTypesCurrent.Keys() {
		i, _ := strconv.Atoi(key.Name())
		if i > maxVehicleKey && i != 999 {
			maxVehicleKey = i
		}
	}

	Particles, _ := rules.GetSection("Particles")
	ParticleSystems, _ := rules.GetSection("ParticleSystems")
	Warheads, _ := rules.GetSection("Warheads")
	Animations, _ := rules.GetSection("Animations")
	BuildingTypes, _ := rules.GetSection("BuildingTypes")
	AircraftTypes, _ := rules.GetSection("AircraftTypes")
	VehicleTypes, _ := rules.GetSection("VehicleTypes")
	InfantryTypes, _ := rules.GetSection("InfantryTypes")

	sound, err := ini.LoadSources(ini.LoadOptions{
		UnparseableSections:     []string{},
		IgnoreInlineComment:     true,
		IgnoreContinuation:      true,
		SkipUnrecognizableLines: true,
	}, "MERGE_INI/sound01.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	art, err := ini.LoadSources(ini.LoadOptions{
		UnparseableSections:     []string{},
		IgnoreInlineComment:     true,
		IgnoreContinuation:      true,
		SkipUnrecognizableLines: true,
	}, "MERGE_INI/art.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	soundOutputCheck, err := ini.LoadSources(ini.LoadOptions{
		UnparseableSections:     []string{},
		IgnoreInlineComment:     true,
		IgnoreContinuation:      true,
		SkipUnrecognizableLines: true,
	}, "INI/sound01.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	soundOutput, err := ini.Load([]byte(""))
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	// var arr map[string]interface{}
	for _, value := range rules.Sections() {
		if value.Name() == "GFSTNK" {
			fmt.Println(value.Name())
			rules2, arts, sounds := FindItems(value, art, sound, rules)
			fmt.Println("Rules", removeDuplicates(rules2))
			fmt.Println("Arts", removeDuplicates(arts))
			fmt.Println("Sounds", removeDuplicates(sounds))

			rules2 = append(rules2, value.Name())
			arts = append(arts, value.Name())

			list, _ := soundOutput.NewSection("SoundList")
			sourceList, _ := soundOutputCheck.GetSection("SoundList")
			var maxKey int
			for _, key := range sourceList.Keys() {
				i, _ := strconv.Atoi(key.Name())
				if i > maxKey && i != 999 {
					maxKey = i
				}
			}

			for _, item := range removeDuplicates(sounds) {
				if soundOutputCheck.HasSection(item) {
					continue
				}
				println("Adding sound", item)

				newSection, err := sound.GetSection(item)
				if err != nil {
					fmt.Printf("no section: %v", err)
					os.Exit(1)
				}
				s, err := soundOutput.NewSection(item)
				if err != nil {
					fmt.Printf("no section: %v", err)
					os.Exit(1)
				}
				for _, key := range newSection.KeyStrings() {
					s.NewKey(key, newSection.Key(key).Value())

					// _, err := s.g
					// if err != nil {
					// 	fmt.Printf("no section: %v", err)
					// 	os.Exit(1)
					// }
				}
				maxKey++
				list.NewKey(fmt.Sprintf("%d", maxKey), item)

				soundOutput.SaveToIndent("temp/sound01_copy.ini", "")
			}

			artOutputCheck, err := ini.LoadSources(ini.LoadOptions{
				UnparseableSections:     []string{},
				IgnoreInlineComment:     true,
				IgnoreContinuation:      true,
				SkipUnrecognizableLines: true,
			}, "INI/art.ini")
			if err != nil {
				fmt.Printf("Fail to read file: %v", err)
				os.Exit(1)
			}

			rulesOutput, err := ini.LoadSources(ini.LoadOptions{
				UnparseableSections:     []string{},
				IgnoreInlineComment:     true,
				IgnoreContinuation:      true,
				SkipUnrecognizableLines: true,
			}, []byte(""))
			if err != nil {
				fmt.Printf("Fail to read file: %v", err)
				os.Exit(1)
			}

			ParticlesOutput, _ := rulesOutput.NewSection("Particles")
			ParticleSystemsOutput, _ := rulesOutput.NewSection("ParticleSystems")
			WarheadsOutput, _ := rulesOutput.NewSection("Warheads")
			AnimationsOutput, _ := rulesOutput.NewSection("Animations")
			BuildingTypesOutput, _ := rulesOutput.NewSection("BuildingTypes")
			AircraftTypesOutput, _ := rulesOutput.NewSection("AircraftTypes")
			VehicleTypesOutput, _ := rulesOutput.NewSection("VehicleTypes")
			InfantryTypesOutput, _ := rulesOutput.NewSection("InfantryTypes")

			artOutput, err := ini.LoadSources(ini.LoadOptions{
				UnparseableSections:     []string{},
				IgnoreInlineComment:     true,
				IgnoreContinuation:      true,
				SkipUnrecognizableLines: true,
			}, []byte(""))
			if err != nil {
				fmt.Printf("Fail to read file: %v", err)
				os.Exit(1)
			}

			for _, item := range removeDuplicates(arts) {

				if Animations.HasValue(item) && !AnimationsCurrent.HasValue(item) {
					maxAnimationKey++
					AnimationsOutput.NewKey(fmt.Sprintf("%d", maxAnimationKey), item)
				}

				for name, value := range mixFiles {
					if strings.Contains(value, item) || strings.Contains(value, strings.ToLower(item)) {
						fmt.Printf("Found %s in %s \n", item, name)
					}
				}

				if artOutputCheck.HasSection(item) {
					continue
				}

				newSection, err := art.GetSection(item)
				if err != nil {
					fmt.Printf("no section: %v", err)
					os.Exit(1)
				}
				s, err := artOutput.NewSection(item)
				if err != nil {
					fmt.Printf("no section: %v", err)
					os.Exit(1)
				}
				for _, key := range newSection.KeyStrings() {
					s.NewKey(key, newSection.Key(key).Value())

					// _, err := s.g
					// if err != nil {
					// 	fmt.Printf("no section: %v", err)
					// 	os.Exit(1)
					// }
				}

				artOutput.SaveToIndent("temp/art_copy.ini", "")
			}

			rulesOutputCheck, err := ini.LoadSources(ini.LoadOptions{
				UnparseableSections:     []string{},
				IgnoreInlineComment:     true,
				IgnoreContinuation:      true,
				SkipUnrecognizableLines: true,
			}, "INI/rules.ini")
			if err != nil {
				fmt.Printf("Fail to read file: %v", err)
				os.Exit(1)
			}

			for _, item := range removeDuplicates(rules2) {

				for name, value := range mixFiles {
					if strings.Contains(value, item) || strings.Contains(value, strings.ToLower(item)) {
						fmt.Printf("Found %s in %s \n", item, name)
					}
				}

				if rulesOutputCheck.HasSection(item) {
					continue
				}

				var isAnimation bool
				var isBuilding bool
				var isVehicle bool
				var isAircraft bool
				var isInfantry bool
				var isWarhead bool
				var isParticleSystem bool
				var isParticle bool

				for _, a := range Particles.Keys() {
					v, _ := Particles.GetKey(a.Name())
					if v.Value() == item {
						isParticle = true
						maxParticleKey++
						ParticlesOutput.NewKey(fmt.Sprintf("%d", maxParticleKey), item)
					}
				}

				for _, a := range ParticleSystems.Keys() {
					v, _ := ParticleSystems.GetKey(a.Name())
					if v.Value() == item {
						isParticleSystem = true
						maxParticleSystemKey++
						ParticleSystemsOutput.NewKey(fmt.Sprintf("%d", maxParticleSystemKey), item)
					}
				}

				for _, a := range Warheads.Keys() {
					v, _ := Warheads.GetKey(a.Name())
					if v.Value() == item {
						isWarhead = true
						maxWarheadKey++
						WarheadsOutput.NewKey(fmt.Sprintf("%d", maxWarheadKey), item)
					}
				}

				for _, a := range InfantryTypes.Keys() {
					v, _ := InfantryTypes.GetKey(a.Name())
					if v.Value() == item {
						isInfantry = true
						maxInfantryKey++
						InfantryTypesOutput.NewKey(fmt.Sprintf("%d", maxInfantryKey), item)
					}
				}

				for _, a := range Animations.Keys() {
					v, _ := Animations.GetKey(a.Name())
					if v.Value() == item {
						isAnimation = true
						maxAnimationKey++
						AnimationsOutput.NewKey(fmt.Sprintf("%d", maxAnimationKey), item)
					}
				}

				for _, a := range BuildingTypes.Keys() {
					v, _ := BuildingTypes.GetKey(a.Name())
					if v.Value() == item {
						isBuilding = true
						maxBuildingKey++
						BuildingTypesOutput.NewKey(fmt.Sprintf("%d", maxBuildingKey), item)
					}
				}

				for _, a := range VehicleTypes.Keys() {
					v, _ := VehicleTypes.GetKey(a.Name())
					if v.Value() == item {
						isVehicle = true
						maxVehicleKey++
						VehicleTypesOutput.NewKey(fmt.Sprintf("%d", maxVehicleKey), item)
					}
				}

				for _, a := range AircraftTypes.Keys() {
					v, _ := AircraftTypes.GetKey(a.Name())
					if v.Value() == item {
						isAircraft = true
						maxAircraftKey++
						AircraftTypesOutput.NewKey(fmt.Sprintf("%d", maxAircraftKey), item)
					}
				}

				if !isAnimation && !isBuilding && !isVehicle && !isAircraft && !isInfantry && !isWarhead && !isParticleSystem && !isParticle {
					fmt.Printf("WARNING: unknown type: %v \n", item)
				}

				newSection, err := rules.GetSection(item)
				if err != nil {
					fmt.Printf("no section: %v", err)
					os.Exit(1)
				}
				s, err := rulesOutput.NewSection(item)
				if err != nil {
					fmt.Printf("no section: %v", err)
					os.Exit(1)
				}
				for _, key := range newSection.KeyStrings() {
					s.NewKey(key, newSection.Key(key).Value())

					// _, err := s.g
					// if err != nil {
					// 	fmt.Printf("no section: %v", err)
					// 	os.Exit(1)
					// }
				}

				rulesOutput.SaveToIndent("temp/rules_copy.ini", "")
			}
		}
	}
}

var checkedRules map[string]bool
var checkedArt map[string]bool
var checkedSound map[string]bool

func init() {
	checkedRules = make(map[string]bool)
	checkedArt = make(map[string]bool)
	checkedSound = make(map[string]bool)
}

func FindItems(section *ini.Section, art *ini.File, sound *ini.File, rules *ini.File) ([]string, []string, []string) {
	re := regexp.MustCompile(`^[a-zA-Z-_0-9,]*$`)
	re2 := regexp.MustCompile(`[A-Z]+`)

	var arts []string
	var sounds []string
	var rules2 []string

	for _, value := range section.Keys() {
		if value.Name() == "Dock" {
			continue
		}
		if value.Name() == "Prerequisite" {
			continue
		}
		if value.Name() == "Owner" {
			continue
		}
		if re.MatchString(value.Value()) && re2.MatchString(value.Value()) {
			items := value.Strings(",")
			for _, item := range items {
				item := strings.Trim(item, " ")
				if sound.HasSection(item) && !checkedSound[item] {
					checkedSound[item] = true
					sounds = append(sounds, item)
					rules3, arts3, sounds3 := FindItems(sound.Section(item), art, sound, rules)
					rules2 = append(rules2, rules3...)
					arts = append(arts, arts3...)
					sounds = append(sounds, sounds3...)
				}
				if art.HasSection(item) && !checkedArt[item] {
					checkedArt[item] = true
					arts = append(arts, item)
					rules3, arts3, sounds3 := FindItems(art.Section(item), art, sound, rules)
					rules2 = append(rules2, rules3...)
					arts = append(arts, arts3...)
					sounds = append(sounds, sounds3...)
				}
				if rules.HasSection(item) && !checkedRules[item] {
					checkedRules[item] = true
					rules2 = append(rules2, item)
					rules3, arts3, sounds3 := FindItems(rules.Section(item), art, sound, rules)
					rules2 = append(rules2, rules3...)
					arts = append(arts, arts3...)
					sounds = append(sounds, sounds3...)
				}
			}
		}
	}

	return rules2, arts, sounds
}
