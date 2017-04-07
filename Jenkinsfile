pipeline {
  agent any
  stages {
    stage('Compile') {
      steps {
        sh 'echo "cool"'
      }
    }
    stage('Tests') {
      steps {
        sh 'exit 1'
      }
    }
  }
}