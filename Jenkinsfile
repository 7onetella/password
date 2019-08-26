pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                sh label: 'build', script: 'execute.sh build'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}
