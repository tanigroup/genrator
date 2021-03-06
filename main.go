package main
import (
    "flag"
    "fmt"
    "bytes"
    "bufio"
    "os"
    "io/ioutil"
    "strings"
    "github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)
const version = "0.10.0"

func selfUpdate(slug string) error {
	selfupdate.EnableLog()

	previous := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(previous, slug)
	if err != nil {
		return err
	}

	if previous.Equals(latest.Version) {
		fmt.Println("Current binary is the latest version", version)
	} else {
		fmt.Println("Update successfully done to version", latest.Version)
		fmt.Println("Release note:\n", latest.ReleaseNotes)
    }
	return nil
}

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
func findAsset(assetSrc string)([]byte){
    asset_template, err := Asset(assetSrc)
    if err != nil {
        fmt.Println(Asset("aset not found"))
        os.Exit(1)
    }
    return asset_template;
}

func createDockerfiles(exposed_port string, base_image_name string, dockerfiles [2]string, docker_template []byte ){
    fmt.Println("Creating Dockerfiles....")
    docker_folder := "dockerfiles"
    os.Mkdir(docker_folder, 0755);

    for _, value := range dockerfiles {
        ioutil.WriteFile(docker_folder+"/"+value, []byte(docker_template), 0644)
        replace(docker_folder+"/"+value, "_EXPOSED_PORT", exposed_port)  
        replace(docker_folder+"/"+value, "_BASE_IMAGE", base_image_name)   
    }
}

func createDockerCompose(project_name string,image_name string,exposed_port string, host_port string, prefix [3]string, dockerfiles [2]string, composefiles [3]string, compose_dev_template []byte, compose_template []byte){
    fmt.Println("Creating Compose Files....")
    for i, value := range composefiles {
        if value == "docker-compose.dev.yaml" {
            ioutil.WriteFile(value, []byte(compose_dev_template), 0644)
            replace(value, "_DOCKER_FILE_NAME", dockerfiles[0])
        }else{
            ioutil.WriteFile(value, []byte(compose_template), 0644)
            replace(value, "_DOCKER_FILE_NAME", dockerfiles[1])
        }
        replace(value, "_PROJECT_NAME", project_name)
        replace(value, "_IMAGE_NAME", prefix[i]+"-"+image_name)
        replace(value, "_EXPOSED_PORT", exposed_port)
        replace(value, "_HOST_PORT", host_port)
    }
}

func createJenkinsfile(image_name string, exposed_port string, namespace_name string, prefix [3]string, composefiles [3]string, jenkinsfiles [2]string, jenkins_template []byte){
    fmt.Println("Creating Jenkinsfiles....")
    for i, value := range jenkinsfiles {
        ioutil.WriteFile(value, []byte(jenkins_template), 0644)
        replace(value, "_COMPOSE_FILE_NAME", composefiles[i+1])
        replace(value, "_IMAGE_NAME", prefix[i+1]+"-"+image_name)
        
        if prefix[i+1] == "s" {
            replace(value, "_UPDATE_SCRIPT_URL", "https://gitlab.tanihub.com/unrestricted-repository/kubernetes-manifest-command/raw/master/updateStagingManifest.sh")
        }else{
            replace(value, "_UPDATE_SCRIPT_URL", "https://gitlab.tanihub.com/unrestricted-repository/kubernetes-manifest-command/raw/master/updateProductionManifest.sh")            
        }
        
        replace(value, "_EXPOSED_PORT", exposed_port)
        replace(value, "_NAMESPACE_NAME", namespace_name)
    }
}
func kubernetesConfig(project_name string, image_name string, exposed_port string, prefix [3]string, kubernetes_template []byte){
    fmt.Println("Creating Kubernetes Config....")
    for i, value := range prefix {
        if value == "d" {
            continue
        }
        ioutil.WriteFile(value+"-"+image_name+".yaml", []byte(kubernetes_template), 0644)
        replace(value+"-"+image_name+".yaml", "_IMAGE_NAME", prefix[i]+"-"+image_name)
        replace(value+"-"+image_name+".yaml", "_EXPOSED_PORT", exposed_port)
        replace(value+"-"+image_name+".yaml", "_PROJECT_NAME", project_name)
    }
}

func main() {
    ver := flag.Bool("version", false, "Show version")
    jenkinsconf := flag.Bool("jenkinsonly", false, "Create only jenkinsfile")
    update := flag.Bool("selfupdate", false, "Try selfupdate via GitHub")
    slug := flag.String("slug", "tanigroup/genrator", "Repository of this command")
    
    flag.Parse()
    
    if *ver {
		fmt.Println(version)
		os.Exit(0)
    }
     
    if *update {
		if err := selfUpdate(*slug); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(0)
	}
     
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("What is your 'GCP Project' name ? ")
    project_name, _ := reader.ReadString('\n')
    project_name = strings.Replace(project_name, "\n", "", -1)
  
    fmt.Print("What 'base image' you use for this APP ? ")
    base_image_name, _ := reader.ReadString('\n')
    base_image_name = strings.Replace(base_image_name, "\n", "", -1)
    
    fmt.Print("What will be the 'name' of your APP ? ")
    image_name, _ := reader.ReadString('\n')
    image_name = strings.Replace(image_name, "\n", "", -1)
    
    fmt.Print("What will be the 'namespace' of your APP ? ")
    namespace_name, _ := reader.ReadString('\n')
    namespace_name = strings.Replace(namespace_name, "\n", "", -1)

    fmt.Print("What will be your base container port ? ")
    exposed_port, _ := reader.ReadString('\n')
    exposed_port = strings.Replace(exposed_port, "\n", "", -1)

    fmt.Print("What port you want your APP listen in your PC ? ")
    host_port, _ := reader.ReadString('\n')
    host_port = strings.Replace(host_port, "\n", "", -1)

    fmt.Println("Initializing Project....")

    prefix := [3]string{"d","s","p"}
    composefiles := [3]string{"docker-compose.dev.yaml","docker-compose.staging.yaml","docker-compose.yaml"}
    dockerfiles := [2]string{"app.docker.dev","app.docker"}
    jenkinsfiles := [2]string{"Jenkinsfile.staging","Jenkinsfile"}
    docker_template :=findAsset("templates/docker-template")
    compose_template :=findAsset("templates/compose-template")
    compose_dev_template:=findAsset("templates/compose-dev-template")
    jenkins_template:=findAsset("templates/jenkins-template")
    kubernetes_template:=findAsset("templates/kubernetes-template")
   
    flag.Parse()
    if *jenkinsconf {
		fmt.Println("Creating jenkinsfile only")
        createJenkinsfile(image_name, exposed_port, namespace_name, prefix, composefiles, jenkinsfiles,jenkins_template)
		os.Exit(0)
    }
    
    createDockerfiles(exposed_port, base_image_name, dockerfiles, docker_template)
    createDockerCompose(project_name, image_name, exposed_port, host_port, prefix, dockerfiles, composefiles, compose_dev_template, compose_template)
    createJenkinsfile(image_name, exposed_port, namespace_name, prefix, composefiles, jenkinsfiles,jenkins_template)
    kubernetesConfig(project_name, image_name, exposed_port, prefix, kubernetes_template)
    fmt.Println("Done....")
}