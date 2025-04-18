pipeline {
    agent any

    environment {
        DOCKER_HUB_CREDENTIALS = credentials('docker-hub-creds')
        IMAGE_NAME = 'jachant/yadro'
        BRANCH_NAME = "${env.BRANCH_NAME ?: 'unknown'}".replaceAll('/', '-')
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
                script {
                    if (!fileExists('Dockerfile')) {
                        error("❌ Dockerfile не найден!")
                    }
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    docker.build("${IMAGE_NAME}:${BRANCH_NAME}", '.')
                }
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-creds') {
                        docker.image("${IMAGE_NAME}:${BRANCH_NAME}").push()
                    }
                }
            }
        }
    }

    post {
        always {
            cleanWs() // Очистка рабочей директории
        }
        success {
            echo "✅ Образ ${IMAGE_NAME}:${BRANCH_NAME} успешно опубликован!"
        }
        failure {
            echo "❌ Сборка завершилась с ошибкой."
        }
    }
}