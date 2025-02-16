import models

appConfiguration: models.Server {
    image = "nginx"
    mainContainer = {
        name = "nginx"
        env: {
            "GET_HOSTS_FROM": {value = "dns"}
        }
        ports = [{containerPort = 80}]
    }
}
