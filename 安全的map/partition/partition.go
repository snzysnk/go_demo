package partition

import "go_demo/安全的map/mp"

const (
	partitionNumber = 6
)

var (
	partitionBlocks = make([]mp.IXMap, partitionNumber)
)

func init() {
	for i := 0; i < partitionNumber; i++ {
		partitionBlocks[i] = mp.CreateSafeXMap()
	}
}

func getBlock(name string) mp.IXMap {
	return partitionBlocks[DJB([]byte(name))%partitionNumber]
}

func Set(key, value string) {
	getBlock(key).Set(key, value)
}

func Get(key string) string {
	return getBlock(key).Get(key)
}

// DJB 使用DJB进行分区
func DJB(str []byte) uint32 {
	var hash uint32 = 5381
	for i := 0; i < len(str); i++ {
		hash += (hash << 5) + uint32(str[i])
	}
	return hash
}
