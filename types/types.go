package types

type Channel []byte

type Image struct {
	Red   Channel `json:"red"`
	Green Channel `json:"green"`
	Blue  Channel `json:"blue"`
}
