pipeline {
    environment {
        registry = "172.9.0.240:5000/iccs/jenkins-test/myweb"
        dockerImage = ""
    }

    // agent { docker { image 'golang'} }
    agent any

    stages {
        stage('Checkout Source') {
            steps {
                'https://92.43.249.202:3000/evolve/golang-test.git'
            }

        }
        
        // stage ("Pre-build") {
        //     steps {
        //         sh 'go version'
        //         sh 'go test'

        //     }
        // }

        stage('Build Image') {
             
            steps {
                script {
                    dockerImage = docker.build registry + ":$BUILD_NUMBER"
                }
            }
            
        }
        stage('Push Image') {
            steps {
                script {
                    docker.withRegistry( "" ) {
                        dockerImage.push()
                    }
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
