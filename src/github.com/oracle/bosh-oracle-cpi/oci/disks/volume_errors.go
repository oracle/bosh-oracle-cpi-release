package disks

import "fmt"

type VolumeTerminatedError struct {
	id string
}

func (e VolumeTerminatedError) Error() string {
	return fmt.Sprintf("Volume %s is terminated", e.id)
}

type VolumeFaultyError struct {
	id string
}

func (e VolumeFaultyError) Error() string {
	return fmt.Sprintf("Volume %s is faulty", e.id)
}
