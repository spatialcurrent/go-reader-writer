# =================================================================
#
# Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

from ctypes import *
import sys


# Load Shared Object
# grw.so must be in the LD_LIBRARY_PATH
# By default, LD_LIBRARY_PATH does not include current directory.
# You can add current directory with LD_LIBRARY_PATH=. python test.py
lib = cdll.LoadLibrary("grw.so")

# Define Function Definitions
schemes = lib.Schemes
schemes.argtypes = []
schemes.restype = c_char_p

# Define Function Definitions
algorithms = lib.Algorithms
algorithms.argtypes = []
algorithms.restype = c_char_p

# Define Function Definition
read_string = lib.ReadString
read_string.argtypes = [c_char_p, c_char_p, POINTER(c_char_p)]
read_string.restype = c_char_p

# Define Function Definition
write_string = lib.WriteString
write_string.argtypes = [c_char_p, c_char_p, c_int, c_char_p]
write_string.restype = c_char_p

# Define input and output variables
# Output must be a ctypec_char_p
input_uri = "https://raw.githubusercontent.com/spatialcurrent/go-reader-writer/master/testdata/doc.txt";
input_alg = "none"
output_string_pointer = c_char_p()

print("Algorithms:", algorithms())

print("Schemes:", schemes())

print("URI:", input_uri)

err = read_string(input_uri.encode('utf_8'), input_alg.encode('utf_8'), byref(output_string_pointer))
if err != None:
    print("error: %s" % (str(err, encoding='utf-8')))
    sys.exit(1)

# Convert from ctype to python string
output_string = output_string_pointer.value

# Print output to stdout (and keep stdout open)
err = write_string("stdout".encode('utf_8'), "none".encode('utf_8'), 0, output_string, 0)
if err != None:
    print("error: %s" % (str(err, encoding='utf-8')))
    sys.exit(1)

print("")
