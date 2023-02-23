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

                sh "docker build -t library-management-ecr ." 

            }

        }

    }

    stage("Push Docker Image"){

        steps {

         withCredentials([

  [

    $class: 'AmazonWebServicesCredentialsBinding',

    accessKeyVariable: 'AWS_ACCESS_KEY_ID',

    secretKeyVariable: 'AWS_SECRET_ACCESS_KEY',

    credentialsId: 'AWS_Credential_Id'

  ]

]){

    sh '''

    aws ecr get-login-password --region ap-southeast-1 | docker login --username AWS --password-stdin 089705641992.dkr.ecr.ap-southeast-1.amazonaws.com/library-management-ecr

    docker tag library-management-ecr:latest 089705641992.dkr.ecr.ap-southeast-1.amazonaws.com/library-management-ecr:latest

    docker push 089705641992.dkr.ecr.ap-southeast-1.amazonaws.com/library-management-ecr:latest

    '''

}

    }

    }

    

    

    

  }

}
pipeline{

    agent {node {label 'SSH'} }

    stages{

        stage('Pull Docker Image'){

            steps{

                sh "docker pull 089705641992.dkr.ecr.ap-southeast-1.amazonaws.com/library-management-ecr:latest"

                sh "docker stop deploy || true"

                sh 'docker run --rm -p 3000:3000 --name deploy -v /home/ubuntu/project:/secrets -d 089705641992.dkr.ecr.ap-southeast-1.amazonaws.com/library-management-ecr:latest'
                sh 'docker ps'

            }

        }

    }

    post {




        always {




            mail bcc: '', body: " Hi Team \n I have forwarded the build status of $JOB_NAME  \n Build : $BUILD_NUMBER  ${currentBuild.currentResult}.\n \n \n Check the console output at ${env.BUILD_URL} to view results\n Thanks and Regards ", cc: 'rohit.sinha@joshsoftware.com', from: '', replyTo: '', subject: 'Test Email From Jenkins Job', to: 'dubeyakshat88@gmail.com'




        }




    }

}
