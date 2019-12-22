package stacksandqueues

// Three in One:
// Describe how you could use a single array to implement three stacks

type MultiStack struct {
	stacksCount int
	sizes       []int
	mem         []int
}

func (m *MultiStack) Init(arraySize int, stacksCount int) {
	if arraySize < stacksCount {
		arraySize = stacksCount
	}

	m.mem = make([]int, arraySize)
	m.stacksCount = stacksCount
	m.sizes = make([]int, stacksCount)
	for i := range m.sizes {
		m.sizes[i] = 0
	}
}

func (m MultiStack) isMemoryFull() bool {
	sum := 0
	for i := range m.sizes {
		sum += m.sizes[i]
	}

	return sum == len(m.sizes)
}

func (m MultiStack) getStackIndex(stack int) int {
	if stack > m.stacksCount {
		return 0
	}

	if stack == 0 {
		return 0
	}

	position := 0
	for i := 0; i < stack-1; i++ {
		position += m.sizes[i]
	}

	return position
}

func (m *MultiStack) shiftLeft(stack int) int {
	start := m.getStackIndex(stack)
	for i := start; i < len(m.mem)-1; i++ {
		m.mem[i] = m.mem[i+1]
	}

	m.mem[len(m.mem)-1] = 0
	return start
}

func (m *MultiStack) shiftRight(stack int) int {
	end := m.getStackIndex(stack)
	for i := len(m.mem) - 1; i > end; i-- {
		m.mem[i] = m.mem[i-1]
	}

	m.mem[end] = 0
	return end
}

func (m *MultiStack) Push(stack int, value int) *int {
	if m.isMemoryFull() {
		return nil
	}

	if stack > m.stacksCount {
		return nil
	}

	position := m.shiftRight(stack)
	m.mem[position] = value
	m.sizes[stack-1]++
	return &m.mem[position]
}

func (m *MultiStack) Pop(stack int) int {
	position := m.getStackIndex(stack)
	item := m.mem[position]
	m.shiftLeft(stack)
	m.sizes[stack-1]--
	return item
}
