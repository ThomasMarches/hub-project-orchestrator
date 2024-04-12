package gen

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

const (
	nameFlag = "name"

	destinationFlag = "dest"

	frontendFlag = "frontend"
	backendFlag  = "backend"

	gitRepo = "git-repo"

	ci = "ci"
)

type generator struct {
	name string

	dest string

	frontend string
	backend  string

	gitRepo bool

	ci bool
}

type backendConfig struct {
	Address string
	Port    string
}

var defaultBackendConfig = &backendConfig{
	Address: "127.0.0.1",
	Port:    "8080",
}

type databaseConfig struct {
	Address string
	Port    string
}

var defaultdDatabaseConfig = &databaseConfig{
	Address: "127.0.0.1",
	Port:    "8080",
}

type ciConfig struct {
	OnPush        bool
	OnPullRequest bool
	Branches      []string
}

var defaultCiConfig = &ciConfig{
	OnPush:        true,
	OnPullRequest: true,
	Branches:      []string{"main", "master"},
}

type frontendConfig struct {
	Address string
	Port    string
	ApiUrl  string
	ApiPort string
}

var defaultFrontendConfig = &frontendConfig{
	Address: "127.0.0.1",
	Port:    "3000",
	ApiUrl:  "127.0.0.1",
	ApiPort: "8080",
}

var Cmd = &cobra.Command{
	Use:   "gen",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		g := &generator{
			name: cmd.Flags().Lookup(nameFlag).Value.String(),

			dest: cmd.Flags().Lookup(destinationFlag).Value.String(),

			frontend: cmd.Flags().Lookup(frontendFlag).Value.String(),
			backend:  cmd.Flags().Lookup(backendFlag).Value.String(),

			gitRepo: cmd.Flags().Lookup(gitRepo).Value.String() == "true",

			ci: cmd.Flags().Lookup(ci).Value.String() == "true",
		}

		if g.dest == "" {
			g.dest = g.name
		}

		err := g.generate()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	Cmd.Flags().StringP(nameFlag, "n", "my-app", "Name of your application.")

	Cmd.Flags().StringP(destinationFlag, "d", "", "Folder to which the project will be generated.")

	Cmd.Flags().StringP(frontendFlag, "f", "", "Frontend you wish to deploy. Leave empty for no frontend.")
	Cmd.Flags().StringP(backendFlag, "b", "", "Backend you wish to deploy. Leave empty for no backend.")

	Cmd.Flags().BoolP(gitRepo, "g", false, "Either you would like to create a git repository in the destination folder.")
	Cmd.Flags().BoolP(ci, "c", false, "Either you would like to generate github-actions CI files.")
}

