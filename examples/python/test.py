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

# Define Function Definition
readAll = lib.ReadAll
readAll.argtypes = [c_char_p, c_char_p, POINTER(c_char_p)]
readAll.restype = c_char_p

version = lib.Version
version.argtypes = []
version.restype = c_char_p

print "version:", version()

# Define input and output variables
# Output must be a ctypec_char_p
input_uri = "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/2.5_month.geojson"
input_alg = "none"
output_string_pointer = c_char_p()

print input_uri

err = readAll(input_uri, input_alg, byref(output_string_pointer))
if err != None:
    print("error: %s" % (str(err, encoding='utf-8')))
    sys.exit(1)

# Convert from ctype to python string
output_string = output_string_pointer.value

# Print output to stdout
print output_string
