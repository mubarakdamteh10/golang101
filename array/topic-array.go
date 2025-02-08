package array

type IArrayTopic interface {
	SumArray(list []int) int
}

type arrayTopic struct {
	name string
}

func NewArrayTopic() IArrayTopic {
	return &arrayTopic{
		name: "ABC",
	}
}

func (arr *arrayTopic) SumArray(list []int) int {
	sum := 0
	for _, item := range list {
		sum += item
	}

	//
	return sum
}
