package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func CloneGraph(node *Node) *Node {
	// Handle kasus khusus jika graf kosong
	if node == nil {
		return nil
	}

	// Gunakan map untuk melacak node yang sudah dikunjungi
	visited := make(map[*Node]*Node)

	// Panggil fungsi cloneGraph untuk memulai proses cloning
	return helper(node, visited)
}

func helper(node *Node, visited map[*Node]*Node) *Node {
	// Jika node sudah dikunjungi, kembalikan node baru yang telah dibuat sebelumnya
	if clonedNode, ok := visited[node]; ok {
		return clonedNode
	}

	// Buat node baru dengan nilai yang sama seperti node yang akan di-clone
	clonedNode := &Node{Val: node.Val}

	// Tandai node asli dan node baru sebagai sudah dikunjungi
	visited[node] = clonedNode

	// Loop melalui tetangga-tetangga dari node asli dan rekursif cloneGraph
	for _, neighbor := range node.Neighbors {
		clonedNeighbor := helper(neighbor, visited)
		clonedNode.Neighbors = append(clonedNode.Neighbors, clonedNeighbor)
	}

	return clonedNode
}
