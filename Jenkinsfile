pipeline {
    environment {
        registry = "172.9.0.240:5000/iccs/jenkins-test/myweb"
        registryCredential = 'ca.crt'
        dockerImage = ""
    }

    agent any

    stages {
  	stage("Fix the permission issue") {
	   // In order to be able to use 'docker commands, root user
	   // needs to get added to the docker group
            steps {
                sh "sudo chown root:docker-evolve /var/run/docker.sock"
            }

        }
	
        stage('Build Image') {
            steps {
                script {
		   // we need to be connected as root (as root belongs in the docker group)
                    sh 'sudo su'
		    // build the image using the desired name, provided as a variable
                    dockerImage = docker.build registry + ":$BUILD_NUMBER"
                }
            }
            
        }
        stage('Push Image') {
            steps {
                script {
		  // push the built image in the private docker registry
		  sh "sudo docker push ${registry}:${BUILD_NUMBER}"
                }
            }
        }
        stage('Deploy App') {
            steps {
                script {
		  // we change the version of our deployment in the yaml file,
		  // and deploy the application in the Kubernetes Cluster
		  sh "sed -i s~${registry}.*~${registry}:${BUILD_NUMBER}~g myweb.yaml"
		  kubernetesDeploy(configs: "myweb.yaml", kubeconfigId: "mykubeconfig")
                }
            }
        }
    }
}
