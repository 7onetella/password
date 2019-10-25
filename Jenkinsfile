pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh label: 'Build',      script: './execute.sh build'
                sh label: 'Dev Deploy', script: './execute.sh deploy'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Release') {
            steps {
                // sh label: 'Releasing', script: './execute.sh release'
            }
        }
    }
}
