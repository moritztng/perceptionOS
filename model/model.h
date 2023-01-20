#pragma once

#include "onnxruntime_c_api.h"

typedef struct Model { 
    const OrtApi *g_ort;
    OrtSession *session;
    OrtMemoryInfo *memory_info;
    const char **input_names;
    const char **output_names;
    int64_t *input_shape;
    size_t input_shape_len;
    size_t model_input_len;
} Model;

void NewModel(const char *filepath, const char **input_names, const char **output_names, int64_t *input_shape, size_t input_shape_len, Model *model);

void Run(Model model, float *input, float **output);
