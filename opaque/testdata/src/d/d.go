package ride

type AssignmentStatus int

const (
	AssignmentStatusPending AssignmentStatus = iota
	AssignmentStatusAccepted
	AssignmentStatusRejected
	AssignmentStatusCancelled
)

type AssignmentFilterOpts struct {
	Status *AssignmentStatus
}

type AssignmentFilterOpt interface {
	configureFilter(*AssignmentFilterOpts)
}

type configureFilterFunc func(*AssignmentFilterOpts)

func (f configureFilterFunc) configureFilter(opt *AssignmentFilterOpts) {
	f(opt)
}

func WithAssignmentStatus(status AssignmentStatus) AssignmentFilterOpt {
	return configureFilterFunc(func(opt *AssignmentFilterOpts) {
		opt.Status = &status
	})
}
