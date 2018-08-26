# Simple Toy Raytracer

This is a very basic raytracer. Its just an excuse to experiment with the Go programming language

It currently supports:
* Object type
    * sphere
* Materials
    * Dielectic
    * Metal
    * Glass
* Texturing
    * Plain color textures
    * Checker color textures
    * Composite textures
* Camera
    * Depth of field
    * Aperture
* Animation
    * Scalar animation
    * Coordinate animation
* Rendering
    * Multicore support with goroutines
    * Output format: PNG
* Custom JSON scene file format

Nice to have improvements:

* Support for other object types, especially triangles as it would allow to import existing meshes
* Support for image textures
* Add volumetric smoke
* Subsurface scattering
* Many many more...

## Example Render

![Render](https://raw.githubusercontent.com/alberthier/SimpleRayTracer/master/render/output.png)

## Example Animation

![Animation](https://raw.githubusercontent.com/alberthier/SimpleRayTracer/master/render/output.gif)
