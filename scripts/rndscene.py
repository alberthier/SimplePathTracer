#!/usr/bin/env python3

import random
import math

materials=["ground",
   "mat1",
   "mat2",
   "mat3",
   "mat4",
   "mat5",
   "mat6",
   "mat7",
   "metal1",
   "metal2",
   "metal3",
   "metal4",
   "metal5",
   "metal6",
   "glass",
   "diamond"]

def makeball(i, j):
    x = i + random.random() - 0.5
    y = 0.5
    z = j + random.random() - 0.5
    r = 0.5
    material = materials[random.randint(0, len(materials) - 1)]
    print('{{ "type": "sphere", "position": {{ "x": {:.1f}, "y": {:.1f}, "z": {:.1f} }}, "radius": {{ "value": {:.1f} }}, "material": "{}" }},'.format(x, y, z, r, material))

for i in range(-20, 21):
    for j in range(-20, 21):
        d = math.sqrt(i*i + j*j)
        if d > 6:
            if random.random() < 0.1:
                makeball(i, j)
