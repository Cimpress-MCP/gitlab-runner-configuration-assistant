package main

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
)

type GitlabRunnerConfig struct {
	Concurrent    int `toml:"concurrent"`
	CheckInterval int `toml:"check_interval"`
	Runners []struct {
		Name     string `toml:"name"`
		URL      string `toml:"url"`
		Token    string `toml:"token"`
		Executor string `toml:"executor"`
		Docker struct {
			TLSVerify    bool     `toml:"tls_verify"`
			Image        string   `toml:"image"`
			Privileged   bool     `toml:"privileged"`
			DisableCache bool     `toml:"disable_cache"`
			Volumes      []string `toml:"volumes"`
			ShmSize      int      `toml:"shm_size"`
		} `toml:"docker"`
		Cache struct {
		} `toml:"cache"`
	} `toml:"runners"`
}

func main() {
	println("Redoing gitlab runner config")
	path := "/etc/gitlab-runner/config.toml"
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		println("Gitlab Runner Config not found...")
		return
	}

	gitlab := GitlabRunnerConfig{}
	toml.Unmarshal(configFile, &gitlab)

	gitlab.Concurrent = 2
	gitlab.Runners[0].Docker.Privileged = true
	gitlab.Runners[0].Docker.Volumes = []string{"/var/run/docker.sock:/var/run/docker.sock", "/cache"}

	output, _ := toml.Marshal(gitlab)
	ioutil.WriteFile(path, []byte(output), 0644)
	println("Done!")
	return
}