func (g *generator) generate() error {
	if g.backend != "" {
		err := g.generateBackend()
		if err != nil {
			return err
		}
	}

	if g.frontend != "" {
		err := g.generateFrontend()
		if err != nil {
			return err
		}
	}

	if g.gitRepo {
		err := g.generateGitRepo()
		if err != nil {
			return err
		}
	}

	if g.ci {
		err := g.generateContinuousIntegration()
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *generator) generateContinuousIntegration() error {
	if !g.gitRepo {
		return fmt.Errorf("[ ModuleOrchestrator - Config ] You need to generate a git repository to generate continuous integration files")
	}

	fmt.Println("[ ModuleOrchestrator - Config ] Generating continuous integration files...")

	if g.backend != "" {
		switch strings.Split(g.backend, "-")[0] {
		case "golang":
			err := changeCITemplate(g.name, g.dest, defaultCiConfig, "template/ci/github-actions-go/")
			if err != nil {
				return err
			}
		case "rust":
			err := changeCITemplate(g.name, g.dest, defaultCiConfig, "template/ci/github-actions-rust/")
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported CI for given backend: %v", g.backend)
		}
	}

	switch g.frontend {
	case "angular":
		err := changeCITemplate(g.name, g.dest, defaultCiConfig, "template/ci/github-actions-angular/")
		if err != nil {
			return err
		}
	case "react":
		err := changeCITemplate(g.name, g.dest, defaultCiConfig, "template/ci/github-actions-react/")
		if err != nil {
			return err
		}
	case "flutter":
		err := changeCITemplate(g.name, g.dest, defaultCiConfig, "template/ci/github-actions-flutter/")
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported CI for given frontend: %v", g.frontend)
	}

	fmt.Println("[ ModuleOrchestrator - Config ] Continuous integration files generated !")

	return nil
}

func (g *generator) generateBackend() error {
	fmt.Println("[ ModuleOrchestrator - Backend ] Generating backend module...")

	switch g.backend {
	case "golang":
		err := g.backendGo()
		if err != nil {
			return err
		}
	case "rust-axum":
		err := g.backendRustAxum()
		if err != nil {
			return err
		}
	case "rust-actix":
		err := g.backendRustActix()
		if err != nil {
			return err
		}
	case "rust-rocket":
		err := g.backendRustRocket()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported backend: %v", g.backend)
	}

	fmt.Println("[ ModuleOrchestrator - Backend ] Backend generated !")

	return nil
}

func (g *generator) generateGitRepo() error {
	fmt.Println("[ ModuleOrchestrator - Config ] Generating git repository...")

	cmd := exec.Command("git", "init")

	// specify the working directory of the command
	cmd.Dir = g.dest

	// create a buffer to store the output of your process
	var out bytes.Buffer

	// define the process standard output
	cmd.Stdout = &out

	// Run the command
	err := cmd.Run()

	if err != nil {
		// error case : status code of command is different from 0
		log.Fatal("generate contiunous integration err : ", err)
		return err
	}

	fmt.Println("[ ModuleOrchestrator - Config ] Repository generated !")

	return nil
}

func (g *generator) generateFrontend() error {
	fmt.Println("[ ModuleOrchestrator - Frontend ] Generating frontend module...")

	switch g.frontend {
	case "angular":
		err := g.frontendAngular()
		if err != nil {
			return err
		}
	case "react":
		err := g.frontendReact()
		if err != nil {
			return err
		}
	case "flutter":
		err := g.frontendFlutter()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported frontend: %v", g.frontend)
	}

	fmt.Println("[ ModuleOrchestrator - Frontend ] Frontend generated !")

	return nil
}

func (g *generator) frontendAngular() error {
	return changeFrontendTemplate(g.name, g.dest, defaultFrontendConfig, "template/frontend/angular/template-angular-basic-front/")
}

func (g *generator) frontendFlutter() error {
	cmd := exec.Command("flutter", "create", g.name)

	// specify the working directory of the command
	cmd.Dir = g.dest

	// create a buffer to store the output of your process
	var out bytes.Buffer

	// define the process standard output
	cmd.Stdout = &out

	// Run the command
	err := cmd.Run()

	if err != nil {
		// error case : status code of command is different from 0
		log.Fatal("generate flutter app err : ", err)
	}

	return err
}

func (g *generator) frontendReact() error {
	return changeFrontendTemplate(g.name, g.dest, defaultFrontendConfig, "template/frontend/react/cra-template-basic-front/")
}

func (g *generator) backendGo() error {
	return changeBackendTemplate(g.name, g.dest, defaultBackendConfig, "template/backend/golang/")
}

func (g *generator) backendRustActix() error {
	return changeBackendTemplate(g.name, g.dest, defaultBackendConfig, "template/backend/rust-actix/")
}

func (g *generator) backendRustRocket() error {
	return changeBackendTemplate(g.name, g.dest, defaultBackendConfig, "template/backend/rust-rocket/")
}

func (g *generator) backendRustAxum() error {
	return changeBackendTemplate(g.name, g.dest, defaultBackendConfig, "template/backend/rust-axum/")
}
