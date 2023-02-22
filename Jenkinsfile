pipeline 
{
    agent any
    tools 
    {
        go 'Go'
        docker 'Docker'
    }
    stages 
    {
        stage('BuildProject')
        {
            steps
            {
                checkout scmGit(branches: [[name: '*/dev']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/RutujaRRohom/LibraryManagement-Go.git']])
            }
        }
    }
}       
