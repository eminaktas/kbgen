[package]
name = "kubeconf"
version = "0.0.1"
description = "Test KCL Package"

[dependencies]
konfig = { git = "https://github.com/kcl-lang/konfig", commit = "c1ca802" }

[profile]
entries = ["main.k", "${konfig:KCL_MOD}/models/kube/render/render.k"]
