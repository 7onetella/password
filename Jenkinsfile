pipeline {
  agent {
    dockerfile {
      filename 'Dockerfile'
    }

  }
  stages {
    stage('build') {
      steps {
        sh 'echo hello world'
      }
    }

  }
  environment {
    stage = 'dev'
  }
}