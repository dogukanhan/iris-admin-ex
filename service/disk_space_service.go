package service

import (
	"app/environment"
	"app/utility"
)

type DiskSpaceService interface {
	GetDiskSpace() (string,error)
}

type DiskSpace struct {

}

type DiskSpaceMock struct {
}

func NewDiskSpaceService(env environment.Env) DiskSpaceService {
	switch env {
	case environment.PROD:
		return &DiskSpace{}
	case environment.DEV:
		return &DiskSpaceMock{}
	default:
		panic("unknown environment")
	}
}

func (d DiskSpace) GetDiskSpace() (string, error) {

	return utility.RunCommand("df")
}

func (d DiskSpaceMock) GetDiskSpace() (string, error) {

	return "/dev/mapper/debian9  49163408 34652364  12310296  74% ", nil
}
