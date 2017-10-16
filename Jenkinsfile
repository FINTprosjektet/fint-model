pipeline {
    agent none
    stages {
        stage('Prepare') {
            agent { label 'master' }
            steps {
                sh 'git log --oneline | nl -nln | perl -lne \'if (/^(\\d+).*Version (\\d+\\.\\d+\\.\\d+)/) { print "$2-$1"; exit; }\' > version.txt'
                stash includes: 'version.txt', name: 'version'
            }
        }
        stage('Build') {
            agent { label 'docker' }
            steps {
                unstash 'version'
                script {
                    VERSION=readFile('version.txt').trim()
                }
                sh "docker build --build-arg VERSION=${VERSION} ."
                archiveArtifacts 'build/**'
            }
        }
        stage('Publish') {
            agent { label 'docker' }
            when {
                branch 'master'
            }
            steps {
                unstash 'version'
                script {
                    VERSION=readFile('version.txt').trim()
                }
            }
        }
    }
}
