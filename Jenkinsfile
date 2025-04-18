pipeline {
    agent any

    environment {
        DOCKER_HUB_CREDENTIALS = credentials('docker-hub-creds')
        IMAGE_NAME = "jachant/yadro"
        APP_VERSION = "1.0.0"
        BRANCH_NAME = "${env.BRANCH_NAME ?: 'unknown'}".replaceAll('/', '-')
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
                script {
                    if (!fileExists('Dockerfile')) {
                        error("❌ Dockerfile не найден в репозитории.")
                    }
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Логинимся в Docker Hub перед сборкой (нужно для приватных базовых образов)
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-creds') {
                        dockerImage = docker.build(
                            "${IMAGE_NAME}:${BRANCH_NAME}-${APP_VERSION}",
                            "--build-arg VERSION=${APP_VERSION} ."
                        )
                    }
                }
            }
        }

        stage('Test Image') {
            steps {
                script {
                    dockerImage.inside {
                        sh 'echo "✅ Запуск тестов..."'
                        // Здесь можно добавить реальные тесты
                        sh 'curl -v http://localhost:8080/health || exit 1'
                    }
                }
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-creds') {
                        // Пуш основного тега
                        dockerImage.push()
                        
                        // Добавляем тег latest для main/master
                        if (BRANCH_NAME == 'main' || BRANCH_NAME == 'master') {
                            dockerImage.push('latest')
                            echo "🚀 Образ ${IMAGE_NAME}:latest обновлен"
                        }
                        
                        // Вывод информации об образе
                        sh "docker inspect ${IMAGE_NAME}:${BRANCH_NAME}-${APP_VERSION}"
                    }
                }
            }
        }

        stage('Cleanup') {
            steps {
                script {
                    // Удаляем локальный образ после пуша
                    sh "docker rmi ${IMAGE_NAME}:${BRANCH_NAME}-${APP_VERSION} || true"
                    if (BRANCH_NAME == 'main' || BRANCH_NAME == 'master') {
                        sh "docker rmi ${IMAGE_NAME}:latest || true"
                    }
                }
            }
        }
    }

    post {
        always {
            // Выходим из Docker Registry
            sh 'docker logout registry.hub.docker.com || true'
        }
        success {
            echo "✅ Сборка ${IMAGE_NAME}:${BRANCH_NAME}-${APP_VERSION} успешно завершена!"
            // slackSend(channel: '#ci-cd', message: "Успех: ${env.BUILD_URL}")
        }
        failure {
            echo "❌ Сборка завершилась с ошибкой"
            // slackSend(channel: '#ci-cd', message: "Провал: ${env.BUILD_URL}")
        }
    }
}