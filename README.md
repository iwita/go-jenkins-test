# Sample Jenkins-Ready git repository

This repository includes a sample web application provided by Gophercises and Jon Calhoun, which displays a story and lets user decide the next steps of it.

Furthermore, there are also included:

### Jenkinsfile

A Jenkinsfile includes all the commands needed to declare and describe every step of your Jenkins pipeline

### Dockerfile

A Dockerfile includes all the needed commands in order to create a customized container using a base image. In this case, it is used by Jenkins in order to build an image using the executable produced by building the code.

### YAML configuration file (myweb.yaml)

This yaml file is the configuration file needed by Kubernetes in order to deploy a container as a Kubernetes understandable object

