pipeline {
    // install golang 1.18 on Jenkins node
    agent any
      environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0
        GOPROXY="https://mirrors.aliyun.com/goproxy/"
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage("Build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                // 执行 Makefile 中的命令
                sh 'make bl_linux'
                // 将 build 后的文件保存
                archiveArtifacts artifacts: 'bin/*', fingerprint: true
                sh 'ps -ef|grep "/opt/simple-web/simple-web-linux"|grep -v "grep"|awk "{print $2}"|xargs kill -9 '
                sh 'cp bin/simple-web-linux /opt/simple-web/simple-web-linux'
                sh '/opt/simple-web/simple-web-linux'
            }
        }
    }
}
