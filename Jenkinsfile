pipeline {
    environment {
        registry = "172.9.0.240:5000/iccs/jenkins-test/myweb"
        registryCredential = 'ca.crt'
        dockerImage = ""
    }

    // agent { docker { image 'golang'} }
    agent any

    stages {
        //stage('Checkout Source') {
       //    steps {
        //      git 'https://https://github.com/iwita/go-jenkins-test.git'
       //    }

        //}
        
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
                    sh 'sudo docker login -u evolve -p 3v0lv3r3g1st2y'
                    dockerImage = docker.build registry + ":$BUILD_NUMBER"
                }
            }
            
        }
        stage('Push Image') {
            steps {
                script {
                  sh 'sudo docker login -u evolve -p 3v0lv3r3g1st2y'
		  sh 'echo' + dockerImage
                 // dockerImage.push()
		  sh 'docker push '
                }
            }
        }
        stage('Deploy App') {
            steps {
                script {
                    kubernetesDeploy(configs: "myweb.yaml", kubeconfig: "mykubeconfig")
                }
            }
        }
    }
}
