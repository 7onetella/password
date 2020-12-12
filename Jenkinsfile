pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh 'echo hello world'
        sh 'docker build -t password .'
        sh 'docker tag password:latest docker-registry.7onetella.net:5000/7onetella/password:0.8.4'
        sh 'docker push docker-registry.7onetella.net:5000/7onetella/password:0.8.4'
      }
    }

  }
  environment {
    stage = 'dev'
  }
}