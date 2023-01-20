#include <assert.h>
#include <stdio.h>
#include "model.h"

#include "onnxruntime_c_api.h"

#define ORT_ABORT_ON_ERROR(expr)                                         \
    do                                                                   \
    {                                                                    \
        OrtStatus *onnx_status = (expr);                                 \
        if (onnx_status != NULL)                                         \
        {                                                                \
            const char *msg = model.g_ort->GetErrorMessage(onnx_status); \
            fprintf(stderr, "%s\n", msg);                                \
            model.g_ort->ReleaseStatus(onnx_status);                     \
            abort();                                                     \
        }                                                                \
    } while (0);

void NewModel(const char *filepath, const char **input_names, const char **output_names, int64_t *input_shape, size_t input_shape_len, Model *output_model)
{
    Model model;
    model.g_ort = OrtGetApiBase()->GetApi(ORT_API_VERSION);
    OrtEnv *env;
    ORT_ABORT_ON_ERROR(model.g_ort->CreateEnv(ORT_LOGGING_LEVEL_WARNING, "test", &env));
    assert(env != NULL);
    OrtSessionOptions *session_options;
    ORT_ABORT_ON_ERROR(model.g_ort->CreateSessionOptions(&session_options));
    ORT_ABORT_ON_ERROR(model.g_ort->CreateSession(env, filepath, session_options, &model.session));
    ORT_ABORT_ON_ERROR(model.g_ort->CreateCpuMemoryInfo(OrtArenaAllocator, OrtMemTypeDefault, &model.memory_info));
    model.input_names = input_names;
    model.output_names = output_names;
    model.input_shape = input_shape;
    model.input_shape_len = input_shape_len;
    size_t input_len = 1;
    for (int dim = 0; dim < model.input_shape_len; dim++)
    {
        input_len *= input_shape[dim];
    }
    model.model_input_len = input_len * sizeof(float);
    *output_model = model;
}

void Run(Model model, float *input, float **output)
{
    OrtValue *input_tensor = NULL;
    ORT_ABORT_ON_ERROR(model.g_ort->CreateTensorWithDataAsOrtValue(model.memory_info, input, model.model_input_len, model.input_shape,
                                                                    model.input_shape_len, ONNX_TENSOR_ELEMENT_DATA_TYPE_FLOAT,
                                                                    &input_tensor));
    assert(input_tensor != NULL);
    int is_tensor;
    ORT_ABORT_ON_ERROR(model.g_ort->IsTensor(input_tensor, &is_tensor));
    assert(is_tensor);
    OrtValue *output_tensor = NULL;
    ORT_ABORT_ON_ERROR(model.g_ort->Run(model.session, NULL, model.input_names, (const OrtValue *const *)&input_tensor, 1, model.output_names, 1, &output_tensor));
    assert(output_tensor != NULL);
    ORT_ABORT_ON_ERROR(model.g_ort->GetTensorMutableData(output_tensor, (void **)output));
}

/*
int main() {
  Model model;
  NewModel("yolov4.onnx", (const char *[]){"input_1:0"}, (const char *[]){"Identity:0"}, (int64_t[]){1, 416, 416, 3}, 4, &model);
  size_t input_len = 1 * 416 * 416 * 3 * sizeof(float);
  float *input = (float *)malloc(input_len);
  memset(input, 0, input_len);
  float *output;
  Run(&model, input, &output);
  printf("%.6f", *output);
}*/
