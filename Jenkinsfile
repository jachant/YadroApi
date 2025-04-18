pipeline {
    agent any

    environment {
        DOCKER_HUB_CREDENTIALS = credentials('docker-hub-creds')
        IMAGE_NAME = "jachant/yadro"
        BRANCH_NAME = "${env.BRANCH_NAME ?: 'unknown'}".replaceAll('/', '-')
        APP_VERSION = "1.0.0"
    }

    stages {
        stage('Check Docker') {
            steps {
                sh 'docker --version || (echo "Docker не установлен!" && exit 1)'
            }
        }

        stage('Checkout') {
            steps {
                checkout scm
                script {
                    if (!fileExists('YadroRest/Dockerfile')) {
                        error("❌ Dockerfile не найден!")
                    }
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    docker.build("${IMAGE_NAME}:${BRANCH_NAME}-${APP_VERSION}", '.')
                }
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-creds') {
                        docker.image("${IMAGE_NAME}:${BRANCH_NAME}-${APP_VERSION}").push()
                        if (BRANCH_NAME == 'main' || BRANCH_NAME == 'master') {
                            docker.image("${IMAGE_NAME}:${BRANCH_NAME}-${APP_VERSION}").tag("${IMAGE_NAME}:latest")
                            docker.image("${IMAGE_NAME}:latest").push()
                        }
                    }
                }
            }
        }
    }

    post {
        always {
            sh 'docker logout registry.hub.docker.com || true'
        }
        success {
            echo "✅ Образ ${IMAGE_NAME}:${BRANCH_NAME}-${APP_VERSION} успешно опубликован!"
        }
        failure {
            echo "❌ Ошибка при публикации образа"
        }
    }
}