package customtype

type BranchPreference string

const (
	BranchPreference_QUEUES     BranchPreference = "queues"
	BranchPreference_DELIVERIES BranchPreference = "deliveries"
	BranchPreference_SPACES     BranchPreference = "spaces"
)
