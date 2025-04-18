pipeline {
    agent any

    environment {
        DOCKER_HUB_CREDENTIALS = credentials('docker-hub-creds')
        IMAGE_NAME = "jachant/yadro"
        APP_VERSION = "1.0.0" // Опционально: параметризация версии
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
                script {
                    // Проверка наличия Dockerfile
                    if (!fileExists('Dockerfile')) {
                        error("❌ Dockerfile не найден в репозитории.")
                    }
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Сборка с кэшированием и версией
                    dockerImage = docker.build(
                        "${IMAGE_NAME}:${env.BRANCH_NAME}-${APP_VERSION}",
                        "--build-arg VERSION=${APP_VERSION} ."
                    )
                }
            }
        }

        stage('Test Image') {
            steps {
                script {
                    // Пример запуска тестов
                    docker.image("${IMAGE_NAME}:${env.BRANCH_NAME}-${APP_VERSION}").inside {
                        sh 'echo "✅ Запуск тестов..."'
                    }
                }
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-creds') {
                        dockerImage.push()
                        // Тег 'latest' только для main
                        if (env.BRANCH_NAME == 'main') {
                            dockerImage.push('latest')
                            echo "🚀 Образ ${IMAGE_NAME}:latest обновлен."
                        }
                    }
                }
            }
        }
    }

    post {
        success {
            echo "✅ Сборка ${IMAGE_NAME}:${env.BRANCH_NAME}-${APP_VERSION} успешна!"
            // slackSend channel: '#ci-cd', message: "Успех: ${env.BUILD_URL}"
        }
        failure {
            echo "❌ Сборка завершилась с ошибкой."
            // slackSend channel: '#ci-cd', message: "Провал: ${env.BUILD_URL}"
        }
    }
}