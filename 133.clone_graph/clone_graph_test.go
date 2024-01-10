package clone_graph_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/133.clone_graph"
)

type testTable struct {
	name   string
	input  Node
	output Node
}

func TestCloneGraph(t *testing.T) {

	// testTable := []testTable{
	// 	{
	// 		name: "test 1",
	// 		input : &Node{
	// 			Val: 1,
	// 			Neighbors: []Node{Node{Val: 2}, Node{Val: 4}},
	// 		}
	// 	},
	// }

	// for _, tt := range testTable {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		assert.Equal(t, tt.output, CloneGraph(nil))
	// 	})
	// }
}
