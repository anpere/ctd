package main

import (
  "testing"
  "reflect"
)
// Convex Hull partitions
// input:
// vgraph with a
//  single point_to_int
//  two identical points
//  two different points
// output:
//  len(output) == len(vgraph) convex hull is equal to the graph
//  len(output) <= len(vgraph) convex hull is a subset of the graph
//
// func TestConvexHullSinglePoint(t *testing.T) {
//   points := []Point{Point{0,0}}
//   vgraph := Vgraph{points}
//   test_convex_hull := vgraph.ConvexHull()
//   if len(test_convex_hull) != 1 {
//     t.Errorf("ConvexHull length was incorrect, got: %d, want: %d.", len(test_convex_hull), 1)
//   }
//   if !reflect.DeepEqual(test_convex_hull, points){
//     t.Errorf("ConvexHull was incorrect, got :%d, want: %d.", test_convex_hull, points)
//   }
// }
//
// func TestConvexHullTwoEqualPoints(t *testing.T) {
//   points := []Point{Point{0, 0}, Point{0, 0}}
//   vgraph := Vgraph{points}
//   test_convex_hull := vgraph.ConvexHull()
//   correct := []Point{Point{0, 0}}
//   if !reflect.DeepEqual(test_convex_hull, correct){
//     t.Errorf("ConvexHull was incorrect, got :%d, want: %d.", test_convex_hull, correct)
//   }
// }
//
// func TestConvexHullTwoUnequalPoints(t *testing.T) {
//   points := []Point{Point{0,0}, Point{1, 0}}
//   vgraph := Vgraph{points}
//   test_convex_hull := vgraph.ConvexHull()
//   correct := []Point{Point{0,0}, Point{1, 0}}
//   if !reflect.DeepEqual(test_convex_hull, correct) {
//     t.Errorf("ConvexHull was incorrect, got: %d, want %d", test_convex_hull, correct)
//   }
// }
// Convex hull of
// * * * *
// * * * *
// * * *  *

func TestConvexHull(t *testing.T) {
  tables := []struct {
    points []Point
    convexhull []Point
  } {{
    //   []Point{Point{0,0}, Point{1,1}, Point{2,2}}, // test with more than 2 points
    //   []Point{Point{0,0}, Point{1,1}, Point{2,2}},
    // }, {
      []Point{Point{0,0}, Point{1,.5}, Point{.5,.2}, Point{1,1}, Point{2,0}},
      []Point{Point{0,0}, Point{1,1}, Point{2,0}},
    }, {
      []Point{Point{1,1}, Point{0,0}, Point{2,2}},
      []Point{Point{0,0}, Point{1,1}, Point{2,2}},
    }, {
      []Point{Point{0,0}, Point{0,2}, Point{2,2}, Point{2,0}, Point{1,1}},
      []Point{Point{0,0}, Point{0,2}, Point{2,2}, Point{2,0}},
    }, {
      []Point{Point{1,2}, Point{3,1}, Point{4,2}, Point{3,3}, Point{5,4}, Point{4,4}, Point{2,4}, Point{4,5}},
      []Point{Point{1,2}, Point{3,1}, Point{4,2}, Point{5,4}, Point{4,5}, Point{2,4}},
    }, {
      []Point{Point{1,1}, Point{2,2}, Point{3,3}, Point{3,1}}, // A point in Q is colinear to two points in CH
      []Point{Point{1,1}, Point{3,1}, Point{3,3}},
    },
  }
  for _, test := range tables {
    vgraph := Vgraph{test.points}
    test_convex_hull := vgraph.ConvexHull()
    if !reflect.DeepEqual(test_convex_hull, test.convexhull) {
      t.Errorf("ConvexHull was incorrect, got: %d, want %d", test_convex_hull, test.convexhull)
    }
  }
}

func TestIsLeftTurn(t *testing.T) {
  
}
func TestSum(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}
