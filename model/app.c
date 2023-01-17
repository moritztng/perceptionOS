#include <assert.h>
#include <stdio.h>

#include "onnxruntime_c_api.h"

#define ORT_ABORT_ON_ERROR(expr)                                   \
    do                                                             \
    {                                                              \
        OrtStatus *onnx_status = (expr);                           \
        if (onnx_status != NULL)                                   \
        {                                                          \
            const char *msg = g_ort->GetErrorMessage(onnx_status); \
            fprintf(stderr, "%s\n", msg);                          \
            g_ort->ReleaseStatus(onnx_status);                     \
            abort();                                               \
        }                                                          \
    } while (0);

typedef struct Model { 
    const OrtApi *g_ort;
    const OrtSession *session;
    const OrtMemoryInfo *memory_info;
    const char *input_names[];
    const char *output_names[];
    const int64_t input_shape[];
    const size_t input_shape_len;
    const size_t model_input_len;
} Model;

void NewModel(const char *filepath, const int64_t input_shape[], chModel *model) {
    model->g_ort = OrtGetApiBase()->GetApi(ORT_API_VERSION);
    OrtEnv *env;
    ORT_ABORT_ON_ERROR(model->g_ort->CreateEnv(ORT_LOGGING_LEVEL_WARNING, "test", &env));
    OrtSessionOptions *session_options;
    ORT_ABORT_ON_ERROR(model->g_ort->CreateSessionOptions(&session_options));
    ORT_ABORT_ON_ERROR(model->g_ort->CreateSession(env, filepath, session_options, &model.session));
    ORT_ABORT_ON_ERROR(model->g_ort->CreateCpuMemoryInfo(OrtArenaAllocator, OrtMemTypeDefault, &model.memory_info));
    model->input_shape = input_shape;
    model->input_shape_len = sizeof(input_shape) / sizeof(input_shape[0]);
    const size_t input_len = 1;
    for (int dim = 0; dim < model->input_shape_len; dim++) {
      input_len *= dim;
    }
    model->model_input_len = input_len;
}

void Run(Model *model, float *input, float *output) {
    OrtValue *input_tensor = NULL;
    ORT_ABORT_ON_ERROR(g_ort->CreateTensorWithDataAsOrtValue(model->memory_info, input, model->model_input_len, model->input_shape,
                                          model->input_shape_len, ONNX_TENSOR_ELEMENT_DATA_TYPE_FLOAT,
                                          &input_tensor));
    assert(input_tensor != NULL);
    int is_tensor;
    ORT_ABORT_ON_ERROR(g_ort->IsTensor(input_tensor, &is_tensor));
    assert(is_tensor);
    OrtValue *output_tensor = NULL;
    ORT_ABORT_ON_ERROR(g_ort->Run(model->session, NULL, model->input_names, (const OrtValue* const*)&input_tensor, 1, model->output_names, 1, &output_tensor));
    assert(output_tensor != NULL);
    ORT_ABORT_ON_ERROR(g_ort->GetTensorMutableData(output_tensor, (void**)&output));
}
