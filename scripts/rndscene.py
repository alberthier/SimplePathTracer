#!/usr/bin/env python3

import random

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

def makeball():
    x = 10 - random.random() * 20.0
    y = 0.5
    z = 10 - random.random() * 20.0
    r = 0.5
    material = materials[random.randint(0, len(materials) - 1)]
    print('{{ "type": "sphere", "x": {:.1f}, "y": {:.1f}, "z": {:.1f}, "radius": {:.1f}, "material": "{}" }},'.format(x, y, z, r, material))


for i in range(100):
    makeball()
