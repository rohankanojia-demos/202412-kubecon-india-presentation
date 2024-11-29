# KubeCon India Presentation : Developing Kubernetes applications beyond go

## Table of Contents:
- [Introduction](#introduction)
- [Database Backup Operator Demo](#database-backup-operator-demo)
- [References](#references)

## Introduction

[Developing Kubernetes applications beyond go](https://kccncind2024.sched.com/event/1mVRN/developing-kubernetes-applications-beyond-go-rohan-kumar-red-hat-sun-tan-sciam)

In this presentation we will showcase various Kubernetes Libraries written in popular languages like:
- [Kubernetes Javascript Client](https://github.com/kubernetes-client/javascript)
- [Kubernetes Python Client](https://github.com/kubernetes-client/python)
- [Kubernetes Rust Client](https://github.com/kube-rs/kube)
- [Official Kubernetes Java Client](https://github.com/kubernetes-client/java)
- [Fabric8 Kubernetes Java Client](https://github.com/fabric8io/kubernetes-client)
- [Java Operator SDK](https://github.com/operator-framework/java-operator-sdk)

Code samples can be found in `kubernetes-client-examples/` folder

## Database Backup Operator Demo

For demo, we will be showcasing  a simple Kubernetes Operator written using [Quarkus Java Operator SDK](https://quarkus.io/extensions/io.quarkiverse.operatorsdk/quarkus-operator-sdk/) , you can find code for the operator in this [GitHub repository](https://github.com/rohankanojia-demos/database-backup-operator-java-operator-sdk).

## References
- [Kubernetes API reference](https://kubernetes.io/docs/reference/)
- [Java Operator SDK Documentation](https://javaoperatorsdk.io/docs/)
- [Kubernetes Operator Pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)
- [Minikube](https://minikube.sigs.k8s.io/docs/)
- [Quarkus Java Operator SDK](https://quarkus.io/extensions/io.quarkiverse.operatorsdk/quarkus-operator-sdk/)
- [Quarkus](https://quarkus.io)