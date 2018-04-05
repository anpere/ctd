package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "math"
  "os"
)


type Stack struct {
  points []Point
}

func (s *Stack) Pop() Point{
  point := s.points[0]
  s.points = s.points[1:]
  return point
}

func (s *Stack) Push(p Point){
  s.points = append(s.points, p)
}

func (s *Stack) Top() Point {
    return s.points[0]
}

func (s *Stack) NextToTop() Point {
  return s.points[1]
}

func (s *Stack) All() []Point {
  ret_items := make([]Point, len(s.points))
  copy(ret_items, s.points)
  return ret_items
}

type Point struct {
  x float64
  y float64
}

type Graph interface {
  Points() []Point
  Tsp(point Point) []Point
}

type Vgraph struct {
  points []Point
}
var DEBUG bool = false;
func debug(format string, a... interface{}) {
  if DEBUG {
    fmt.Printf(format+"\n", a...)
  }
}
func newGraph(p_points [][]float64) *Vgraph {
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
func (g Vgraph) TspApprox() []Point {
  path, _ := tsp_helper(g.Points())
  return path
}

func (v Vgraph) Points() []Point {
  tmp := make([]Point, len(v.points))
  copy(tmp, v.points)
  return tmp
}

func (this Point) Equals(other Point) bool {
  return this.x == other.x && this.y == other.y
}
func polar(a Point, b Point) float64 {
  angle_a := math.Tan(float64(a.y)/float64(a.x))
  angle_b := math.Tan(float64(b.y)/float64(b.x))
  return angle_a - angle_b
}
func quicksortByPolar(points []Point, pivot Point) []Point {
  // returns a sorted list of points, sorted by polar angle with respect to the polar angle
  var left, right []Point
  if len(points) <= 1 {
    return points
  }
  sorting_pivot := points[len(points)/2]
  points_to_partition := remove(points, sorting_pivot)
  for _, point := range points_to_partition {
    if polar(pivot, point) <= polar(pivot, sorting_pivot) {
      left = append(left, point)
    } else {
      right = append(right, point)
    }
  }
  left = quicksortByPolar(left, pivot)
  right = quicksortByPolar(right, pivot)
  left = append(left, sorting_pivot)
  return append(left, right...)
}

func remove(s []Point, p Point) []Point {
    var index int
    for i, other := range s {
      if other.Equals(p){
        index = i
        break
      }
    }
    s[index] = s[len(s)-1]
    return s[:len(s)-1]
}

func (v Vgraph) ConvexHull() []Point {
  // Takes in a Vgraph
  // and returns the Convex Hull (in order defining the hull)
  // requires len(v.points) >= 3
  _points := make([]Point, len(v.points))
  copy(_points, v.points)
  var min_y, min_x float64 = math.MaxInt64, math.MaxInt64
  var min_point Point
  for _, point := range v.points {
    if point.y == min_y {
      if point.x > min_x {
        continue
      }
    }
    if point.y > min_y {
      continue
    }
      min_point = point
      min_y = point.y
  }
  stack := Stack{}
  // sort points by polar angle around min_point
  stack.Push(min_point)
  _points = remove(_points, min_point)
  debug("before sort")
  debug("min_point: %+v points: %+v\n", min_point, _points)
  quicksortByPolar(_points, min_point)
  debug("after sort")
  // push the first two points on the sorted list of points
  stack.Push(_points[0])
  stack.Push(_points[1])



  for i:=2 ; i < len(_points); i++ {
    for !IsLeftTurn(stack.NextToTop(), stack.Top(), _points[i]){
      stack.Pop()
    }
    stack.Push(_points[i])
  }

  return stack.All()
}

func IsLeftTurn(a, b, c Point) bool {
  // calculate (c-a) X (b-a)
  first := Point{(c.x - a.x), (c.y - a.y)}
  second := Point{(b.x - a.x), (b.y - a.y)}
  return (first.x * second.y - second.x*first.y) < 0
}
func Sum(x int, y int) int {
    return x + y
}

/*


  */
func tsp_helper(points []Point) ([]Point, float64) {
  if len(points) == 1 {
     return points, 0
  }
  min := math.MaxFloat64
  var min_path []Point
  for i, point := range points[1:] {
    debug("before append")
    target_points := append(points[1:i], points[i+1:]...)
    debug("after append")
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
  return math.Sqrt(math.Pow(math.Abs(float64(p.x - q.x)),2) + math.Pow(math.Abs(float64(p.x - q.x)), 2))
}
func main() {

	points_file := os.Args[1]

	points_corpus, err := ioutil.ReadFile(points_file)
	if err != nil {
		log.Fatal(err)
	}
  var points [][]float64
  err = json.Unmarshal(points_corpus, &points)
  fmt.Printf("%+v", points)
  vgraph := newGraph(points)
  tsp := vgraph.TspApprox()
  fmt.Printf("%+v", tsp)
}
