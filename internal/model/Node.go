package model

type Node struct {
	Id int
	isEntry bool
	isExit bool
}

func BuildNormalNode(id int) Node {
	return Node{Id: id, isEntry: false, isExit: false}
}

func BuildEntryNode(id int) Node {
	return Node{Id: id, isEntry: true, isExit: false}
}

func BuildExitNode(id int) Node {
	return Node{Id: id, isEntry: false, isExit: true}
}