pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh 'echo hello world'
        sh 'docker build -t password .'
      }
    }

  }
  environment {
    stage = 'dev'
  }
}