import models
import konfig.models.kube.frontend

_appConfiguration: models.Server {
    mainContainer = {
        name = "nginx"
        image = "nginx"
        env: {
            "GET_HOSTS_FROM": {value = "dns"}
        }
        ports = [{containerPort = 80}]
    }
}

appConfiguration: frontend.Server {
    image = _appConfiguration.mainContainer.image
    mainContainer = {
        name = _appConfiguration.mainContainer.name
        env = _appConfiguration.mainContainer.env
        ports = _appConfiguration.mainContainer.ports
    }
}
