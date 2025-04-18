pipeline {
    agent any

    environment {
        DOCKER_HUB_CREDENTIALS = credentials('docker-hub-creds')
        IMAGE_NAME = "jachant/yadro"  // Полное имя репозитория Docker Hub
        BRANCH_NAME = "${env.BRANCH_NAME ?: 'unknown'}".replaceAll('/', '-')
        APP_VERSION = "1.0.0"
    }

   stage('Checkout') {
    steps {
        checkout scm
        script {
            if (!fileExists('YadroRest/Dockerfile')) {  // ← Указываем путь
                error("❌ Dockerfile не найден!")
            }
        }
    }


        stage('Build Docker Image') {
            steps {
                script {
                    // Сборка с полным путем Docker Hub
                    docker.build("${IMAGE_NAME}:${BRANCH_NAME}-${APP_VERSION}", '.')
                }
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-creds') {
                        // Пуш основного тега
                        docker.image("${IMAGE_NAME}:${BRANCH_NAME}-${APP_VERSION}").push()
                        
                        // Добавляем тег latest для main/master
                        if (BRANCH_NAME == 'main' || BRANCH_NAME == 'master') {
                            docker.image("${IMAGE_NAME}:${BRANCH_NAME}-${APP_VERSION}")
                                .tag("${IMAGE_NAME}:latest")
                            docker.image("${IMAGE_NAME}:latest").push()
                            echo "🚀 Образ ${IMAGE_NAME}:latest обновлен"
                        }
                    }
                }
            }
        }

        stage('Cleanup') {
            steps {
                script {
                    // Удаляем локальные образы после пуша
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