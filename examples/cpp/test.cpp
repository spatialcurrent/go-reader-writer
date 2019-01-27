// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

#include <iostream>
#include <string>
#include <cstring>
#include "grw.h"

// readall is a example of a C++ function that can convert between formats using some std::string variables.
// In production, you would want to write the function definition to match the use case.
char* readall(std::string input_uri, std::string input_alg, char** output_string_c) {

  char *input_uri_c = new char[input_uri.length() + 1];
  std::strcpy(input_uri_c, input_uri.c_str());
  char *input_alg_c = new char[input_alg.length() + 1];
  std::strcpy(input_alg_c, input_alg.c_str());

  char *err = ReadAll(input_uri_c, input_alg_c, output_string_c);

  free(input_uri_c);
  free(input_alg_c);

  return err;

}

int main(int argc, char **argv) {

  // Since Go requires non-const values, we must define our parameters as variables
  // https://stackoverflow.com/questions/4044255/passing-a-string-literal-to-a-function-that-takes-a-stdstring
  std::string input_uri("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/2.5_month.geojson");
  std::string input_alg("none");
  char *output_char_ptr;

  // Write input to stderr
  std::cout << input_uri << std::endl;

  char *v = Version();
  std::cout << "version: " << std::string(v) << std::endl;

  char *err = readall(input_uri, input_alg, &output_char_ptr);
  if (err != NULL) {
    // Write output to stderr
    std::cerr << std::string(err) << std::endl;
    // Return exit code indicating error
    return 1;
  }
  std::string output_string = std::string(output_char_ptr);

  // Write output to stdout
  std::cout << output_string << std::endl;

  // Return exit code indicating success
  return 0;
}
