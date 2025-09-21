Quadchecker

This program checks ASCII rectangle patterns and identifies which of the known quad functions created them. The quad functions are called quadA, quadB, quadC, quadD, and quadE.

The program reads a shape from standard input, calculates its width and height, and then runs each of the quad generators with those same dimensions. It compares the result of each generator with the given input. If there is a match, the program outputs the name of the matching quad together with the dimensions. If no match is found, it outputs that the shape is not a quad function.

The purpose of this tool is to verify the correctness of quad functions and to make it easy to know exactly which quad implementation was used to produce a given rectangle.