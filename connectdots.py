import argparse
import json
import logging
from svgpathtools import (
    svg2paths,
    wsvg,
)
import svgwrite


s_prefix = 'svg_'
connect_prefix = 'connect_'
json_prefix = 'json_'
if __name__ == '__main__':
    logging.basicConfig(format='%(levelname)s:%(message)s', level=logging.DEBUG)
    parser = argparse.ArgumentParser(description='Process some integers.')
    parser.add_argument('svg_file')

    args = parser.parse_args()
    scale = 3
    dwg = svgwrite.Drawing(connect_prefix+args.svg_file, size=(1700,1700))
    logging.info("getting paths...")
    paths, attributes = svg2paths(args.svg_file)
    logging.info("done!")
    logging.info("Number of paths: {}".format(len(paths)))
    i = 0
    point_to_int = {}
    for path in paths:
       for line in path:
           for point in [line.start, line.end]:
               x = int(point.real)
               y = int(point.imag)
               if (x,y) not in point_to_int:
                   point_to_int[(x,y)] = i
                   i+=1
    for (x,y), val in point_to_int.items():
        x *= scale
        y *= scale
        dwg.add(dwg.text(str(val), insert=(x+3,y+3), fill='black'))
        dwg.add(dwg.circle(r=3,center=(x,y), fill='blue'))

    dwg.save()
    # Try to actually draw it
    inverted = {v: k for k,v in point_to_int.items()}
    logging.debug(sorted(inverted))
    drew = svgwrite.Drawing('post_'+args.svg_file, size=(1700,1700))
    stroke=svgwrite.rgb(10, 10, 16, '%')
    for j in sorted(inverted):
       logging.debug(inverted[j])
       try:
           drew.add(drew.line(inverted[j-1], inverted[j], stroke=stroke))
       except KeyError:
           drew.add(drew.line(inverted[j], inverted[i-1], stroke=stroke))
    drew.save()

    with open(json_prefix+args.svg_file, 'w') as outfile:
        json.dump(point_to_int.keys(), outfile)
