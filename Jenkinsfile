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
                script {
                    // Обновляем статус GitHub
                    githubNotify(context: "Jenkins/Checkout", 
                                status: "PENDING", 
                                description: "Checkout in progress")
                }
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
                    githubNotify(context: "Jenkins/Build", 
                                status: "PENDING", 
                                description: "Building Docker image")
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-creds') {
                        dockerImage = docker.build(
                            "${IMAGE_NAME}:${BRANCH_NAME}-${APP_VERSION}",
                            "--build-arg VERSION=${APP_VERSION} ."
                        )
                    }
                    githubNotify(context: "Jenkins/Build", 
                                status: "SUCCESS", 
                                description: "Image built successfully")
                }
            }
        }

        stage('Test Image') {
            steps {
                script {
                    githubNotify(context: "Jenkins/Test", 
                                status: "PENDING", 
                                description: "Running tests")
                    dockerImage.inside {
                        sh 'echo "✅ Запуск тестов..."'
                        // Пример теста
                        sh 'curl -v http://localhost:8080/health || exit 1'
                    }
                    githubNotify(context: "Jenkins/Test", 
                                status: "SUCCESS", 
                                description: "Tests passed")
                }
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    githubNotify(context: "Jenkins/Push", 
                                status: "PENDING", 
                                description: "Pushing to Docker Hub")
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-creds') {
                        dockerImage.push()
                        if (BRANCH_NAME == 'main' || BRANCH_NAME == 'master') {
                            dockerImage.push('latest')
                        }
                    }
                    githubNotify(context: "Jenkins/Push", 
                                status: "SUCCESS", 
                                description: "Image pushed successfully")
                }
            }
        }

        stage('Cleanup') {
            steps {
                script {
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
            githubNotify(context: "Jenkins/Overall", 
                        status: "SUCCESS", 
                        description: "Build completed")
            echo "✅ Сборка ${IMAGE_NAME}:${BRANCH_NAME}-${APP_VERSION} успешна!"
        }
        failure {
            githubNotify(context: "Jenkins/Overall", 
                        status: "FAILURE", 
                        description: "Build failed")
            echo "❌ Сборка завершилась с ошибкой"
        }
    }
}