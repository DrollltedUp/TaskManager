package taskmanager

type Manager struct {
	Tasks    []Task
	NextID   int
	DataFile string
}
