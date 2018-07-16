package main
import (
    "fmt"
    "bytes"
    "bufio"
    "os"
    "io/ioutil"
    "strings"
)

func replace(src, replace, replacement string) {

    input, err := ioutil.ReadFile(src)
    if err != nil {
            fmt.Println(err)
            os.Exit(1)
    }

    output := bytes.Replace(input, []byte(replace), []byte(replacement), -1)

    if err = ioutil.WriteFile(src, output, 0666); err != nil {
            fmt.Println(err)
            os.Exit(1)
    }
}

func main() {
    prefix := [3]string{"d","s","p"}
    
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("What is your project name (GCP Project name) ? ")
    project_name, _ := reader.ReadString('\n')
    project_name = strings.Replace(project_name, "\n", "", -1)
  
    fmt.Print("What is your base image name (Your base image) ? ")
    base_image_name, _ := reader.ReadString('\n')
    base_image_name = strings.Replace(base_image_name, "\n", "", -1)
    
    fmt.Print("What is your image name (Your image and container name) ? ")
    image_name, _ := reader.ReadString('\n')
    image_name = strings.Replace(image_name, "\n", "", -1)


    fmt.Print("What is your exposed container port ? ")
    exposed_port, _ := reader.ReadString('\n')
    exposed_port = strings.Replace(exposed_port, "\n", "", -1)


    fmt.Print("What is your host to container port ? ")
    host_port, _ := reader.ReadString('\n')
    host_port = strings.Replace(host_port, "\n", "", -1)


    fmt.Println("Initializing Project....")
    docker_folder := "dockerfiles"
    os.Mkdir(docker_folder, 0755);

    docker_template, err := Asset("templates/docker-template")
    if err != nil {
        fmt.Println(Asset("aset not found"))
    }
    
    compose_template, err := Asset("templates/compose-template")
    if err != nil {
        fmt.Println(Asset("aset not found"))
    }

    compose_dev_template, err := Asset("templates/compose-dev-template")
    if err != nil {
        fmt.Println(Asset("aset not found"))
    }

    jenkins_template, err := Asset("templates/jenkins-template")
    if err != nil {
        fmt.Println(Asset("aset not found"))
    }

    fmt.Println("Creating Dockerfiles....")
    dockerfiles := [3]string{"app.docker.dev","app.docker.staging","app.docker"}

    for _, value := range dockerfiles {
        ioutil.WriteFile(docker_folder+"/"+value, []byte(docker_template), 0644)
        replace(docker_folder+"/"+value, "_EXPOSED_PORT", exposed_port)  
        replace(docker_folder+"/"+value, "_BASE_IMAGE", base_image_name)   
    }

    fmt.Println("Creating Compose Files....")
    composefiles := [3]string{"docker-compose.dev.yaml","docker-compose.staging.yaml","docker-compose.yaml"}

    for i, value := range composefiles {
        if value == "docker-compose.dev.yaml" {
            ioutil.WriteFile(value, []byte(compose_dev_template), 0644)
        }else{
            ioutil.WriteFile(value, []byte(compose_template), 0644)
        }
        replace(value, "_PROJECT_NAME", project_name)
        replace(value, "_IMAGE_NAME", prefix[i]+"-"+image_name)
        replace(value, "_EXPOSED_PORT", exposed_port)
        replace(value, "_HOST_PORT", host_port)
        replace(value, "_DOCKER_FILE_NAME", dockerfiles[i])
    }

    fmt.Println("Creating Jenkinsfiles....")
    jenkinsfiles := [2]string{"Jenkinsfile.staging","Jenkinsfile"}

    for i, value := range jenkinsfiles {
        ioutil.WriteFile(value, []byte(jenkins_template), 0644)
        replace(value, "_COMPOSE_FILE_NAME", composefiles[i+1])
        replace(value, "_IMAGE_NAME", prefix[i+1]+"-"+image_name)
    }

    fmt.Println("Done....")
}