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

// read_string is a example of a C++ function that can convert between formats using some std::string variables.
// In production, you would want to write the function definition to match the use case.
char* read_string(std::string input_uri, std::string input_alg, char** output_string_c) {

  char *input_uri_c = new char[input_uri.length() + 1];
  std::strcpy(input_uri_c, input_uri.c_str());
  char *input_alg_c = new char[input_alg.length() + 1];
  std::strcpy(input_alg_c, input_alg.c_str());

  char *err = ReadString(input_uri_c, input_alg_c, output_string_c);

  free(input_uri_c);
  free(input_alg_c);

  return err;

}

// write_string is a example of a C++ function that can convert between formats using some std::string variables.
// In production, you would want to write the function definition to match the use case.
char* write_string(std::string input_uri, std::string input_alg, std::string input_string) {

  char *input_uri_c = new char[input_uri.length() + 1];
  std::strcpy(input_uri_c, input_uri.c_str());

  char *input_alg_c = new char[input_alg.length() + 1];
  std::strcpy(input_alg_c, input_alg.c_str());

  char *input_string_c = new char[input_string.length() + 1];
  std::strcpy(input_string_c, input_string.c_str());

  char *err = WriteString(input_uri_c, input_alg_c, 1, input_string_c);

  free(input_uri_c);
  free(input_alg_c);
  free(input_string_c);

  return err;

}

int main(int argc, char **argv) {

  std::cout << "Algorithms: " << Algorithms() << std::endl;
  std::cout << "Schemes: " << Schemes() << std::endl;

  // Since Go requires non-const values, we must define our parameters as variables
  // https://stackoverflow.com/questions/4044255/passing-a-string-literal-to-a-function-that-takes-a-stdstring
  std::string input_uri("https://raw.githubusercontent.com/spatialcurrent/go-reader-writer/master/testdata/doc.txt.bz2");
  std::string input_alg("bzip2");
  char *output_char_ptr;

  // Write input to stderr
  std::cout << "Uri: " << input_uri << std::endl;

  char *errReadString = read_string(input_uri, input_alg, &output_char_ptr);
  if (errReadString != NULL) {
    // Write output to stderr
    std::cerr << std::string(errReadString) << std::endl;
    // Return exit code indicating error
    return 1;
  }
  std::string output_uri("stdout");
  std::string output_alg("none");
  std::string output_string = std::string(output_char_ptr);

  char *errWriteString = write_string(output_uri, output_alg, output_string);
  if (errWriteString != NULL) {
    // Write output to stderr
    std::cerr << std::string(errWriteString) << std::endl;
    // Return exit code indicating error
    return 1;
  }
  std::cout << std::endl; // just add an extra new line

  // Return exit code indicating success
  return 0;
}
