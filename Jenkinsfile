pipeline {
    environment {
        registry = "172.9.0.240:5000/iccs/jenkins-test/myweb"
        registryCredential = 'ca.crt'
        dockerImage = ""
    }

    // agent { docker { image 'golang'} }
    agent any

    stages {
    //   stage('Checkout Source') {
        //     steps {
	//	sh 'git config --global --unset http.proxy'
	//	sh 'git config --global --unset https.proxy'
        //      git 'https://https://github.com/iwita/go-jenkins-test.git'
       //    }

     //  }
        
        // stage ("Pre-build") {
        //     steps {
        //         sh 'go version'
        //         sh 'go test'

        //     }
        // }


  	stage("Fix the permission issue") {

            steps {
                sh "sudo chown tzenetoa:docker-evolve /var/run/docker.sock"
            }

        }

        stage('Build Image') {
            steps {
                script {
                    sh 'sudo su'
                    dockerImage = docker.build registry + ":$BUILD_NUMBER"
                }
            }
            
        }
        stage('Push Image') {
            steps {
                script {
		  sh 'sudo su'
                //  dockerImage.push()
		  sh "sudo docker push ${registry}"
                }
            }
        }
        stage('Deploy App') {
            steps {
                script {
                    step([$class: 'KubernetesEngineBuilder', location: env.LOCATION, manifestPattern: 'myweb.yaml'])
                }
            }
        }
    }
}
