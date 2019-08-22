// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "grw.h"

int
main(int argc, char **argv) {
    char *err;

    printf("Algorithms: %s\n", Algorithms());

    printf("Schemes: %s\n", Schemes());

    char *input_uri = "https://raw.githubusercontent.com/spatialcurrent/go-reader-writer/master/test/doc.txt";
    char *input_alg = "none";
    char *output_string;

    printf("%s\n", input_uri);

    err = ReadString(input_uri, input_alg, &output_string);

    if (err != NULL) {
        fprintf(stderr, "error: %s\n", err);
        free(err);
        return 1;
    }

    printf("%s\n", output_string);

    return 0;
}
