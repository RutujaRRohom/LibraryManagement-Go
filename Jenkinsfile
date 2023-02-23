pipeline {
  agent any
  tools {
      go 'Go'
      dockerTool 'Docker'
  }
  stages {
    stage('Fetch code from Github') {
      steps {
          checkout([$class: 'GitSCM', branches: [[name: '*/deployment']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[url: 'https://github.com/RutujaRRohom/LibraryManagement-Go.git/']]])
      }
    }
    stage("Build Docker Image"){
        steps{
            script{
//                 sh 'sudo chmod 666 /var/run/docker.sock -S'
                sh "docker build -t library-management-ecr ." 
            }
        }
    }
    stage("Push Docker Image"){
        steps {
         withCredentials([[
    $class: 'AmazonWebServicesCredentialsBinding',
    accessKeyVariable: 'AWS_ACCESS_KEY_ID',
    secretKeyVariable: 'AWS_SECRET_ACCESS_KEY',
    credentialsId: 'AWS_Credential_Id']]){
    sh '''
    aws ecr get-login-password --region ap-southeast-1 | docker login --username AWS --password-stdin 089705641992.dkr.ecr.ap-southeast-1.amazonaws.com/library-management-ecr
    docker tag library-management-ecr:latest 089705641992.dkr.ecr.ap-southeast-1.amazonaws.com/library-management-ecr:latest
    docker push 089705641992.dkr.ecr.ap-southeast-1.amazonaws.com/library-management-ecr:latest'''
    }
    }
    }

    

    

    

  }

}
