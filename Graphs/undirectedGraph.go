// Creates an undirected graph represented with an adjacency list
package main

type successors struct {
    sucNum      int
    isSorted    bool
    listOfSuc   []int
}

type uGraph struct {
    vertNum     int
    edgeNum     int
    suc         *successors
}

func (graph *uGraph) createGraph(verticies int) {
    graph.vertNum = verticies
    graph.edgeNum = 0
    graph.suc.sucNum = 0
    graph.suc.
}
