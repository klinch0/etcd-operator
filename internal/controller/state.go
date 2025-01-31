package controller

type State interface {
	ClusterExists() bool
}
