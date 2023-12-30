package types

type SubArray []byte

type Array [][]byte

type ArrayJSON struct {
	Value Array `json:"value"`
}

type ArrayLenTable []int
