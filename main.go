package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "math"
  "os"
)


type Point struct {
  x int
  y int
}

type Graph interface {
  Points() []Point
  Tsp(point Point) []Point
}

type Vgraph struct {
  points []Point
}

func newGraph(p_points [][]int) *Vgraph {
  var points []Point
  for _, p_point := range p_points {
    point := Point{p_point[0], p_point[1]}
    points = append(points, point)
  }
  return &Vgraph{points}
}

/*
  Returns an ordered list of points in which the total distance traversed by
  visiting each point in the order of the list is a minimum
*/
func (g Vgraph) Tsp() []Point {
  path, _ := tsp_helper(g.Points())
  return path
}

func (v Vgraph) Points() []Point {
  tmp := make([]Point, len(v.points))
  copy(tmp, v.points)
  return tmp
}
/*

  */
func tsp_helper(points []Point) ([]Point, float64) {
  min := math.MaxFloat64
  var min_path []Point
  for i, point := range points[1:] {
    fmt.Printf("before append")
    target_points := append(points[1:i], points[i+1:]...)
    fmt.Printf("after append")
    path, length := tsp_helper(target_points)
    // Pretty sure we need to find the best place to put this point, doesn't
    // make sense to put it in the beginning... TODO!
    dist := point.Distance(path[0])
    if dist + length <= min {
      min = dist + length
      min_path = append([]Point{point}, target_points...)
    }
  }
  return min_path, min
}
func (p Point) Distance (q Point) float64 {
  return math.Sqrt(math.Pow(math.Abs(float64(p.x - q.x)),2)+ math.Pow(math.Abs(float64(p.x - q.x)), 2))
}
func main() {

	points_file := os.Args[1]

	points_corpus, err := ioutil.ReadFile(points_file)
	if err != nil {
		log.Fatal(err)
	}
  var points [][]int
  err = json.Unmarshal(points_corpus, &points)
  fmt.Printf("%+v", points)
  vgraph := newGraph(points)
  tsp := vgraph.Tsp()
  fmt.Printf("%+v", tsp)
}
