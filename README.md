# hello-kubic
Hello World POD for openSUSE Kubic

![Screenshot](images/Hello-Kubic-Screenshot.png "Screenshot")

This POD consisting of one container image can be deployed on a
Kubernetes cluster and accessed via a load balancer. It displays:
  * The kubernetes logo and string
  * "Hello Kubic!", which can be changed via the environment variable "MESSAGE"
  * The openSUSE Kubic logo and string
  * Informations about the Node and POD

The default port is 80 and can be changed in the service section of the yaml
file. 

To deploy: ``kubectl apply -f https://raw.githubusercontent.com/thkukuk/hello-kubic/master/yaml/hello-kubic.yaml``

The image is available from registry.opensuse.org: 
`registry.opensuse.org/kubic/hello-kubic:latest`
