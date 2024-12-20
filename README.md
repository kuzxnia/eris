# eris


- workflow
    > select resource
        - pod affected percentage 
            (how many resources will be affected from selected app-resource)
        - target pod
            - with regex
        - target container
            - with regex (multiple)


    - attack

    - probe (resilience)





todo:
* add testcontainers for k8s https://golang.testcontainers.org/modules/k3s/

* sposób generowania worklows (chaos init) / online editor???





wejście do aplikacji / usecase:

1. przez kod - tworzenie testów integracyjnych itd. my w python

- python-sdk - client
- golang-sdk - client

    > from url, example 

    import docker
    client = docker.APIClient(base_url='unix://var/run/docker.sock')
    client.version()
    {u'ApiVersion': u'1.33',
     u'Arch': u'amd64',
     u'BuildTime': u'2017-11-19T18:46:37.000000000+00:00',
     u'GitCommit': u'f4ffd2511c',
     u'GoVersion': u'go1.9.2',
     u'KernelVersion': u'4.14.3-1-ARCH',
     u'MinAPIVersion': u'1.12',
     u'Os': u'linux',
     u'Version': u'17.10.0-ce'}

     class APIClient(base_url=None, version=None, timeout=60, tls=False, user_agent='docker-sdk-python/7.1.0', num_pools=None, credstore_env=None, use_ssh_client=False, max_pool_size=10)


2. osobno jako zwykły test

    2.1. client webowy uruchamiający testy + jakiś monitoring etc.

    2.2. client cli uruchamiający to jako chaos as code
        * defiujemy yaml, organizujemy worklowy



Pytania:
1. czy potrzebujemy to móc odpalać na wielu kontekstach???
    - tak, dla operatorów multiple cluster chcemy mieć możliwość definiowania testów 
    
    # to mogłoby być płatne
    across multiple clusters

2. czy centralny agent
    - będzie tutaj wymagany centranly element, do którego będą się rejestrować agenci
        - wystawiony ws/grcp, agent incjuje połaczenie



architektura:

- (optional for now) web-ui
- (optional for now) python-sdk client for / control-plane 
- (optional for now) golang-sdk client for / control-plane 
    - z poziomu sdk musimy mieć możliwość deploy zasobów na k8s, np. za pomocą helmów 

- eris-control-plane (python??) - swagger out of the box, easy 
    - high level api for for running tests across multiple agents
    - może być odpalony lokalnie,
    - można go skonfigurować, żeby łączył się do agentów, albo agentom ustawić adress 
    gdzie mają udeżać

    - może mieć bazke docelowo do storowania workflowów, wyników itp 


- eris-agent (golang, we need k8s fancy tooling like k3s)
    - agent - from here all probes & attacks will run

    / probe/attack jako jobk8s, wszystko wrzucane z control-plane lub bezpośrednio z agenta
    - wady:
    jak kontrolować kiedy wykonane będą probe(przed,przed i po,calyczas)




basic flow:
test jak zachowuje się klaster przy awarii 1 node

    - suspend process - kill -19
    - wait until probe failed (http check) success or max 10s 
        - curl for status with jq. filtering
    - kill process -9
    - wait until probe success max 120s

test jak zachowuje się klaster przy awarii 2 node(1dc lub 1 node z 2 dc)

(potrzebujemy 2 job / lub 2 agentów / lub jeden wykonuje 2 operacje)

    - suspend process - kill -19
    - wait until probe failed (http check) success or max 10s 
        - curl for status with jq. filtering
    - kill process -9
    - wait until probe success max 120s

test jak zachowuje się klaster przy awarii 2 node(1dc lub 1 node z 2 dc)

    - suspend process - kill -19
    - wait until probe failed (http check) success or max 10s 
        - curl for status with jq. filtering
    - kill process -9
    - wait until probe success max 120s


todo: 
* https://github.com/itchyny/gojq for check data
