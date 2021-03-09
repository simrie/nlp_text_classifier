package utils

/*
MapKeysAsStrings retruns a slice of string keys from a map[string]
https://stackoverflow.com/questions/21362950/getting-a-slice-of-keys-from-a-map
*/
func MapKeysAsStrings(m map[string]interface{}) []string {
	// TODO: consider if this can be used to
	// replace MapKeysAsString for Block and WordsSeen
	keys := make([]string, 0, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

/*
ContainsString returns true if item string is in string slice
https://stackoverflow.com/questions/10485743/contains-method-for-a-slice
*/
func ContainsString(slice []string, item string) bool {
	set := MapStringSlice(slice)
	_, ok := set[item]
	return ok
}

/*
MapStringSlice to a map of empty structs
*/
func MapStringSlice(slice []string) map[string]struct{} {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	return set
}

/*
KeysInCommon returns an array of strings that are in both arrays
*/
func KeysInCommon(a1, a2 []string) []string {
	var keysInCommon []string
	setA2 := MapStringSlice(a2)
	for _, a1Str := range a1 {
		_, ok := setA2[a1Str]
		if ok {
			keysInCommon = append(keysInCommon, a1Str)
		}
	}
	return keysInCommon
}

/*
KeysDiff returns an array of strings in a1 but not in a2,
   and an array of strings in a2 but not in a2
*/
func KeysDiff(a1, a2 []string) ([]string, []string) {
	var onlyInA1 []string
	var onlyInA2 []string
	var keysInCommon = KeysInCommon(a1, a2)
	setInCommon := MapStringSlice(keysInCommon)

	for _, a1Str := range a1 {
		_, ok := setInCommon[a1Str]
		if !ok {
			onlyInA1 = append(onlyInA1, a1Str)
		}
	}
	for _, a2Str := range a2 {
		_, ok := setInCommon[a2Str]
		if !ok {
			onlyInA2 = append(onlyInA2, a2Str)
		}
	}
	return onlyInA1, onlyInA2
}
