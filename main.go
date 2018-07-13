package main
import (
    "fmt"
    "bytes"
    "bufio"
    "os"
    "io/ioutil"
    "io"
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

func copyFileContents(src, dst string) (err error) {
    in, err := os.Open(src)
    if err != nil {
        return
    }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil {
        return
    }
    defer func() {
        cerr := out.Close()
        if err == nil {
            err = cerr
        }
    }()
    if _, err = io.Copy(out, in); err != nil {
        return
    }
    err = out.Sync()
    return
}

func main() {
    prefix := [3]string{"d","s","p"}
    
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("What is your project name ? ")
    project_name, _ := reader.ReadString('')
    
    fmt.Print("What is your image name ? ")
    image_name, _ := reader.ReadString('')

    fmt.Print("What is your exposed port ? ")
    exposed_port, _ := reader.ReadString('')

    fmt.Print("What is your host port ? ")
    host_port, _ := reader.ReadString('')

    fmt.Println("Initializing Project....")
    docker_folder := "dockerfiles"
    os.Mkdir(docker_folder, 0755);

    fmt.Println("Creating Dockerfiles....")
    dockerfiles := [3]string{"app.docker.dev","app.docker.staging","app.docker"}

    for _, value := range dockerfiles {
        copyFileContents("templates/docker-template", docker_folder+"/"+value)
        replace(docker_folder+"/"+value, "_EXPOSED_PORT", exposed_port)    
    }

    fmt.Println("Creating Compose Files....")
    composefiles := [3]string{"docker-compose.dev.yaml","docker-compose.staging.yaml","docker-compose.yaml"}

    for i, value := range composefiles {
        copyFileContents("templates/compose-template", value)
        replace(value, "_PROJECT_NAME", project_name)
        replace(value, "_IMAGE_NAME", prefix[i]+"-"+image_name)
        replace(value, "_EXPOSED_PORT", exposed_port)
        replace(value, "_HOST_PORT", host_port)
        replace(value, "_DOCKER_FILE_NAME", dockerfiles[i])
    }

    fmt.Println("Creating Jenkinsfiles....")
    jenkinsfiles := [2]string{"Jenkinsfile.staging","Jenkinsfile"}

    for i, value := range jenkinsfiles {
        copyFileContents("templates/jenkins-template", value)
        replace(value, "_COMPOSE_FILE_NAME", composefiles[i+1])
        replace(value, "_IMAGE_NAME", prefix[i+1]+"-"+image_name)
    }

    fmt.Println("Done....")
}