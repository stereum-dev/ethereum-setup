pipeline {
    agent any
    stages {
        stage('Build Bundle') {
            steps {
                checkout scm
                dir('base-installer') {                    
                    sh 'sed -i s/%%CONTROLCENTER_RELEASE%%/${CONTROLCENTER_RELEASE}/ roles/stereum-controlcenter/defaults/main.yaml'
                    sh 'bundle-playbook -f ./playbook.yaml'
                    sh 'mv playbook.run base-installer-${RELEASE}.run'
                    sh 'rm -f /var/jenkins_home/publish/*'
                    sh 'cp base-installer-${RELEASE}.run /var/jenkins_home/publish'
                }                
            }
        }
        stage('Build Launcher') {
            steps {
                checkout scm                
                dir('launcher') {                    
                    sh 'sed -i s/%%RELEASE%%/${RELEASE}/ launcher.py'
                    sh 'pyinstaller --onefile launcher.py'
                    sh 'cp dist/launcher /var/jenkins_home/publish'
                }                
            }
        }
        stage('Push') {
            steps {
                // TODO: build windows binary and copy to share
                sh '/var/jenkins_home/release_publish.sh'
            }                
        }        
    }
    post { 
        always { 
            cleanWs()
        }
    }
}
