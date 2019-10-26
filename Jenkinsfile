pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh label: 'Build',      script: './execute.sh build'
                sh label: 'Dev Deploy', script: './execute.sh deploy'
                input 'continue?'
                sh label: 'Test',       script: './execute.sh test'
            }
        }
        stage('Test') {
            steps {
                echo test place holder
            }
        }
        stage('Release') {
            steps {
                sh label: 'Releasing', script: 'echo releasing to production'
            }
        }
    }
}
