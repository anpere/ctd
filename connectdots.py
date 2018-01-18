import argparse
from svgpathtools import (
    svg2paths,
    wsvg,
)
import svgwrite

s_prefix = 'svg_'
connect_prefix = 'connect_'
if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Process some integers.')
    parser.add_argument('svg_file')
    
    args = parser.parse_args()
    scale = 10
    dwg = svgwrite.Drawing(connect_prefix+args.svg_file, size=(1700,1700))
    paths, attributes = svg2paths(args.svg_file)
    i = 0
    point_to_int = {}
    for path in paths:
       for line in path:
           if i == 23 or i == 0:
               print line
               print point_to_int
               print line.start not in point_to_int
               print line.end not in point_to_int
           if line.start not in point_to_int:
               point_to_int[line.start] = i
               i+=1
           if line.end not in point_to_int:
               point_to_int[line.end] = i
               i+=1
    for point, val in point_to_int.items():
        x = scale*point.real
        y = scale*point.imag
        dwg.add(dwg.text(str(val), insert=(x,y), fill='black'))
        dwg.add(dwg.circle(r=3,center=(x,y), fill='blue'))

    dwg.save()

    wsvg(paths, attributes=attributes, filename=s_prefix+args.svg_file)
    
    
