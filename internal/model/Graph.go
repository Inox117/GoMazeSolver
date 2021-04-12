package model

type Graph struct {
	Nodes []Node
	Edges []Edge
}

func (g *Graph) getNbNodes() int {
	return len(g.Nodes)
}

func (g *Graph) getNbEdges() int {
	return len(g.Edges)
}

func CreateEmptyGraph() Graph {
	return Graph{}
}

func BuildFromMatrix(matrix [][]int) Graph {
	var nodes []Node
	var edges []Edge

	// TODO vérifier la consistance de la matrice
	// 0 Mur
	// 1 Chemin standard
	// 2 Entrée
	// 3 Sortie

	rowNumber := len(matrix)
	colNumber := len(matrix[0])

	//Build nodes
	for rowIndex:= 0; rowIndex < rowNumber; rowIndex++ {
		for colIndex := 0; colIndex < colNumber; colIndex++{
			point := matrix[rowIndex][colIndex]
			var id int
			// TODO Would be fun to use anonymous function
			switch point {
			case 0:
				continue
			case 1:
				id = 1 + (rowIndex * colNumber) + colIndex
				nodes = append(nodes, BuildNormalNode(id))
			case 2:
				id = 1 + (rowIndex * colNumber) + colIndex
				nodes = append(nodes, BuildEntryNode(id))
			case 3:
				id = 1 + (rowIndex * colNumber) + colIndex
				nodes = append(nodes, BuildExitNode(id))
			}
			// TODO Make it more understandable
			//build BackEdge
 			if colIndex > 0 && matrix[rowIndex][colIndex - 1] > 0 {
				edges = append(edges, Edge{
					from: id,
					to:	id - 1,
				})
			}
			//build FrontEdge
			if colIndex < colNumber - 1  && matrix[rowIndex][colIndex + 1] > 0 {
				edges = append(edges, Edge{
					from: id,
					to:	id + 1,
				})
			}
			// build UpperEdge
			if rowIndex > 0 && matrix[rowIndex - 1][colIndex] > 0 {
				edges = append(edges, Edge{
					from: id,
					to:	id - rowNumber,
				})
			}
			// build LowerEdge
			if rowIndex < rowNumber - 1 && matrix[rowIndex + 1][colIndex] > 0 {
				edges = append(edges, Edge{
					from: id,
					to:	id + rowNumber,
				})
			}
		}
	}
	return Graph{
		Nodes: nodes,
		Edges: edges,
	}
}

func BuildFromEdgesAndNodes(nodes []Node, edges []Edge) Graph {
	return Graph{
		Nodes: nodes,
		Edges: edges,
	}
}