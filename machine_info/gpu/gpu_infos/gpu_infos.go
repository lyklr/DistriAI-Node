package gpu

type GpuInfo struct {
	Name string
	Fp16 float32
	Fp32 float32
	Fp64 float32
}

func InitGpuInfos() []GpuInfo {
	var gpuInfos []GpuInfo
	gpuInfos = append(gpuInfos,
		GpuInfo{Name: "NVIDIA GeForce RTX 3090", Fp16: 35.58, Fp32: 35.58, Fp64: 0.556},
		GpuInfo{Name: "NVIDIA GeForce RTX 3090 Ti", Fp16: 40.00, Fp32: 40.00, Fp64: 0.625},
		GpuInfo{Name: "NVIDIA GeForce RTX 3050 6 GB", Fp16: 6.021, Fp32: 6.021, Fp64: 0.094},
		GpuInfo{Name: "NVIDIA GeForce RTX 4070 SUPER", Fp16: 35.48, Fp32: 35.48, Fp64: 0.554},
		GpuInfo{Name: "NVIDIA GeForce RTX 4070 Ti SUPER", Fp16: 44.10, Fp32: 44.10, Fp64: 0.689},
		GpuInfo{Name: "NVIDIA GeForce GTX 1080", Fp16: 0.138, Fp32: 8.87, Fp64: 0.277},
		GpuInfo{Name: "NVIDIA GeForce GTX 1080 11Gbps", Fp16: 0.138, Fp32: 8.87, Fp64: 0.277},
		GpuInfo{Name: "NVIDIA GeForce GTX 1080 Ti", Fp16: 0.177, Fp32: 11.34, Fp64: 0.354},
		GpuInfo{Name: "NVIDIA GeForce GTX 1080 Ti 10 GB", Fp16: 0.167, Fp32: 10.69, Fp64: 0.334},
		GpuInfo{Name: "NVIDIA GeForce GTX 1080 Ti 12 GB", Fp16: 0.167, Fp32: 10.69, Fp64: 0.334},
		GpuInfo{Name: "NVIDIA Tesla C1080", Fp16: 0.0, Fp32: 0.622, Fp64: 0.077},
		GpuInfo{Name: "NVIDIA GeForce RTX 4080 SUPER", Fp16: 51.3, Fp32: 51.3, Fp64: 0.801},
		GpuInfo{Name: "NVIDIA GeForce RTX 4090D", Fp16: 73.54, Fp32: 73.54, Fp64: 1.149},
		GpuInfo{Name: "NVIDIA GeForce RTX 4060 AD106", Fp16: 19.47, Fp32: 19.47, Fp64: 0.304},
		GpuInfo{Name: "NVIDIA GeForce RTX 4080 Ti", Fp16: 67.58, Fp32: 67.58, Fp64: 1.056},
		GpuInfo{Name: "NVIDIA GeForce RTX 4050", Fp16: 13.52, Fp32: 13.52, Fp64: 0.211},
		GpuInfo{Name: "NVIDIA RTX 4000 Ada Generation", Fp16: 26.73, Fp32: 26.73, Fp64: 0.417},
		GpuInfo{Name: "NVIDIA RTX 4500 Ada Generation", Fp16: 39.63, Fp32: 39.63, Fp64: 0.619},
		GpuInfo{Name: "NVIDIA RTX 5000 Ada Generation", Fp16: 65.28, Fp32: 65.28, Fp64: 1.02},
		GpuInfo{Name: "NVIDIA GeForce RTX 4060", Fp16: 15.11, Fp32: 15.11, Fp64: 0.236},
		GpuInfo{Name: "NVIDIA GeForce RTX 4060 Ti 16 GB", Fp16: 22.06, Fp32: 22.06, Fp64: 0.344},
		GpuInfo{Name: "NVIDIA GeForce RTX 4060 Ti 8 GB", Fp16: 22.06, Fp32: 22.06, Fp64: 0.344},
		GpuInfo{Name: "NVIDIA GeForce RTX 4070", Fp16: 29.15, Fp32: 29.15, Fp64: 0.455},
		GpuInfo{Name: "NVIDIA H100 CNX", Fp16: 215.4, Fp32: 53.84, Fp64: 26.92},
		GpuInfo{Name: "NVIDIA H100 PCIe 80 GB", Fp16: 204.9, Fp32: 51.22, Fp64: 25.61},
		GpuInfo{Name: "NVIDIA H100 PCIe 96 GB", Fp16: 248.3, Fp32: 62.08, Fp64: 31.04},
		GpuInfo{Name: "NVIDIA H100 SXM5 64 GB", Fp16: 267.6, Fp32: 66.91, Fp64: 33.45},
		GpuInfo{Name: "NVIDIA H100 SXM5 80 GB", Fp16: 267.6, Fp32: 66.91, Fp64: 33.45},
		GpuInfo{Name: "NVIDIA H100 SXM5 96 GB", Fp16: 248.3, Fp32: 62.08, Fp64: 31.04},
		GpuInfo{Name: "NVIDIA H800 PCIe 80 GB", Fp16: 204.9, Fp32: 51.22, Fp64: 25.61},
		GpuInfo{Name: "NVIDIA H800 SXM5", Fp16: 237.2, Fp32: 59.3, Fp64: 29.65},
		GpuInfo{Name: "NVIDIA L4", Fp16: 31.33, Fp32: 31.33, Fp64: 0.489},
		GpuInfo{Name: "NVIDIA RTX 4000 SFF Ada Generation", Fp16: 19.17, Fp32: 19.17, Fp64: 0.299},
		GpuInfo{Name: "NVIDIA GeForce RTX 4090 Ti", Fp16: 93.24, Fp32: 93.24, Fp64: 1.457},
		GpuInfo{Name: "NVIDIA TITAN Ada", Fp16: 92.9, Fp32: 92.9, Fp64: 1.452},
		GpuInfo{Name: "NVIDIA GeForce RTX 4070 Ti", Fp16: 40.09, Fp32: 40.09, Fp64: 0.626},
		GpuInfo{Name: "NVIDIA GeForce RTX 3050 8 GB GA107", Fp16: 9.09, Fp32: 9.09, Fp64: 0.142},
		GpuInfo{Name: "NVIDIA RTX 6000 Ada Generation", Fp16: 91.06, Fp32: 91.06, Fp64: 1.423},
		GpuInfo{Name: "NVIDIA A800 PCIe 40 GB", Fp16: 77.97, Fp32: 19.49, Fp64: 9.74},
		GpuInfo{Name: "NVIDIA A800 PCIe 80 GB", Fp16: 77.97, Fp32: 19.49, Fp64: 9.74},
		GpuInfo{Name: "NVIDIA GeForce RTX 3070 TiM", Fp16: 16.6, Fp32: 16.6, Fp64: 0.254},
		GpuInfo{Name: "NVIDIA GeForce RTX 3070 Ti 8 GB GA102", Fp16: 21.75, Fp32: 21.75, Fp64: 0.339},
		GpuInfo{Name: "NVIDIA GeForce RTX 3060 Ti GDDR6X", Fp16: 16.2, Fp32: 16.2, Fp64: 0.253},
		GpuInfo{Name: "NVIDIA L40", Fp16: 90.52, Fp32: 90.52, Fp64: 1.414},
		GpuInfo{Name: "NVIDIA L40 CNX", Fp16: 89.97, Fp32: 89.97, Fp64: 1.406},
		GpuInfo{Name: "NVIDIA L40G", Fp16: 89.97, Fp32: 89.97, Fp64: 1.406},
		GpuInfo{Name: "NVIDIA L40S", Fp16: 91.61, Fp32: 91.61, Fp64: 1.431},
		GpuInfo{Name: "NVIDIA GeForce RTX 3060 8 GB", Fp16: 12.74, Fp32: 12.74, Fp64: 0.199},
		GpuInfo{Name: "NVIDIA GeForce RTX 3060 8 GB GA104", Fp16: 12.74, Fp32: 12.74, Fp64: 0.199},
		GpuInfo{Name: "NVIDIA GeForce RTX 4080", Fp16: 48.74, Fp32: 48.74, Fp64: 0.761},
		GpuInfo{Name: "NVIDIA GeForce RTX 4080 12 GB", Fp16: 40.09, Fp32: 40.09, Fp64: 0.626},
		GpuInfo{Name: "NVIDIA GeForce RTX 4090", Fp16: 82.58, Fp32: 82.58, Fp64: 1.29},
		GpuInfo{Name: "NVIDIA GeForce RTX 4090 D", Fp16: 73.54, Fp32: 73.54, Fp64: 1.149},
		GpuInfo{Name: "NVIDIA Tesla T40 24 GB", Fp16: 28.75, Fp32: 14.38, Fp64: 0.449},
		GpuInfo{Name: "NVIDIA A800 SXM4 80 GB", Fp16: 77.97, Fp32: 19.49, Fp64: 9.74},
		GpuInfo{Name: "NVIDIA GeForce GTX 1630", Fp16: 3.65, Fp32: 1.82, Fp64: 0.057},
		GpuInfo{Name: "NVIDIA RTX A5500", Fp16: 34.1, Fp32: 34.1, Fp64: 1.066},
		GpuInfo{Name: "NVIDIA GeForce RTX 3060 Ti GA103", Fp16: 16.2, Fp32: 16.2, Fp64: 0.253},
		GpuInfo{Name: "NVIDIA GeForce RTX 3070 Ti 16 GB", Fp16: 21.75, Fp32: 21.75, Fp64: 0.339},
		GpuInfo{Name: "NVIDIA GeForce RTX 3050 4 GB", Fp16: 8.01, Fp32: 8.01, Fp64: 0.125},
		GpuInfo{Name: "NVIDIA GeForce RTX 3080 12 GB", Fp16: 30.64, Fp32: 30.64, Fp64: 0.478},
		GpuInfo{Name: "NVIDIA GeForce RTX 3050 8 GB", Fp16: 9.09, Fp32: 9.09, Fp64: 0.142},
		GpuInfo{Name: "NVIDIA GeForce RTX 3050 OEM", Fp16: 8.08, Fp32: 8.08, Fp64: 0.126},
		GpuInfo{Name: "NVIDIA GeForce RTX 3080 Ti 20 GB", Fp16: 34.1, Fp32: 34.1, Fp64: 0.532},
		GpuInfo{Name: "NVIDIA GeForce RTX 2060 12 GB", Fp16: 14.36, Fp32: 7.18, Fp64: 0.224},
		GpuInfo{Name: "NVIDIA RTX A2000 12 GB", Fp16: 7.98, Fp32: 7.98, Fp64: 0.124},
		GpuInfo{Name: "NVIDIA RTX A4500", Fp16: 23.65, Fp32: 23.65, Fp64: 0.739},
		GpuInfo{Name: "NVIDIA A2", Fp16: 4.53, Fp32: 4.53, Fp64: 0.07},
		GpuInfo{Name: "NVIDIA A2 PCIe", Fp16: 4.53, Fp32: 4.53, Fp64: 0.141},
		GpuInfo{Name: "NVIDIA RTX A500", Fp16: 7.25, Fp32: 7.25, Fp64: 0.113},)
	return gpuInfos
}
