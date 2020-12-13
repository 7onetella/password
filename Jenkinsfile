pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh './build.sh'
      }
    }
  }
  environment {
    stage = 'dev'
  }
}