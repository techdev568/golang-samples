package main

import "fmt"

func main() {
	/* Maps are golang builtin datatype similar to the hash table which maps a key to a value.
	Map is an unordered collection where each key is unique while values can be the same for two or more different keys.
	The advantages of using a map are that it provides fast retrieval, search, insert, and delete operations.
	Maps are referenced data types. When you assign one map to another both refer to the same underlying map
	*/

	/* go maps are not safe for concurrent(go routines) use.
	so better to use mutex. lock and unlock
	*/
	fmt.Println("Sample Map")
	sampleMapWithLength()
	fmt.Println("Declaration map using make")
	declareMapUsingMake()
	updateMap()
	daleteMap()
	checkMapKey()
	referenceMap()
	iterateMap()
	sampleMap := map[string]int{
		"sal":   1000,
		"phone": 9910,
	}

	fmt.Println(getAllKeys(sampleMap))
}

func sampleMapWithLength() {
	sampleMap := map[string]int{}
	fmt.Println("Sample Map: ", sampleMap)
	sampleMap = map[string]int{
		"sal":   1000,
		"phone": 9910,
	}
	fmt.Println("Map length ", len(sampleMap))

	sampleMap["hallticket"] = 3214
	fmt.Println("after adding data: ", sampleMap)
}

func declareMapUsingMake() {
	employeeSalary := make(map[string]int)
	//Adding a key value
	employeeSalary["Tom"] = 2000
	fmt.Println(employeeSalary)
}

func updateMap() {
	sampleMap := map[string]int{
		"sal":   1000,
		"phone": 9910,
	}
	fmt.Println("Sample Map: ", sampleMap)
	// update map value
	sampleMap["sal"] = 2000
	fmt.Println("Sample Map after update: ", sampleMap)
}

func daleteMap() {
	sampleMap := map[string]int{
		"sal":   1000,
		"phone": 9910,
	}
	fmt.Println("Sample Map: ", sampleMap)
	// update map value
	delete(sampleMap, "sal")
	fmt.Println("Sample Map after delete: ", sampleMap)
}

func checkMapKey() {
	sampleMap := map[string]int{
		"sal":   1000,
		"phone": 9910,
	}
	fmt.Println("Sample Map: ", sampleMap)
	// update map value
	val, ok := sampleMap["sal"]
	fmt.Println("check value and status ", val, ok)

	val, ok = sampleMap["sl"]
	fmt.Println("check value and status wrong ", val, ok)
}

func referenceMap() {
	sampleMap := map[string]int{
		"sal":   1000,
		"phone": 9910,
	}
	fmt.Println("Sample Map: ", sampleMap)
	sm := sampleMap
	fmt.Println("SM : ", sm)
	// update map value
	delete(sampleMap, "sal")
	fmt.Println("Sample Map after delete: ", sampleMap)
	fmt.Println("Sample Map after delete: ", sm)
}

func iterateMap() {
	sampleMap := map[string]int{
		"sal":   1000,
		"phone": 9910,
	}
	for k, v := range sampleMap {
		fmt.Println("key : value: \n", k, v)
	}
}

func getAllKeys(sample map[string]int) []string {
	var keys []string
	for k := range sample {
		keys = append(keys, k)
	}
	return keys
}
