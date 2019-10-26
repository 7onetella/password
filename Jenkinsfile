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
                input 'continue?'
                sh label: 'Test',       script: './execute.sh test'
            }
        }
        stage('Release') {
            steps {
                sh label: 'Releasing', script: 'echo releasing to production'
            }
        }
    }
}
