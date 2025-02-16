import konfig.models.kube.frontend.service
import konfig.models.kube.frontend.container
import konfig.models.kube.frontend.volume

schema Server:
    """Server is abstaction of Deployment and StatefulSet.

    Attributes
    ----------
    name: str, default is Undefined, optional.
        The name of the workload and service. 
        If not defined, a generated name ("{__META_APP_NAME}-{__META_ENV_TYPE_NAME}") will be used.
        The value of metadata.__META_APP_NAME will be extracted from the value of the "name" defined through option("app"),
        and the value of __META_ENV_TYPE_NAME will be extracted from the value of the "name" defined through option("env").
    namespace: str, default is Undefined, optional.
        The namespace of the workload and service.
        If not defined, a generated name ("{__META_APP_NAMESPACE}-{__META_ENV_TYPE_NAME}") will be used.
        The value of metadata.__META_APP_NAMESPACE will be extracted from the value of the "name" defined through option("appns"),
        and the value of __META_ENV_TYPE_NAME will be extracted from the value of the "name" defined through option("env").
    workloadType: "Deployment" | "StatefulSet", default is "Deployment", required.
        Application workload type, default to 'Deployment'
    replicas: int, default is 1, required.
        Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
    image: str, default is Undefined, required.
        Container image name.
        More info: https://kubernetes.io/docs/concepts/containers/images
    imagePullPolicy : str, default is Undefined, optional
        A Container-level attribute.
        The policy for image pulling. The value of
        this field must be one of "Always", "IfNotPresent" or
        "Never". Defaults to "IfNotPresent".
    mainContainer: container.Main, default is Undefined, required.
        MainContainer describes the main container configuration that is expected to be run on the host.
    labels: {str:str}, default is Undefined, optional.
        Labels is a map of string keys and values that can be used to organize and categorize (scope and select) objects.
        More info: http://kubernetes.io/docs/user-guide/labels
    annotations: {str:str}, default is Undefined, optional
        Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
        More info: http://kubernetes.io/docs/user-guide/annotations
    volumes: [volume.Volume], default is Undefined, optional.
        Volumes represents a named volume and corresponding mounts in containers.
    services: [service.Service], default is Undefined, optional.
        Services is a list of Service which partition a single Kubernetes cluster into multiple virtual clusters.

    Examples
    --------
    import models.kube.frontend
    import models.kube.frontend.container
    import models.kube.templates.resource as res_tpl

    appConfiguration: frontend.Server {
        mainContainer = container.Main {
            name = "php-redis"
            env: {
                "GET_HOSTS_FROM": {value = "dns"}
            }
            ports = [{containerPort = 80}]
        }
    }
    """

    # workload name
    name?: str

    # workload namespace
    namespace?: str

    # Application workload type, default to 'Deployment'
    workloadType: "Deployment" | "StatefulSet" = "Deployment"

    # Application replicas
    replicas: int = 1

    # Main container image
    image: str
    # Main container imagePullPolicy
    imagePullPolicy?: "Always" | "IfNotPresent" | "Never"
    # Main container resource

    # Main container configuration
    mainContainer: container.Main
    # Sidecar container configurations

    # Workload configuration
    labels?: {str:str}
    annotations?: {str:str}
    volumes?: [volume.Volume]

    services?: [service.Service]
