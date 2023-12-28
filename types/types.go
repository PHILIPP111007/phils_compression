package types

type SubArray []byte

type Array struct {
	Value [][]byte `json:"value"`
}

type ArrayLenTable []int
