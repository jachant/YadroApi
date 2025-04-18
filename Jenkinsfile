pipeline {
    agent any

    environment {
        // Учетные данные Docker Hub (ID из Jenkins Credentials)
        DOCKER_HUB_CREDENTIALS = credentials('docker-hub-creds')
        // Название образа в формате: <логин-docker>/<название-репозитория>
        IMAGE_NAME = "jachant/yadro"
    }

    stages {
        stage('Checkout') {
            steps {
                // Получение кода из репозитория GitHub
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Сборка Docker-образа с тегом = имени ветки
                    dockerImage = docker.build("${IMAGE_NAME}:${env.BRANCH_NAME}")
                }
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    // Авторизация в Docker Hub
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-creds') {
                        // Публикация образа
                        dockerImage.push()
                        // Опционально: добавить тег 'latest' для основной ветки
                        if (env.BRANCH_NAME == 'main') {
                            dockerImage.push('latest')
                        }
                    }
                }
            }
        }
    }

    post {
        success {
            // Уведомление о успешной сборке (опционально)
            echo "Сборка и публикация образа ${IMAGE_NAME}:${env.BRANCH_NAME} успешно завершены!"
        }
        failure {
            // Уведомление об ошибке (опционально)
            echo "Сборка завершилась с ошибкой."
        }
    }
}