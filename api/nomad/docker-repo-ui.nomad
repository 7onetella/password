job "docker-registry-ui" {
  datacenters = ["dc1"]

  type = "service"

  update {
    stagger      = "30s"
    max_parallel = 1
  }

  group "dev" {
    count = 1
        
    task "service" {
        driver = "docker"
        config {
            image = "joxit/docker-registry-ui:static"
            port_map {
                http = 80
            }        
        }

        resources {
            cpu = 20
            memory = 128 # MB
            network {
                mbits = 100
                port "http" {}
            }            
        }

        service {
            tags = ["urlprefix-docker-repo.7onetella.net/"]
            port = "http"
            check {
                type     = "http"
                path     = "/"
                interval = "10s"
                timeout  = "2s"
            }
        }

        env {
            URL = "http://tmt-vm10.7onetella.net:5000"
            DELETE_IMAGES = "true"
        }

    }
  }
}