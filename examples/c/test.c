// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "reader.h"

int
main(int argc, char **argv) {
    char *err;

    char *input_uri = "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/2.5_month.geojson";
    char *input_alg = "none";
    char *output_string;

    printf("%s\n", input_uri);

    char *version = Version();
    printf("version: %s\n", version);

    err = ReadAll(input_uri, input_alg, &output_string);

    if (err != NULL) {
        fprintf(stderr, "error: %s\n", err);
        free(err);
        return 1;
    }

    printf("%s\n", output_string);

    return 0;
}
