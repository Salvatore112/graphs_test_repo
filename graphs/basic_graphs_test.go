package graphs

import (
	"reflect"
	"sort"
	"testing"
)

func TestAddVertex(t *testing.T) {
	graph := NewSimpleGraph()

	graph.AddVertex("A")
	if _, exists := graph.vertices["A"]; !exists {
		t.Errorf("Expected vertex 'A' to be added, but it wasn't")
	}

	graph.AddVertex("A")
	if len(graph.vertices) != 1 {
		t.Errorf("Expected only one vertex 'A', but got %d vertices", len(graph.vertices))
	}
}

func TestAddEdge(t *testing.T) {
	graph := NewSimpleGraph()

	graph.AddEdge("A", "B")
	if _, exists := graph.edges["A"]["B"]; !exists {
		t.Errorf("Expected edge between 'A' and 'B', but it doesn't exist")
	}
	if _, exists := graph.edges["B"]["A"]; !exists {
		t.Errorf("Expected edge between 'B' and 'A', but it doesn't exist")
	}

	graph.AddEdge("C", "D")
	if _, exists := graph.vertices["C"]; !exists {
		t.Errorf("Expected vertex 'C' to be added")
	}
	if _, exists := graph.vertices["D"]; !exists {
		t.Errorf("Expected vertex 'D' to be added")
	}
	if _, exists := graph.edges["C"]["D"]; !exists {
		t.Errorf("Expected edge between 'C' and 'D', but it doesn't exist")
	}
	if _, exists := graph.edges["D"]["C"]; !exists {
		t.Errorf("Expected edge between 'D' and 'C', but it doesn't exist")
	}
}

func TestGetNeighbors(t *testing.T) {
	graph := NewSimpleGraph()

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")

	neighborsA := graph.GetNeighbors("A")
	expectedNeighborsA := []string{"B", "C"}
	if !reflect.DeepEqual(neighborsA, expectedNeighborsA) {
		t.Errorf("Expected neighbors of 'A' to be %v, but got %v", expectedNeighborsA, neighborsA)
	}

	neighborsB := graph.GetNeighbors("B")
	expectedNeighborsB := []string{"A"}
	if !reflect.DeepEqual(neighborsB, expectedNeighborsB) {
		t.Errorf("Expected neighbors of 'B' to be %v, but got %v", expectedNeighborsB, neighborsB)
	}

	neighborsNonExistent := graph.GetNeighbors("NonExistent")
	if len(neighborsNonExistent) != 0 {
		t.Errorf("Expected no neighbors for non-existent vertex, but got %v", neighborsNonExistent)
	}
}

func TestGetVertices(t *testing.T) {
	graph := NewSimpleGraph()

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")

	vertices := graph.GetVertices()
	expectedVertices := []string{"A", "B", "C"}

	if !reflect.DeepEqual(sortStrings(vertices), sortStrings(expectedVertices)) {
		t.Errorf("Expected vertices to be %v, but got %v", expectedVertices, vertices)
	}
}

func sortStrings(slice []string) []string {
	sorted := append([]string(nil), slice...)
	sort.Strings(sorted)
	return sorted
}
