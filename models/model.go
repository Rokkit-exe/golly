package models

type Model struct {
	Name          string `json:"name"`
	Model         string `json:"model"`
	Modified_at   string `json:"modified_at"`
	Size          int64  `json:"size"`
	Digest        string `json:"digest"`
	ParameterSize int64  `json:"parameter_size"`
	Quantization  string `json:"quantization"`
}
