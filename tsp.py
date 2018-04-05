from copy import deepcopy
from random import shuffle

def tsp(points):
    """
    Returns the shortest cycle visiting each point exactly once
    """
    starting_point = points[0]
    point_to_tsp = tsp_helper(points[1:])
    min_length = 999999
    min_path = []
    for point, tsp in point_to_tsp.items():
        new_length = dist(starting_point, point) + path_length(tsp)
        if new_length < min_path:
            min_path = [starting_point] + tsp
            min_length = new_length
    return min_path

def tsp_helper(points):
    '''
    Returns a map from ponits to the shortest hamilton cycle visiting all those paths
    '''
    if len(points) == 1:
        return {points[0]: points}
    point_to_tsp = {}
    min_path = []
    min_length = 999999
    for point in points:
        others = deepcopy(points)
        before = len(others)
        others.remove(point)
        assert len(others) < before
        others_to_tsp = tsp_helper(others)
        for other_point, local_tsp in others_to_tsp.items():
            new_length = dist(point, other_point) + path_length(local_tsp)
            if new_length < min_length:
                min_length = new_length
                min_path = [point] + local_tsp
        point_to_tsp[point] = min_path
    if len(points) == 1:
        print(point_to_tsp)
    return point_to_tsp

def path_length(points):
    ''' Return the euclidean distance from traversing the points in order'''
    length = 0
    for i in range(len(points)-1):
        length += dist(points[i], points[i+1])
    return length
    return reduce(dist, points)

def dist(a, b):
    ''' returns the euclidean distance between two points '''
    return (abs(a.x -b.x)**2 + abs(a.y - b.y)**2)**.5

class Point(object):
    def __init__(self, x, y):
        self.x = x
        self.y = y
    def __eq__(self, other):
        return self.x == other.x and self.y == other.y
    def __str__(self):
        return "({}, {})".format(self.x, self.y)


if __name__ == '__main__':

    simple_points = [Point(0, 0), Point(0, 1), Point(0, 2)]
    p = tsp(simple_points)
    print([str(point) for point in p])
    assert path_length(p) == 2, "{} != 2".format(path_length(p))
    points = [
        Point(0,0), Point(1,2), Point(2,1), Point(3,1),
        Point(4,1), Point(5,1), Point(6,1), Point(7,1),
    ]
    points = [Point(0, i) for i in range(10)]
    l_p = path_length(points)
    assert Point(0,0) == Point(0,0)
    assert l_p == 9, "{} != 10".format(l_p)
    least_path = tsp(points)

    assert path_length(least_path) == 9, "{} != 9".format(l_p)

    counter_example = [
        Point(0,0), Point(1,2), Point(2,0), Point(3,0),
        Point(4,0), Point(5,0), Point(6,0), Point(7,0),
    ]
    inp = counter_example[1:]
    # shuffle(inp)
    # least_path = tsp([Point(0,0)]+inp)
    least_path = tsp(counter_example)
    least_path_length = path_length(least_path)
    print("LPL for counter_example: {}".format(least_path_length))

    for i in range(len(least_path)):
        assert least_path[i] == counter_example[i], "{} != {}, {}".format(least_path[i], counter_example[i], i)
    assert least_path == counter_example
